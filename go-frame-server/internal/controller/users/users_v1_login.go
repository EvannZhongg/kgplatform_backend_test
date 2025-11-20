package users

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"time"

	v1 "kgplatform-backend/api/users/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	token, err := c.users.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	if req.RememberMe {
		r := ghttp.RequestFromCtx(ctx)
		r.Cookie.SetCookie("username", req.Username, "", "/", 30*24*time.Hour)
		r.Cookie.SetCookie("password", req.Password, "", "/", 30*24*time.Hour)
	}

	return &v1.LoginRes{
		Token: token,
	}, nil
}
