package users

import (
	"context"
	v1 "kgplatform-backend/api/users/v1"
	"kgplatform-backend/internal/logic/users"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	err = c.users.Register(ctx, &users.RegisterInput{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
		Email:    req.Email,
	})
	return nil, err
}
