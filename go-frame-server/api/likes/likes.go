// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package likes

import (
	"context"

	"kgplatform-backend/api/likes/v1"
)

type ILikesV1 interface {
	LikeProject(ctx context.Context, req *v1.LikeProjectReq) (res *v1.LikeProjectRes, err error)
	UnlikeProject(ctx context.Context, req *v1.UnlikeProjectReq) (res *v1.UnlikeProjectRes, err error)
	GetProjectLikes(ctx context.Context, req *v1.GetProjectLikesReq) (res *v1.GetProjectLikesRes, err error)
}
