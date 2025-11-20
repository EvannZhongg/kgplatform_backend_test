package likes

import (
	"context"
	v1 "kgplatform-backend/api/likes/v1"
	"kgplatform-backend/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	//"github.com/gogf/gf/v2/frame/g"
)

type CLike struct{}

func NewV1() *CLike {
	return &CLike{}
}

//func (c *CLike) getUserIdFromToken(ctx context.Context) (int, error) {
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

func (c *CLike) LikeProject(ctx context.Context, req *v1.LikeProjectReq) (res *v1.LikeProjectRes, err error) {
	// 从上下文获取当前登录用户ID
	userId := g.RequestFromCtx(ctx).GetCtxVar("userID").Int()
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	err = service.Like().LikeProject(ctx, userId, req.ProjectId)
	if err != nil {
		return nil, err
	}

	return &v1.LikeProjectRes{
		Message: "点赞成功",
	}, nil
}

func (c *CLike) UnlikeProject(ctx context.Context, req *v1.UnlikeProjectReq) (res *v1.UnlikeProjectRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userID").Int()
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	err = service.Like().UnlikeProject(ctx, userId, req.ProjectId)
	if err != nil {
		return nil, err
	}

	return &v1.UnlikeProjectRes{
		Message: "取消点赞成功",
	}, nil
}

func (c *CLike) GetProjectLikes(ctx context.Context, req *v1.GetProjectLikesReq) (res *v1.GetProjectLikesRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userID").Int()
	if userId == 0 {
		userId = 0
	}

	// 检查项目可见性
	visibility, err := g.Model("projects").
		Where("id", req.ProjectId).
		Fields("visibility").
		Value()
	if err != nil {
		return nil, err
	}

	// 如果项目不存在或visibility不为1，返回错误
	if visibility.Int() != 1 {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "项目不可见或不存在")
	}

	likeCount, isLiked, err := service.Like().GetProjectLikes(ctx, req.ProjectId, userId)
	if err != nil {
		return nil, err
	}

	return &v1.GetProjectLikesRes{
		LikeCount: likeCount,
		IsLiked:   isLiked,
	}, nil
}
