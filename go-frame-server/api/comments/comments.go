// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package comments

import (
	"context"

	"kgplatform-backend/api/comments/v1"
)

type ICommentsV1 interface {
	CreateComment(ctx context.Context, req *v1.CreateCommentReq) (res *v1.CreateCommentRes, err error)
	DeleteComment(ctx context.Context, req *v1.DeleteCommentReq) (res *v1.DeleteCommentRes, err error)
	GetComments(ctx context.Context, req *v1.GetCommentsReq) (res *v1.GetCommentsRes, err error)
	LikeComment(ctx context.Context, req *v1.LikeCommentReq) (res *v1.LikeCommentRes, err error)
	UnlikeComment(ctx context.Context, req *v1.UnlikeCommentReq) (res *v1.UnlikeCommentRes, err error)
	GetCommentLikes(ctx context.Context, req *v1.GetCommentLikesReq) (res *v1.GetCommentLikesRes, err error)
}
