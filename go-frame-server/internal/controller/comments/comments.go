package comments

import (
	"context"
	v1 "kgplatform-backend/api/comments/v1"
	"kgplatform-backend/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type CComment struct{}

func NewV1() *CComment {
	return &CComment{}
}

//func (c *CComment) getUserIdFromToken(ctx context.Context) (int, error) {
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

func (c *CComment) CreateComment(ctx context.Context, req *v1.CreateCommentReq) (res *v1.CreateCommentRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userID").Int()
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	commentId, err := service.Comment().CreateComment(ctx, userId, req.ProjectId, req.ParentId, req.Content)
	if err != nil {
		return nil, err
	}

	return &v1.CreateCommentRes{
		CommentId: commentId,
		Message:   "评论成功",
	}, nil
}

func (c *CComment) DeleteComment(ctx context.Context, req *v1.DeleteCommentReq) (res *v1.DeleteCommentRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userID").Int()
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	err = service.Comment().DeleteComment(ctx, userId, req.CommentId)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteCommentRes{
		Message: "删除成功",
	}, nil
}

func (c *CComment) GetComments(ctx context.Context, req *v1.GetCommentsReq) (res *v1.GetCommentsRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userID").Int()
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	list, total, err := service.Comment().GetComments(ctx, req.ProjectId, userId, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	return &v1.GetCommentsRes{
		List:  list,
		Total: total,
		Page:  req.Page,
	}, nil
}

func (c *CComment) LikeComment(ctx context.Context, req *v1.LikeCommentReq) (res *v1.LikeCommentRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userID").Int()
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	err = service.Comment().LikeComment(ctx, userId, req.CommentId)
	if err != nil {
		return nil, err
	}

	return &v1.LikeCommentRes{
		Message: "点赞成功",
	}, nil
}

func (c *CComment) UnlikeComment(ctx context.Context, req *v1.UnlikeCommentReq) (res *v1.UnlikeCommentRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userID").Int()
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	err = service.Comment().UnlikeComment(ctx, userId, req.CommentId)
	if err != nil {
		return nil, err
	}

	return &v1.UnlikeCommentRes{
		Message: "取消点赞成功",
	}, nil
}

func (c *CComment) GetCommentLikes(ctx context.Context, req *v1.GetCommentLikesReq) (res *v1.GetCommentLikesRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userID").Int()
	if userId == 0 {
		// 即使未登录也可以获取点赞数，只是isLiked为false
		userId = 0
	}

	likeCount, isLiked, err := service.Comment().GetCommentLikes(ctx, req.CommentId, userId)
	if err != nil {
		return nil, err
	}

	return &v1.GetCommentLikesRes{
		LikeCount: likeCount,
		IsLiked:   isLiked,
	}, nil
}
