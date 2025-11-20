// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sms

import (
	"context"

	"kgplatform-backend/api/sms/v1"
)

type ISmsV1 interface {
	SendSms(ctx context.Context, req *v1.SendSmsReq) (res *v1.SendSmsRes, err error)
	VerifySms(ctx context.Context, req *v1.VerifySmsReq) (res *v1.VerifySmsRes, err error)
}
