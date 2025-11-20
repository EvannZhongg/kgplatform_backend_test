package users

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"time"

	v1 "kgplatform-backend/api/users/v1"
)

func (c *ControllerV1) LoginByPhone(ctx context.Context, req *v1.LoginByPhoneReq) (res *v1.LoginByPhoneRes, err error) {
	token, err := c.users.LoginByPhone(ctx, req.Phone, req.Code)
	if err != nil {
		return nil, err
	}

	if req.RememberMe {
		r := ghttp.RequestFromCtx(ctx)
		r.Cookie.SetCookie("phone", req.Phone, "", "/", 30*24*time.Hour)
	}

	return &v1.LoginByPhoneRes{
		Token: token,
	}, nil
}
