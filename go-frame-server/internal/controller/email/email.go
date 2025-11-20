package email

import (
	"context"

	v1 "kgplatform-backend/api/email/v1"
	"kgplatform-backend/internal/logic/email"
)

type EmailV1 struct{}

func NewV1() *EmailV1 {
	return &EmailV1{}
}

func (c *EmailV1) Send(ctx context.Context, req *v1.SendEmailReq) (res *v1.SendEmailRes, err error) {
	emailLogic := email.New()
	err = emailLogic.SendVerificationCode(ctx, req.Email)
	if err != nil {
		return &v1.SendEmailRes{
			Message: "发送失败: " + err.Error(),
		}, err
	}

	return &v1.SendEmailRes{
		Message: "验证码发送成功",
	}, nil
}

func (c *EmailV1) Verify(ctx context.Context, req *v1.VerifyEmailReq) (res *v1.VerifyEmailRes, err error) {
	emailLogic := email.New()
	valid, err := emailLogic.VerifyCode(ctx, req.Email, req.Code)
	if err != nil {
		return &v1.VerifyEmailRes{
			Valid:   false,
			Message: "验证失败: " + err.Error(),
		}, err
	}

	if valid {
		return &v1.VerifyEmailRes{
			Valid:   true,
			Message: "验证成功",
		}, nil
	}

	return &v1.VerifyEmailRes{
		Valid:   false,
		Message: "验证码错误或已过期",
	}, nil
}
