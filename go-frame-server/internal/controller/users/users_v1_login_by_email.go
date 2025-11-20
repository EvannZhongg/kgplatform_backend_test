package users

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	v1 "kgplatform-backend/api/users/v1"
	"time"
)

func (c *ControllerV1) LoginByEmail(ctx context.Context, req *v1.LoginByEmailReq) (res *v1.LoginByEmailRes, err error) {
	token, err := c.users.LoginByEmail(ctx, req.Email, req.Code)
	if err != nil {
		return nil, err
	}

	if req.RememberMe {
		r := ghttp.RequestFromCtx(ctx)
		r.Cookie.SetCookie("email", req.Email, "", "/", 30*24*time.Hour)
	}

	return &v1.LoginByEmailRes{
		Token: token,
	}, nil
}
