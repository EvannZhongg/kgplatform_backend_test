package py_service

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/google/uuid"

	"io"
	"kgplatform-backend/internal/consts"
	"kgplatform-backend/internal/dao"
	"kgplatform-backend/internal/logic/upload"
	"kgplatform-backend/internal/model/entity"
	"kgplatform-backend/internal/neo4j"
	"kgplatform-backend/internal/utils"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PythonClient Python 服务客户端
type PythonClient struct {
	baseURL    string
	httpClient *http.Client
	sseManager *SSEManager
}

// SSEManager SSE 连接管理器
type SSEManager struct {
	connections map[string]*SSEConnection // task_id -> connection
	mutex       sync.RWMutex
}

// SSEConnection SSE 连接
type SSEConnection struct {
	taskID     string
	ctx        context.Context
	cancel     context.CancelFunc
	response   *http.Response
	reader     *bufio.Scanner
	statusChan chan *PythonTaskStatus
	errorChan  chan error
	closed     bool
	mutex      sync.RWMutex
}

// PythonCreateTaskRequest Python 任务请求
type PythonCreateTaskRequest struct {
	Files      []File `json:"files"`
	PromptText string `json:"prompt_text"`
	Provider   string `json:"provider"`
	Model      string `json:"model,omitempty"`
	APIKey     string `json:"api_key"`
	BaseURL    string `json:"base_url,omitempty"`
}

type File struct {
	MaterialId int    `json:"material_id"`
	URL        string `json:"url"`
}

// PythonCreateTaskResponse Python 任务响应
type PythonCreateTaskResponse struct {
	TaskID  string `json:"task_id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// PythonTaskStatus Python 任务状态
type PythonTaskStatus struct {
	TaskID      string     `json:"task_id"`
	Status      string     `json:"status"`
	Progress    float64    `json:"progress,omitempty"`
	Message     string     `json:"message,omitempty"`
	Error       string     `json:"error,omitempty"`
	Result      []TaskFile `json:"results,omitempty"`
	CreatedAt   string     `json:"created_at,omitempty"`
	StartedAt   string     `json:"started_at,omitempty"`
	CompletedAt string     `json:"completed_at,omitempty"`
	Type        string     `json:"type,omitempty"`
	Timestamp   string     `json:"timestamp,omitempty"`
}

// TaskFile 任务文件结果
type TaskFile struct {
	FileName     string            `json:"file_name"`
	MaterialId   int               `json:"material_id"`
	Status       string            `json:"status"`
	TriplesCount int               `json:"triples_count"`
	OutputFiles  map[string]string `json:"output_files"`
	Error        string            `json:"error,omitempty"`
}

type PythonGenPromptRequest struct {
	SchemaURL              string    `json:"schema_url"`                        // 必填
	SampleTextURL          *string   `json:"sample_text_url,omitempty"`         // 选填
	SampleXLSXURL          *string   `json:"sample_xlsx_url,omitempty"`         // 选填
	TargetDomain           *string   `json:"target_domain,omitempty"`           // 选填
	DictionaryURL          *string   `json:"dictionary_url,omitempty"`          // 选填
	PriorityExtractions    *[]string `json:"priority_extractions,omitempty"`    // 选填
	ExtractionRequirements *string   `json:"extraction_requirements,omitempty"` // 选填
	BaseInstruction        *string   `json:"base_instruction,omitempty"`        // 选填
}

type PythonGenPromptResponse struct {
	Prompt                 string   `json:"prompt"`
	SchemaURL              string   `json:"schema_url,omitempty"`
	SampleTextURL          string   `json:"sample_text_url,omitempty"`
	SampleXLSXURL          string   `json:"sample_xlsx_url,omitempty"`
	TargetDomain           string   `json:"target_domain,omitempty"`
	DictionaryURL          string   `json:"dictionary_url,omitempty"`
	PriorityExtractions    []string `json:"priority_extractions,omitempty"`
	ExtractionRequirements string   `json:"extraction_requirements,omitempty"`
	Error                  *string  `json:"error,omitempty"`
	Message                string   `json:"message,omitempty"`
}

var (
	pythonClient *PythonClient
	clientOnce   sync.Once
)

// GetPythonClient 获取 Python 客户端单例
func GetPythonClient() *PythonClient {
	clientOnce.Do(func() {
		pythonClient = NewPythonClient()
	})
	return pythonClient
}

// NewPythonClient 创建新的 Python 客户端
func NewPythonClient() *PythonClient {
	// 从配置中获取 Python 服务地址
	baseURL := g.Cfg().MustGet(context.Background(), "python.baseUrl", "http://localhost:8000").String()

	return &PythonClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 0,
		},
		sseManager: &SSEManager{
			connections: make(map[string]*SSEConnection),
		},
	}
}

// CreateTask 创建 Python 三元组抽取任务
func (c *PythonClient) CreateTask(ctx context.Context, req *PythonCreateTaskRequest) (*PythonCreateTaskResponse, error) {
	url := fmt.Sprintf("%s/api/v1/tasks", c.baseURL)

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, gerror.Newf("序列化请求失败: %v", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, gerror.Newf("创建HTTP请求失败: %v", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, gerror.Newf("发送HTTP请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, gerror.Newf("Python服务返回错误: %d, %s", resp.StatusCode, string(body))
	}

	var response PythonCreateTaskResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, gerror.Newf("解析响应失败: %v", err)
	}

	g.Log().Infof(ctx, "Python任务创建成功: %s", response.TaskID)
	return &response, nil
}

// GetTaskStatus 获取 Python 任务状态
func (c *PythonClient) GetTaskStatus(ctx context.Context, taskID string) (*PythonTaskStatus, error) {
	url := fmt.Sprintf("%s/api/v1/tasks/%s", c.baseURL, taskID)

	httpReq, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, gerror.Newf("创建HTTP请求失败: %v", err)
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, gerror.Newf("发送HTTP请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, gerror.Newf("Python服务返回错误: %d, %s", resp.StatusCode, string(body))
	}

	var status PythonTaskStatus
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return nil, gerror.Newf("解析响应失败: %v", err)
	}

	return &status, nil
}

// CancelTask 取消 Python 任务
func (c *PythonClient) CancelTask(ctx context.Context, taskID string) error {
	url := fmt.Sprintf("%s/api/v1/tasks/%s", c.baseURL, taskID)

	httpReq, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return gerror.Newf("创建HTTP请求失败: %v", err)
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return gerror.Newf("发送HTTP请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return gerror.Newf("Python服务返回错误: %d, %s", resp.StatusCode, string(body))
	}

	g.Log().Infof(ctx, "Python任务取消成功: %s", taskID)
	return nil
}

// StartSSEConnection 启动 SSE 连接监听任务状态
// pyTaskID: 调用 py 侧的 createTask 生成的任务ID
// goTaskID: go 侧 Task 的 ID
func (c *PythonClient) StartSSEConnection(ctx context.Context, pyTaskID string, goTaskID int) error {
	c.sseManager.mutex.Lock()
	defer c.sseManager.mutex.Unlock()

	// 检查是否已存在连接
	if _, exists := c.sseManager.connections[pyTaskID]; exists {
		return gerror.Newf("任务 %s 的SSE连接已存在", pyTaskID)
	}

	url := fmt.Sprintf("%s/api/v1/tasks/%s/stream", c.baseURL, pyTaskID)

	sseCtx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)

	httpReq, err := http.NewRequestWithContext(sseCtx, "GET", url, nil)
	if err != nil {
		cancel()
		return gerror.Newf("创建SSE请求失败: %v", err)
	}

	httpReq.Header.Set("Accept", "text/event-stream")
	httpReq.Header.Set("Cache-Control", "no-cache")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		cancel()
		return gerror.Newf("发送SSE请求失败: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		cancel()
		resp.Body.Close()
		return gerror.Newf("SSE连接失败: %d", resp.StatusCode)
	}

	connection := &SSEConnection{
		taskID:     pyTaskID,
		ctx:        sseCtx,
		cancel:     cancel,
		response:   resp,
		reader:     bufio.NewScanner(resp.Body),
		statusChan: make(chan *PythonTaskStatus, 10),
		errorChan:  make(chan error, 1),
	}

	c.sseManager.connections[pyTaskID] = connection

	// 启动 SSE 读取协程
	go func() {
		err = c.handleSSEConnection(connection, goTaskID)
		if err != nil {
			g.Log().Errorf(ctx, "SSE处理异常: %v", err)
			_, err = dao.Tasks.Ctx(ctx).Where("id", goTaskID).Update(g.Map{
				"status":        consts.TaskStatusFailed,
				"error_message": err.Error(),
				"updated_at":    gtime.Now(),
			})
			c.closeSSEConnection(pyTaskID)
		}
	}()

	g.Log().Infof(ctx, "SSE连接已启动: %s", pyTaskID)
	return nil
}

// handleSSEConnection 处理 SSE 连接
func (c *PythonClient) handleSSEConnection(conn *SSEConnection, goTaskID int) error {
	defer func() {
		c.closeSSEConnection(conn.taskID)
		if r := recover(); r != nil {
			g.Log().Errorf(context.Background(), "SSE连接处理异常: %v", r)
		}
	}()

	for conn.reader.Scan() {
		line := strings.TrimSpace(conn.reader.Text())

		// 跳过空行和非数据行
		if line == "" || !strings.HasPrefix(line, "data: ") {
			continue
		}

		// 提取数据部分
		data := strings.TrimPrefix(line, "data: ")
		if data == "" {
			continue
		}

		// 解析 JSON 数据
		var status PythonTaskStatus
		if err := json.Unmarshal([]byte(data), &status); err != nil {
			g.Log().Errorf(context.Background(), "解析SSE数据失败: %v, data: %s", err, data)
			return err
		}

		// 处理心跳消息
		if status.Type == "heartbeat" {
			continue
		}

		// 跳过 created 和 processing 状态
		if status.Status == "created" || status.Status == "pending" || status.Status == "processing" {
			continue
		}

		ctx := context.Background()
		tx, err := g.DB().Begin(ctx)
		if err != nil {
			return err
		}
		defer func() {
			if err != nil {
				tx.Rollback()
			} else {
				tx.Commit()
			}
		}()

		var task entity.Tasks
		err = dao.Tasks.Ctx(ctx).TX(tx).Where("id", goTaskID).Scan(&task)
		if err != nil {
			return err
		}

		var project entity.Projects
		err = dao.Projects.Ctx(ctx).Where("id", task.ProjectId).Scan(&project)
		if err != nil {
			g.Log().Errorf(ctx, "获取项目信息失败: %v", err)
			return gerror.New("获取项目信息失败")
		}

		// 更新 Go 任务状态
		if err := c.updateGoTaskStatus(ctx, tx, task.Id, &status); err != nil {
			g.Log().Errorf(ctx, "更新Go任务状态失败: %v", err)
			return err
		}

		// 更新素材提取URL
		if status.Status == "completed" {
			if err := c.updateMaterialsExtractURL(ctx, tx, &task, &project, &status); err != nil {
				g.Log().Errorf(ctx, "更新素材提取URL失败: %v", err)
				return err
			}
		}

		// 更新任务进度
		if project.ProjectProgress < 3 {
			_, err = dao.Projects.Ctx(ctx).Where("id", project.Id).Update(g.Map{
				"project_progress": 3,
			})
			if err != nil {
				g.Log().Errorf(ctx, "更新项目进度失败: %v", err)
				return gerror.Newf("更新项目进度失败: %v", err)
			}
		}

		// 如果任务完成，退出循环
		if status.Status == "completed" || status.Status == "failed" || status.Status == "cancelled" {
			g.Log().Infof(context.Background(), "Python任务完成: %s, 状态: %s", conn.taskID, status.Status)
			break
		}
	}

	// 检查扫描错误
	if err := conn.reader.Err(); err != nil {
		g.Log().Errorf(context.Background(), "SSE连接读取错误: %v", err)
	}

	return nil
}

// updateGoTaskStatus 更新 Go 任务状态
func (c *PythonClient) updateGoTaskStatus(ctx context.Context, tx gdb.TX, goTaskID int, pythonStatus *PythonTaskStatus) error {
	var goStatus string
	var errorMessage string
	var finishTime *gtime.Time

	switch pythonStatus.Status {
	case "created", "pending":
		goStatus = consts.TaskStatusPending
	case "processing":
		goStatus = consts.TaskStatusProcessing
	case "completed":
		goStatus = consts.TaskStatusCompleted
		finishTime = gtime.Now()
	case "failed":
		goStatus = consts.TaskStatusFailed
		errorMessage = pythonStatus.Error
		if errorMessage == "" {
			errorMessage = pythonStatus.Message
		}
		finishTime = gtime.Now()
	case "cancelled":
		goStatus = consts.TaskStatusFailed
		errorMessage = "任务已取消"
		finishTime = gtime.Now()
	default:
		return nil // 忽略未知状态
	}

	updateData := g.Map{
		"status":     goStatus,
		"updated_at": gtime.Now(),
	}

	if errorMessage != "" {
		updateData["error_message"] = errorMessage
	}

	if finishTime != nil {
		updateData["finish_time"] = finishTime
	}

	_, err := dao.Tasks.Ctx(ctx).TX(tx).Where("id", goTaskID).Update(updateData)
	if err != nil {
		return gerror.Newf("更新任务状态失败: %v", err)
	}

	g.Log().Infof(ctx, "Go Task状态已更新: ID=%d, Status=%s", goTaskID, goStatus)
	return nil
}

// closeSSEConnection 关闭 SSE 连接
func (c *PythonClient) closeSSEConnection(taskID string) {
	c.sseManager.mutex.Lock()
	defer c.sseManager.mutex.Unlock()

	if conn, exists := c.sseManager.connections[taskID]; exists {
		conn.mutex.Lock()
		if !conn.closed {
			conn.cancel()
			conn.response.Body.Close()
			close(conn.statusChan)
			close(conn.errorChan)
			conn.closed = true
		}
		conn.mutex.Unlock()

		delete(c.sseManager.connections, taskID)
		g.Log().Infof(context.Background(), "SSE连接已关闭: %s", taskID)
	}
}

// StopSSEConnection 停止指定任务的 SSE 连接
func (c *PythonClient) StopSSEConnection(taskID string) {
	c.closeSSEConnection(taskID)
}

// StopAllSSEConnections 停止所有 SSE 连接
func (c *PythonClient) StopAllSSEConnections() {
	c.sseManager.mutex.Lock()
	defer c.sseManager.mutex.Unlock()

	for taskID := range c.sseManager.connections {
		c.closeSSEConnection(taskID)
	}

	g.Log().Info(context.Background(), "所有SSE连接已关闭")
}

// HealthCheck 健康检查
func (c *PythonClient) HealthCheck(ctx context.Context) error {
	url := fmt.Sprintf("%s/health", c.baseURL)

	httpReq, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return gerror.Newf("创建健康检查请求失败: %v", err)
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return gerror.Newf("健康检查请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return gerror.Newf("Python服务不健康: %d", resp.StatusCode)
	}

	return nil
}

func (c *PythonClient) updateMaterialsExtractURL(ctx context.Context, tx gdb.TX, task *entity.Tasks, project *entity.Projects, result *PythonTaskStatus) error {
	var err error
	uploadLogic := upload.NewUpload()
	var materialList []entity.Materials
	err = dao.Materials.Ctx(ctx).TX(tx).WhereIn("id", task.MaterialIdList).Scan(&materialList)
	if err != nil {
		return err
	}

	// 创建文件名到输出文件路径的映射
	extractMapping := make(map[int]string)
	for _, r := range result.Result {
		if r.Status == "success" && r.OutputFiles != nil {
			// 使用jsonl文件路径作为提取结果路径
			if jsonlPath, exists := r.OutputFiles["jsonl"]; exists {
				extractMapping[r.MaterialId] = jsonlPath
			}
		}
	}

	var projectTripleList []neo4j.SimpleTriple
	// 计算并更新文字量
	totalWords := 0
	userId := project.UserId
	// 获取当前用户的订阅记录和套餐信息
	var subscription entity.UserSubscriptions
	err = dao.UserSubscriptions.Ctx(ctx).TX(tx).
		Where("user_id", userId).
		Scan(&subscription)
	if err != nil {
		g.Log().Errorf(ctx, "获取用户订阅记录失败: %v, 用户ID: %d", err, userId)
	}

	// 定义套餐字数配额
	free_quota := g.Cfg().MustGet(ctx, "plans.free.words_quota").Int()
	professional_quota := g.Cfg().MustGet(ctx, "plans.professional.words_quota").Int()
	team_quota := g.Cfg().MustGet(ctx, "plans.team.words_quota").Int()
	plansQuota := map[string]int{
		"free":         free_quota,
		"professional": professional_quota,
		"team":         team_quota,
	}

	// 计算当前任务的总字数（先获取所有文件的字数）
	for _, material := range materialList {
		extractResultPath, exists := extractMapping[material.Id]
		if !exists {
			continue
		}

		text, err := utils.ExtractTextFromFile(ctx, extractResultPath)
		if err != nil {
			g.Log().Errorf(ctx, "提取文件文本失败: %v, 文件路径: %s", err, extractResultPath)
			continue
		}

		taskWords := len(text)
		totalWords += taskWords
	}

	// 应用AI模型的成本系数
	var modelCostMultiplier float64 = 1.0
	if subscription.SelectedAiModel != "" {
		// 尝试从配置中获取模型的成本系数
		modelName := strings.ToLower(subscription.SelectedAiModel)
		// 首先检查中文模型
		modelCostMultiplier = g.Cfg().MustGet(ctx, fmt.Sprintf("ai_models.chinese.%s.cost_multiplier", modelName)).Float64()

		// 如果中文模型中没有找到，尝试检查英文模型
		if modelCostMultiplier == 0 {
			modelCostMultiplier = g.Cfg().MustGet(ctx, fmt.Sprintf("ai_models.english.%s.cost_multiplier", modelName)).Float64()
		}

		// 如果配置中仍未找到，使用默认值1.0
		if modelCostMultiplier == 0 {
			modelCostMultiplier = 1.0
			g.Log().Warningf(ctx, "未找到模型 %s 的成本系数配置，使用默认值: %f", subscription.SelectedAiModel, modelCostMultiplier)
		}
	}
	// 计算实际消耗的字数（原始字数 × 成本系数）
	actualWordsUsed := int(float64(totalWords) * modelCostMultiplier)

	// 确保至少计算1个字
	if actualWordsUsed < 1 {
		actualWordsUsed = 1
	}

	// 检查免费版用户是否超过配额，免费版不允许超额使用
	if subscription.UserPlan == "free" {
		quota := plansQuota[subscription.UserPlan]
		currentUsage := int(subscription.WordsUsed) + actualWordsUsed
		if currentUsage > quota {
			// 使用独特的错误码1002表示免费版用户字数超出配额
			return gerror.NewCodef(
				gcode.New(1002, "FreePlanWordQuotaExceeded", "免费版字数超出配额"),
				"免费版用户字数已超出配额，当前用量: %d, 配额: %d", currentUsage, quota,
			)
		}
	} else {
		quota := plansQuota[subscription.UserPlan]
		currentUsage := int(subscription.WordsUsed) + actualWordsUsed
		if currentUsage > quota {
			// 使用独特的错误码1003表示付费版用户字数超出配额
			return gerror.NewCodef(
				gcode.New(1003, "PaidPlanWordQuotaExceeded", "付费版字数超出配额"),
				"用户字数已超出配额，请注意用量", currentUsage, quota,
			)
		}

		//TODO: 添加字数超额提醒（目前已使用特殊code来表示，0表示成功）
	}

	// 更新每个素材的三元组URL
	for _, material := range materialList {
		extractResultPath, exists := extractMapping[material.Id]
		if !exists {
			continue
		}

		text, err := utils.ExtractTextFromFile(ctx, extractResultPath)
		if err != nil {
			g.Log().Errorf(ctx, "提取文件文本失败: %v, 文件路径: %s", err, extractResultPath)
			continue
		}

		// 解析三元组（包含溯源信息）
		var materialTripleDataList []map[string]interface{}
		err = json.Unmarshal([]byte(text), &materialTripleDataList)
		if err != nil {
			g.Log().Errorf(ctx, "三元组格式错误: %v, 文件路径: %s", err, extractResultPath)
			continue
		}

		// 转换为SimpleTriple并添加溯源信息
		materialTripleListProcessed := []neo4j.SimpleTriple{}

		for _, tripleData := range materialTripleDataList {
			triple := neo4j.SimpleTriple{}

			// 解析标准字段（head, relationship, tail）
			if headData, ok := tripleData["head"].(map[string]interface{}); ok {
				triple.Head.Type = getStringValue(headData, "type")
				triple.Head.Label = getStringValue(headData, "label")
			}
			if relData, ok := tripleData["relationship"].(map[string]interface{}); ok {
				triple.Relationship.Type = getStringValue(relData, "type")
				triple.Relationship.Label = getStringValue(relData, "label")
			}
			if tailData, ok := tripleData["tail"].(map[string]interface{}); ok {
				triple.Tail.Type = getStringValue(tailData, "type")
				triple.Tail.Label = getStringValue(tailData, "label")
			}

			// 添加溯源信息（如果存在_chunk_index和_source_text）
			if chunkIndexVal, ok := tripleData["_chunk_index"].(float64); ok {
				sourceText := ""
				if textVal, ok := tripleData["_source_text"].(string); ok {
					sourceText = textVal
				}

				// 获取素材名称（简化处理，使用URL的文件名部分）
				materialName := material.Url
				if lastSlash := strings.LastIndex(material.Url, "/"); lastSlash >= 0 {
					materialName = material.Url[lastSlash+1:]
				}

				triple.SourceInfo = &neo4j.TripleSourceInfo{
					MaterialId:   material.Id,
					MaterialName: materialName,
					ChunkIndex:   int(chunkIndexVal),
					SourceText:   sourceText,
				}
			}

			materialTripleListProcessed = append(materialTripleListProcessed, triple)
		}

		// 添加到项目三元组列表
		projectTripleList = append(projectTripleList, materialTripleListProcessed...)

		// 上传三元组（包含SourceInfo）
		uploadFileName := utils.RemoveExt(material.Url)
		materialTripleContent, _ := json.Marshal(materialTripleListProcessed)

		saveDataOutput, err := uploadLogic.SaveData(ctx, &upload.SaveDataInput{
			FileName: uploadFileName,
			Content:  string(materialTripleContent),
			DataType: "json",
			UserId:   userId,
		})
		if err != nil {
			g.Log().Errorf(ctx, "保存数据失败: %v, 文件名: %s", err, material.Url)
			continue
		}

		_, err = dao.Materials.Ctx(ctx).TX(tx).Where("id", material.Id).Update(g.Map{
			"triple_url": saveDataOutput.FileName,
		})
		if err != nil {
			g.Log().Errorf(ctx, "更新素材三元组URL失败: %v, 素材ID: %d", err, material.Id)
			continue
		}
	}

	// 更新项目的三元组url
	uuidStr := uuid.New().String()
	timestamp := time.Now().Format("20060102150405")
	uploadFileName := utils.RemoveExt("triples_project_" + strconv.Itoa(task.ProjectId) + "_" + timestamp + "_" + uuidStr[:8])

	// 序列化包含SourceInfo的三元组列表
	uploadContent, _ := json.Marshal(projectTripleList)
	saveDataOutput, err := uploadLogic.SaveData(ctx, &upload.SaveDataInput{
		FileName: uploadFileName,
		Content:  string(uploadContent),
		DataType: "json",
		UserId:   userId,
	})
	if err != nil {
		g.Log().Errorf(ctx, "保存数据失败: %v, 文件名: %s", err, uploadFileName)
		return err
	}
	_, err = dao.Projects.Ctx(ctx).TX(tx).Where("id", task.ProjectId).Update(g.Map{
		"triple_url": saveDataOutput.FileName,
	})

	// 按照三元组type分类存储
	var tripeTypeMap = make(map[string][]neo4j.SimpleTriple)
	for _, triple := range projectTripleList {
		tripleType := fmt.Sprintf("%s-%s-%s", triple.Head.Type, triple.Relationship.Type, triple.Tail.Type)
		tripeTypeMap[tripleType] = append(tripeTypeMap[tripleType], triple)
	}
	var tripleTypeFilenameMap = make(map[string]string)
	for tripleType, triples := range tripeTypeMap {
		uuidStr = uuid.New().String()
		timestamp = time.Now().Format("20060102150405")
		uploadFileName = utils.RemoveExt("triples_project_type" + tripleType + "_" + timestamp + "_" + uuidStr[:8])
		uploadContent, _ = json.Marshal(triples)
		saveMapOutput, err := uploadLogic.SaveData(ctx, &upload.SaveDataInput{
			FileName: uploadFileName,
			Content:  string(uploadContent),
			DataType: "json",
			UserId:   userId,
		})
		if err != nil {
			g.Log().Errorf(ctx, "保存数据失败: %v, 文件名: %s", err, uploadFileName)
			continue
		}
		tripleTypeFilenameMap[tripleType] = saveMapOutput.FileName
	}

	// 修改projects表
	_, err = dao.Projects.Ctx(ctx).TX(tx).Where("id", task.ProjectId).Update(g.Map{
		"triple_url":      saveDataOutput.FileName,
		"triple_type_url": tripleTypeFilenameMap,
	})

	// 更新用户订阅表中的文字用量
	if totalWords > 0 && userId > 0 {
		// 再次获取用户订阅记录（确保数据最新）
		var subscriptionUpdate entity.UserSubscriptions
		err := dao.UserSubscriptions.Ctx(ctx).TX(tx).
			Where("user_id", userId).
			Scan(&subscriptionUpdate)
		if err != nil {
			g.Log().Errorf(ctx, "获取用户订阅记录失败: %v, 用户ID: %d", err, userId)
		}

		// 更新文字用量
		newWordsUsed := int64(subscriptionUpdate.WordsUsed) + int64(totalWords)
		_, err = dao.UserSubscriptions.Ctx(ctx).TX(tx).
			Where("user_id", userId).
			Update(g.Map{
				"words_used": newWordsUsed,
				"updated_at": gtime.Now(),
			})
		if err != nil {
			g.Log().Errorf(ctx, "更新用户文字用量失败: %v, 用户ID: %d", err, userId)
		} else {
			g.Log().Infof(ctx, "用户文字用量已更新: 用户ID=%d, 新增字数=%d, 总字数=%d",
				userId, totalWords, newWordsUsed)
		}
	}

	return nil
}

type ExtractConfig struct {
	Prompt                 string   `json:"prompt"`
	ModelId                int      `json:"modelId"`
	Method                 string   `json:"method"`
	SchemaURL              string   `json:"schemaUrl"`
	SampleTextURL          string   `json:"sampleTextUrl"`
	SampleXLSXURL          string   `json:"sampleXlsxUrl"`
	TargetDomain           string   `json:"targetDomain"`
	DictionaryId           *int     `json:"dictionaryId"`
	DictionaryURL          string   `json:"dictionaryUrl"`
	PriorityExtractions    []string `json:"priorityExtractions"`
	ExtractionRequirements string   `json:"extractionRequirements"`
}

//func getUserIdFromToken(ctx context.Context) (int, error) {
//	r := ghttp.RequestFromCtx(ctx)
//	tokenString := r.Header.Get("Authorization")
//
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		return []byte(consts.JwtKey), nil
//	})
//	if err != nil || !token.Valid {
//		return 0, gerror.New("无效的认证token")
//	}
//
//	claims, ok := token.Claims.(jwt.MapClaims)
//	if !ok {
//		return 0, gerror.New("无效的token claims")
//	}
//
//	userIdFloat, ok := claims["Id"].(float64)
//	if !ok {
//		return 0, gerror.New("token中缺少用户ID")
//	}
//
//	return int(userIdFloat), nil
//}

func (c *PythonClient) GenPrompt(ctx context.Context, req *PythonGenPromptRequest) (*PythonGenPromptResponse, error) {
	url := fmt.Sprintf("%s/api/v1/genprompt", c.baseURL)
	g.Log().Infof(ctx, "请求Python服务: %s", url)
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, gerror.Newf("序列化请求失败: %v", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, gerror.Newf("创建HTTP请求失败: %v", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, gerror.Newf("发送HTTP请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		errorMsg := string(body)
		// 解码Unicode转义字符
		decodedErrorMsg := decodeUnicodeEscapes(errorMsg)
		return nil, gerror.Newf("Python服务返回错误: %d, %s", resp.StatusCode, decodedErrorMsg)
	}

	var response PythonGenPromptResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, gerror.Newf("解析响应失败: %v", err)
	}

	g.Log().Infof(ctx, "Prompt生成成功")
	return &response, nil
}

// decodeUnicodeEscapes 解码Unicode转义字符
func decodeUnicodeEscapes(s string) string {
	re := regexp.MustCompile(`\\u([0-9a-fA-F]{4})`)
	return re.ReplaceAllStringFunc(s, func(match string) string {
		unicodeStr := strings.TrimPrefix(match, "\\u")
		if code, err := strconv.ParseInt(unicodeStr, 16, 32); err == nil {
			return string(rune(code))
		}
		return match
	})
}

// getStringValue 安全地从map[string]interface{}中提取字符串值
func getStringValue(data map[string]interface{}, key string) string {
	if val, ok := data[key]; ok {
		if strVal, ok := val.(string); ok {
			return strVal
		}
		// 如果value不是string类型，尝试转换为string
		return fmt.Sprintf("%v", val)
	}
	return ""
}
