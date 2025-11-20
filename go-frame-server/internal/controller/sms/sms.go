package sms

import (
	"context"

	v1 "kgplatform-backend/api/sms/v1"
	"kgplatform-backend/internal/logic/sms"
)

type V1 struct{}

func NewV1() *V1 {
	return &V1{}
}

func (c *V1) SendSms(ctx context.Context, req *v1.SendSmsReq) (res *v1.SendSmsRes, err error) {
	smsLogic := sms.New()
	err = smsLogic.SendVerificationCode(ctx, req.Phone)
	if err != nil {
		return &v1.SendSmsRes{
			Message: "发送失败: " + err.Error(),
		}, err
	}

	return &v1.SendSmsRes{
		Message: "验证码发送成功",
	}, nil
}

func (c *V1) VerifySms(ctx context.Context, req *v1.VerifySmsReq) (res *v1.VerifySmsRes, err error) {
	smsLogic := sms.New()
	valid, err := smsLogic.VerifyCode(ctx, req.Phone, req.Code)
	if err != nil {
		return &v1.VerifySmsRes{
			Valid:   false,
			Message: "验证失败: " + err.Error(),
		}, err
	}

	if valid {
		return &v1.VerifySmsRes{
			Valid:   true,
			Message: "验证成功",
		}, nil
	}

	return &v1.VerifySmsRes{
		Valid:   false,
		Message: "验证码错误或已过期",
	}, nil
}
