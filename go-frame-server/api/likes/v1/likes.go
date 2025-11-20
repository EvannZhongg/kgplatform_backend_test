package v1

import "github.com/gogf/gf/v2/frame/g"

// LikeProjectReq 点赞项目请求
type LikeProjectReq struct {
	g.Meta    `path:"/likes/project/:project_id" method:"post" tags:"点赞" summary:"点赞项目"`
	ProjectId int `json:"project_id" v:"required" dc:"项目ID"`
}

type LikeProjectRes struct {
	Message string `json:"message"`
}

// UnlikeProjectReq 取消点赞项目请求
type UnlikeProjectReq struct {
	g.Meta    `path:"/likes/project/:project_id" method:"delete" tags:"点赞" summary:"取消点赞项目"`
	ProjectId int `json:"project_id" v:"required" dc:"项目ID"`
}

type UnlikeProjectRes struct {
	Message string `json:"message"`
}

// GetProjectLikesReq 获取项目点赞数请求
type GetProjectLikesReq struct {
	g.Meta    `path:"/likes/project/:project_id" method:"get" tags:"点赞" summary:"获取项目点赞数"`
	ProjectId int `json:"project_id" v:"required" dc:"项目ID"`
}

type GetProjectLikesRes struct {
	LikeCount int  `json:"like_count" dc:"点赞数"`
	IsLiked   bool `json:"is_liked" dc:"当前用户是否点赞"`
}
