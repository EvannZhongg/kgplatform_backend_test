// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package alipay

import (
	"context"

	"kgplatform-backend/api/alipay/v1"
)

type IAlipayV1 interface {
	CreatePay(ctx context.Context, req *v1.CreatePayReq) (res *v1.CreatePayRes, err error)
	Notify(ctx context.Context, req *v1.NotifyReq) (res *v1.NotifyRes, err error)
	QueryOrder(ctx context.Context, req *v1.QueryOrderReq) (res *v1.QueryOrderRes, err error)
	Refund(ctx context.Context, req *v1.RefundReq) (res *v1.RefundRes, err error)
	CreateBillingPay(ctx context.Context, req *v1.CreateBillingPayReq) (res *v1.CreateBillingPayRes, err error)
	CreateSubscriptionPay(ctx context.Context, req *v1.CreateSubscriptionPayReq) (res *v1.CreateSubscriptionPayRes, err error)
}
