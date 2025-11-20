package v1

import "github.com/gogf/gf/v2/frame/g"

// CreateCommentReq 创建评论请求
type CreateCommentReq struct {
	g.Meta    `path:"/comments" method:"post" tags:"评论" summary:"创建评论"`
	ProjectId int    `json:"project_id" v:"required" dc:"项目ID"`
	ParentId  *int   `json:"parent_id" dc:"父评论ID（回复时填写）"`
	Content   string `json:"content" v:"required|length:1,1000" dc:"评论内容"`
}

type CreateCommentRes struct {
	CommentId int    `json:"comment_id"`
	Message   string `json:"message"`
}

// DeleteCommentReq 删除评论请求
type DeleteCommentReq struct {
	g.Meta    `path:"/comments/:comment_id" method:"delete" tags:"评论" summary:"删除评论"`
	CommentId int `json:"comment_id" v:"required" dc:"评论ID"`
}

type DeleteCommentRes struct {
	Message string `json:"message"`
}

// GetCommentsReq 获取项目评论列表请求
type GetCommentsReq struct {
	g.Meta    `path:"/comments/project/:project_id" method:"get" tags:"评论" summary:"获取项目评论列表"`
	ProjectId int `json:"project_id" v:"required" dc:"项目ID"`
	Page      int `json:"page" d:"1" dc:"页码"`
	PageSize  int `json:"page_size" d:"20" dc:"每页数量"`
}

type CommentItem struct {
	Id        int           `json:"id"`
	ProjectId int           `json:"project_id"`
	UserId    int           `json:"user_id"`
	Username  string        `json:"username"`
	ParentId  *int          `json:"parent_id"`
	Content   string        `json:"content"`
	LikeCount int           `json:"like_count"`
	IsLiked   bool          `json:"is_liked"`
	CreatedAt string        `json:"created_at"`
	Replies   []CommentItem `json:"replies,omitempty"`
}

type GetCommentsRes struct {
	List  []CommentItem `json:"list"`
	Total int           `json:"total"`
	Page  int           `json:"page"`
}

// LikeCommentReq 点赞评论请求
type LikeCommentReq struct {
	g.Meta    `path:"/comments/:comment_id/like" method:"post" tags:"评论" summary:"点赞评论"`
	CommentId int `json:"comment_id" v:"required" dc:"评论ID"`
}

type LikeCommentRes struct {
	Message string `json:"message"`
}

// UnlikeCommentReq 取消点赞评论请求
type UnlikeCommentReq struct {
	g.Meta    `path:"/comments/:comment_id/like" method:"delete" tags:"评论" summary:"取消点赞评论"`
	CommentId int `json:"comment_id" v:"required" dc:"评论ID"`
}

type UnlikeCommentRes struct {
	Message string `json:"message"`
}

// GetCommentLikesReq 获取评论点赞数请求
type GetCommentLikesReq struct {
	g.Meta    `path:"/comments/:comment_id/likes" method:"get" tags:"评论" summary:"获取评论点赞数"`
	CommentId int `json:"comment_id" v:"required" dc:"评论ID"`
}

type GetCommentLikesRes struct {
	LikeCount int  `json:"like_count" dc:"点赞数"`
	IsLiked   bool `json:"is_liked" dc:"当前用户是否点赞"`
}
