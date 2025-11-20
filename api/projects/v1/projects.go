package v1

import (
	"kgplatform-backend/external/py_service"
	"kgplatform-backend/internal/logic/materials"
	"kgplatform-backend/internal/logic/projects"
	"kgplatform-backend/internal/logic/tasks"
	"kgplatform-backend/internal/model/entity"
	"kgplatform-backend/internal/neo4j"
	"kgplatform-backend/internal/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type CreateProjectReq struct {
	g.Meta      `path:"projects/create" method:"post" sm:"创建项目" tags:"项目管理"`
	ProjectName string `json:"projectName" v:"required|length:1,200" dc:"项目名称"`
	Visibility  int    `json:"visibility" dc:"项目可见性, 0-private, 1-public"`
}

type CreateProjectRes struct {
	ProjectId  int `json:"projectId" dc:"项目ID"`
	PipelineId int `json:"pipelineId" dc:"关联的图谱ID"`
}

type GetProjectReq struct {
	g.Meta `path:"projects/{id}" method:"get" sm:"获取项目详情" tags:"项目管理"`
	Id     int `json:"id" v:"required|min:1" dc:"项目ID"`
}

type GetProjectRes struct {
	Id              int                         `json:"id" dc:"项目ID"`
	UserId          int                         `json:"userId" dc:"用户ID"`
	ProjectName     string                      `json:"projectName" dc:"项目名称"`
	ProjectProgress int                         `json:"projectProgress" dc:"项目进度(0-100)"`
	GraphId         *int                        `json:"graphId" dc:"关联的图谱ID"`
	Materials       []*materials.MaterialDetail `json:"materials" dc:"材料详情列表"`
	ExtractType     string                      `json:"extractType" dc:"抽取类型"`
	CreatedAt       *gtime.Time                 `json:"createdAt" dc:"创建时间"`
	UpdatedAt       *gtime.Time                 `json:"updatedAt" dc:"更新时间"`
	PipelineId      int                         `json:"pipelineId" dc:"关联的图谱ID"`

	// TripleList 三元组列表
	TripleList []neo4j.SimpleTriple `json:"tripleList" dc:"三元组列表"`

	// Schema 主体结构
	Schema []tasks.Schema `json:"schema" dc:"主体结构"`
	// SampleTextUrl 示例原文的url
	SampleTextUrl string `json:"sampleTextUrl" dc:"示例原文的下载link"`
	// SampleXlsxUrl 示例抽取结果的url
	SampleXlsxUrl string `json:"sampleXlsxUrl" dc:"示例抽取结果的下载link"`

	// ExtractConfig 抽取配置
	ExtractConfig *py_service.ExtractConfig `json:"extractConfig" dc:"抽取配置"`
}

type UpdateProjectReq struct {
	g.Meta          `path:"projects/{id}" method:"put" sm:"更新项目" tags:"项目管理"`
	Id              int     `json:"id" v:"required|min:1" dc:"项目ID"`
	ProjectName     *string `json:"projectName" v:"required|length:1,200" dc:"项目名称"`
	ProjectProgress *int    `json:"projectProgress" v:"min:0|max:100" dc:"项目进度(0-100)"`
	Visibility      *int    `json:"visibility" dc:"项目可见性, 0-private, 1-public"`
}

type UpdateProjectRes struct {
	Success bool `json:"success" dc:"更新是否成功"`
}

type UploadSchemaReq struct {
	g.Meta    `path:"upload/domain" method:"post" sm:"上传本体结构" tags:"文件上传"`
	ProjectId int            `json:"projectId" v:"required|min:1" dc:"项目ID"`
	Triples   []tasks.Schema `json:"triples" dc:"主体结构"`
}
type UploadSchemaRes struct {
}

type UploadSampleTextReq struct {
	g.Meta    `path:"upload/sample" method:"post" sm:"上传样本文本" tags:"文件上传"`
	ProjectId int    `json:"projectId" v:"required|min:1" dc:"项目ID"`
	FilePath  string `json:"filePath" v:"required" dc:"文件路径（存储桶中的位置）"`
}
type UploadSampleTextRes struct {
}

type UploadSampleXLSXReq struct {
	g.Meta    `path:"upload/xlsx" method:"post" sm:"上传样本XLSX" tags:"文件上传"`
	ProjectId int    `json:"projectId" v:"required|min:1" dc:"项目ID"`
	FilePath  string `json:"filePath" v:"required" dc:"文件路径（存储桶中的位置）"`
}

type UploadSampleXLSXRes struct {
}

type ListProjectReq struct {
	g.Meta     `path:"projects/list" method:"post" sm:"获取项目列表" tags:"项目管理"`
	Page       utils.IPage `json:"page" dc:"分页参数"`
	Visibility int         `json:"visibility" dc:"项目可见性, 0-private, 1-public"`
}

type ListProjectRes struct {
	Total    int               `json:"total"`
	Projects []entity.Projects `json:"projects"`
}

type UploadProjectSnapshotPhotoReq struct {
	g.Meta    `path:"projects/uploadPhoto" method:"post" sm:"上传项目的图谱缩略图" tags:"项目管理"`
	ProjectId int    `json:"projectId" v:"required|min:1" dc:"项目ID"`
	FilePath  string `json:"filePath" v:"required" dc:"文件路径"`
}
type UploadProjectSnapshotPhotoRes struct {
}

type GetTriplesTypeByProjectReq struct {
	g.Meta    `path:"projects/getTriplesType" method:"post" sm:"获取项目下的三元组类型" tags:"项目管理"`
	ProjectId int `json:"projectId" v:"required|min:1" dc:"项目ID"`
}

type GetTriplesTypeByProjectRes struct {
	TripleTypeEntity []neo4j.TripleTypeEntity `json:"tripleTypes"`
}

// GetTriplesByTypeReq 按照三元组类型获取项目下的三元组
type GetTriplesByTypeReq struct {
	g.Meta    `path:"projects/getTriplesByType" method:"post" sm:"按照三元组类型获取项目下的三元组" tags:"项目管理"`
	ProjectId int `json:"projectId" v:"required|min:1" dc:"项目ID"`
	// type的形式为 head.type-relationship.type-tail.type
	TripleType *neo4j.TripleTypeEntity `json:"tripleType" v:"required" dc:"三元组类型"`
}
type GetTriplesByTypeRes struct {
	Triples     []neo4j.SimpleTripleVO           `json:"triples" dc:"该type的三元组"`
	Percentages []projects.MaterialToPercentages `json:"percentages" dc:"该type的各个materials中三元组所占比例, key:materials, value:百分比"`
}

// GetProjectEntitiesReq 获取项目下的实体结构
type GetProjectEntitiesReq struct {
	g.Meta    `path:"projects/getEntities" method:"post" sm:"获取项目下的实体结构" tags:"项目管理"`
	ProjectId int `json:"projectId" v:"required|min:1" dc:"项目ID"`
}
type GetProjectEntitiesRes struct {
	TripleEntities map[string][]string `json:"tripleEntities"`
}

// ExportTriplesToZipReq 导出项目下的三元组到Excel
type ExportTriplesToZipReq struct {
	g.Meta    `path:"projects/exportTriplesToZip" method:"post" sm:"导出项目下的三元组到zip, 内含三元组按照type分类的excel" tags:"项目管理"`
	ProjectId int `json:"projectId" v:"required|min:1" dc:"项目ID"`
}
type ExportTriplesToZipRes struct {
}

type PublishProjectReq struct {
	g.Meta     `path:"projects/publish" method:"post" sm:"发布项目" tags:"项目管理"`
	ProjectId  int  `json:"projectId" v:"required|min:1" dc:"项目ID"`
	Visibility *int `json:"visibility" dc:"可见性, 0-private, 1-public"`
	//Price            int     `json:"price" v:"required" dc:"价格"`
	ProjectName      *string   `json:"projectName" dc:"项目名称"`
	SnapshotPhotoURL *string   `json:"snapshotPhotoURL" dc:"项目缩略图/封面"`
	ReadPriceCent    string    `json:"readPriceCent" dc:"阅读价格(元)"`
	BuyPriceCent     string    `json:"buyPriceCent" dc:"购买价格(元)"`
	Description      *string   `json:"description" dc:"详细描述"`
	Tags             *[]string `json:"tags" dc:"关键词数组"`
}

type PublishProjectRes struct {
}

type ListExtractPriorityReq struct {
	g.Meta    `path:"projects/listExtractPriority" method:"post" sm:"获取项目下的抽取优先级" tags:"项目管理"`
	ProjectId int `json:"projectId" v:"required|min:1" dc:"项目ID"`
}
type ListExtractPriorityRes struct {
	EntityList       []string `json:"entityList" dc:"实体类型"`
	RelationshipList []string `json:"relationshipList" dc:"关系类型"`
}

// AddTripleToProjectReq 添加三元组到项目
type AddTripleToProjectReq struct {
	g.Meta     `path:"projects/addTriples" method:"post" sm:"添加三元组到项目，triple为新增的三元组" tags:"项目管理"`
	ProjectId  int                    `json:"projectId" v:"required|min:1" dc:"项目ID"`
	TripleType neo4j.TripleTypeEntity `json:"tripleType" v:"required" dc:"三元组类型"`
	Triples    []neo4j.SimpleTriple   `json:"triples" v:"required" dc:"三元组列表"`
}
type AddTripleToProjectRes struct {
}

// 购买项目请求
type PurchaseProjectReq struct {
	g.Meta      `path:"projects/purchase" method:"post" sm:"购买项目" tags:"项目管理"`
	ProjectId   int    `json:"projectId" v:"required|min:1" dc:"项目ID"`
	ProjectName string `json:"projectName" v:"required|length:1,200" dc:"项目名称"`
	PayType     string `json:"pay_type" v:"required|in:web,wap,app" dc:"支付类型: web-电脑网站, wap-手机网站, app-APP"`
}

// 购买项目响应
type PurchaseProjectRes struct {
	OrderString string `json:"order_string" dc:"支付订单信息(HTML或URL)"`
	OutTradeNo  string `json:"out_trade_no" dc:"商户订单号"`
	TotalAmount string `json:"total_amount" dc:"支付金额"`
}

// 查询项目购买状态请求
type QueryProjectPurchaseStatusReq struct {
	g.Meta     `path:"projects/purchase/status" method:"get" sm:"查询项目购买状态" tags:"项目管理"`
	ProjectId  int    `json:"projectId" v:"required|min:1" dc:"项目ID"`
	OutTradeNo string `json:"out_trade_no" v:"required" dc:"商户订单号"`
}

// 查询项目购买状态响应
type QueryProjectPurchaseStatusRes struct {
	IsPurchased bool   `json:"is_purchased" dc:"是否已购买"`
	TradeStatus string `json:"trade_status" dc:"交易状态"`
	Message     string `json:"message" dc:"状态描述"`
}

// DownloadExtractExampleReq 下载抽取样例文件
type DownloadExtractExampleReq struct {
	g.Meta `path:"projects/download/downloadExtractExample" method:"post" sm:"下载抽取样例文件" tags:"项目管理"`
}
type DownloadExtractExampleRes struct {
}

type GetProjectPurchaseCountReq struct {
	g.Meta    `path:"/purchase/project/:project_id" method:"get" tags:"项目" summary:"获取项目购买次数"`
	ProjectId int `json:"project_id" v:"required" dc:"项目ID"`
}

type GetProjectPurchaseCountRes struct {
	PurchaseCount int `json:"purchase_count" dc:"购买次数"`
}

// GetProjectFavorCountReq 获取项目收藏数请求
type GetProjectFavorCountReq struct {
	g.Meta    `path:"/projects/favor/count/:project_id" method:"get" tags:"项目" summary:"获取项目收藏数"`
	ProjectId int `json:"project_id" v:"required" dc:"项目ID"`
}

// GetProjectFavorCountRes 获取项目收藏数响应
type GetProjectFavorCountRes struct {
	FavorCount int `json:"favor_count" dc:"收藏数"`
}

type GetProjectViewCountReq struct {
	g.Meta    `path:"/view/project/:project_id" method:"get" tags:"项目" summary:"获取项目浏览量"`
	ProjectId int `json:"project_id" v:"required" dc:"项目ID"`
}

type GetProjectViewCountRes struct {
	ViewCount int `json:"view_count" dc:"浏览量"`
}

type GetTripleSourceInfoReq struct {
	g.Meta    `path:"projects/{projectId}/triplets/source" method:"post" sm:"获取三元组来源信息" tags:"项目管理"`
	ProjectId int                `json:"projectId" v:"required" dc:"项目ID"`
	Triple    neo4j.SimpleTriple `json:"triple" v:"required" dc:"三元组"`
}

type GetTripleSourceInfoRes struct {
	MaterialId   int    `json:"materialId" dc:"素材ID"`
	MaterialName string `json:"materialName" dc:"素材名称"`
	SourceText   string `json:"sourceText" dc:"来源文本片段"`
	ChunkIndex   int    `json:"chunkIndex" dc:"Chunk索引"`
}

// GetPublicProjectsReq 不需要登录的公开项目列表请求参数
type GetPublicProjectsReq struct {
	g.Meta `path:"projects/public/list" method:"get" sm:"获取所有公开项目列表" tags:"项目管理"`
	Page   utils.IPage `json:"page" dc:"分页参数"`
}

// GetPublicProjectsRes 不需要登录的公开项目列表响应参数
type GetPublicProjectsRes struct {
	Total    int               `json:"total"`
	Projects []entity.Projects `json:"projects"`
}
