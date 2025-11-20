package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"kgplatform-backend/internal/model/entity"
	"kgplatform-backend/internal/utils"
)

type CreateTaskReq struct {
	g.Meta         `path:"/task/create" method:"post" tags:"任务" sm:"创建任务"`
	PipelineId     int    `json:"pipelineId" v:"required#请选择管道"`
	MaterialIdList []int  `json:"materialIdList" v:"required#请选择素材"`
	ProjectId      int    `json:"projectId" v:"required#请选择项目"`
	Type           string `json:"type" v:"required#请选择任务类型"`
	Prompt         string `json:"prompt"`
	ModelId        *int   `json:"modelId" v:"required#请选择模型"`
}

type CreateTaskRes struct {
	TaskId int    `json:"id"`
	Status string `json:"status"`
	Type   string `json:"type"`
}

type GetTaskReq struct {
	g.Meta `path:"/task/get/{id}" method:"get" tags:"任务" sm:"获取任务"`
	Id     int `path:"id" v:"required#请选择任务"`
}

type GetTaskRes struct {
	Id         int    `json:"id"`
	Status     string `json:"status" dc:"任务状态, pending, processing, completed, failed"`
	Type       string `json:"type" dc:"任务类型, ocr, extract, graph"`
	CreateTime string `json:"createTime"`
	StartTime  string `json:"startTime" dc:"任务开始处理的时间"`
	FinishTime string `json:"finishTime"`
}

type CreateOCRTaskReq struct {
	g.Meta         `path:"/task/ocr" method:"post" tags:"任务" sm:"创建OCR任务"`
	MaterialIDList []int `json:"materialIdList" v:"required#请选择素材"`
	ProjectId      int   `json:"projectId" v:"required#请选择项目"`
	PipelineId     int   `json:"pipelineId" v:"required#请选择管道"`
}

type CreateOCRTaskRes struct {
	TaskId int    `json:"id"`
	Status string `json:"status" dc:"任务状态, pending, processing, completed, failed"`
}

type CreateExtractTaskReq struct {
	g.Meta         `path:"/task/extract" method:"post" tags:"任务" sm:"创建抽取任务"`
	MaterialIDList []int  `json:"materialIdList" v:"required#请选择素材"`
	ProjectId      int    `json:"projectId" v:"required#请选择项目"`
	PipelineId     int    `json:"pipelineId" v:"required#请选择管道"`
	Method         string `json:"method" v:"required#请选择抽取方式"`
	Prompt         string `json:"prompt" v:"required#请输入提示词"`
	ModelId        *int   `json:"modelId" v:"required#请选择大模型"`

	// genprompt 接口返回的配置
	SchemaURL              string   `json:"schemaUrl"`
	SampleTextURL          string   `json:"sampleTextUrl"`
	SampleXLSXURL          string   `json:"sampleXlsxUrl"`
	TargetDomain           string   `json:"targetDomain"`
	DictionaryId           *int     `json:"dictionaryId"`
	DictionaryURL          string   `json:"dictionaryUrl"`
	PriorityExtractions    []string `json:"priorityExtractions"`
	ExtractionRequirements string   `json:"extractionRequirements"`
}

type CreateExtractTaskRes struct {
	TaskId int    `json:"id"`
	Status string `json:"status" dc:"任务状态, pending, processing, completed, failed"`
}

type CreateGraphTaskReq struct {
	g.Meta         `path:"/task/graph" method:"post" tags:"任务" sm:"创建图谱任务"`
	MaterialIDList []int `json:"materialIdList" v:"required#请选择素材"`
	ProjectId      int   `json:"projectId" v:"required#请选择项目"`
	PipelineId     int   `json:"pipelineId" v:"required#请选择管道"`
}

type CreateGraphTaskRes struct {
	TaskId int    `json:"id"`
	Status string `json:"status" dc:"任务状态, pending, processing, completed, failed"`
}

type ListTaskReq struct {
	g.Meta     `path:"/task/list" method:"get" tags:"任务" sm:"获取任务列表"`
	Page       utils.IPage `json:"page" dc:"分页"`
	ProjectId  *int        `json:"projectId"`
	PipelineId *int        `json:"pipelineId"`
}
type ListTaskRes struct {
	Total int             `json:"total" dc:"总记录数"`
	List  []*entity.Tasks `json:"list" dc:"任务列表"`
}
