// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package account

import (
	"context"

	"kgplatform-backend/api/account/v1"
)

type IAccountV1 interface {
	Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error)
}
