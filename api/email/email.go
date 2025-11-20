// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package email

import (
	"context"

	"kgplatform-backend/api/email/v1"
)

type IEmailV1 interface {
	SendEmail(ctx context.Context, req *v1.SendEmailReq) (res *v1.SendEmailRes, err error)
	VerifyEmail(ctx context.Context, req *v1.VerifyEmailReq) (res *v1.VerifyEmailRes, err error)
}
