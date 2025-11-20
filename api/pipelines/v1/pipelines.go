package v1

import (
	"kgplatform-backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type CreatePipelineReq struct {
	g.Meta    `path:"/pipeline/create" method:"post" tags:"工作流" sm:"创建工作流"`
	ProjectId int `json:"projectId" v:"required#请选择项目"`
}

type CreatePipelineRes struct {
	Id        int         `json:"id" dc:"id"`
	ProjectId int         `json:"projectId" dc:"项目id"`
	CreatedAt *gtime.Time `json:"created" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated" dc:"更新时间"`
}

type GetPipelineReq struct {
	g.Meta    `path:"/pipeline/get/{projectId}" method:"get" tags:"工作流" sm:"获取工作流"`
	ProjectId int `path:"projectId" v:"required#请选择工作流"`
}
type GetPipelineRes struct {
	Id          int           `json:"id" dc:"id"`
	ProjectId   int           `json:"projectId" dc:"项目id"`
	OCRTask     *entity.Tasks `json:"ocrTask" dc:"ocr任务"`
	ExtractTask *entity.Tasks `json:"extractTask" dc:"提取任务"`
	GraphTask   *entity.Tasks `json:"graphTask" dc:"图谱化任务"`
	CreatedAt   *gtime.Time   `json:"created" dc:"创建时间"`
	UpdatedAt   *gtime.Time   `json:"updated" dc:"更新时间"`
}

// LikePipelineReq 点赞工作流请求
type LikePipelineReq struct {
	g.Meta     `path:"/pipeline/:pipeline_id/like" method:"post" tags:"工作流" sm:"点赞工作流"`
	PipelineId int `json:"pipeline_id" v:"required" dc:"工作流ID"`
}

type LikePipelineRes struct {
	Message string `json:"message"`
}

// UnlikePipelineReq 取消点赞工作流请求
type UnlikePipelineReq struct {
	g.Meta     `path:"/pipeline/:pipeline_id/like" method:"delete" tags:"工作流" sm:"取消点赞工作流"`
	PipelineId int `json:"pipeline_id" v:"required" dc:"工作流ID"`
}

type UnlikePipelineRes struct {
	Message string `json:"message"`
}

// GetPipelineLikesReq 获取工作流点赞数请求
type GetPipelineLikesReq struct {
	g.Meta     `path:"/pipeline/:pipeline_id/likes" method:"get" tags:"工作流" sm:"获取工作流点赞数"`
	PipelineId int `json:"pipeline_id" v:"required" dc:"工作流ID"`
}

type GetPipelineLikesRes struct {
	LikeCount int  `json:"like_count" dc:"点赞数"`
	IsLiked   bool `json:"is_liked" dc:"当前用户是否点赞"`
}

// FavorPipelineReq 收藏工作流请求
type FavorPipelineReq struct {
	g.Meta     `path:"/pipeline/:pipeline_id/favor" method:"post" tags:"工作流" sm:"收藏工作流"`
	PipelineId int `json:"pipeline_id" v:"required" dc:"工作流ID"`
}

type FavorPipelineRes struct {
	Message string `json:"message"`
}

// UnfavorPipelineReq 取消收藏工作流请求
type UnfavorPipelineReq struct {
	g.Meta     `path:"/pipeline/:pipeline_id/favor" method:"delete" tags:"工作流" sm:"取消收藏工作流"`
	PipelineId int `json:"pipeline_id" v:"required" dc:"工作流ID"`
}

type UnfavorPipelineRes struct {
	Message string `json:"message"`
}

// GetPipelineFavorsReq 获取工作流收藏数请求
type GetPipelineFavorsReq struct {
	g.Meta     `path:"/pipeline/:pipeline_id/favors" method:"get" tags:"工作流" sm:"获取工作流收藏数"`
	PipelineId int `json:"pipeline_id" v:"required" dc:"工作流ID"`
}

type GetPipelineFavorsRes struct {
	FavorCount int  `json:"favor_count" dc:"收藏数"`
	IsFavored  bool `json:"is_favored" dc:"当前用户是否收藏"`
}
