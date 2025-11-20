package alipay

import (
	"context"
	v1 "kgplatform-backend/api/alipay/v1"
	"kgplatform-backend/internal/service"
)

type ControllerV1 struct{}

func NewV1() *ControllerV1 {
	return &ControllerV1{}
}

// CreateSubscriptionPay 创建订阅套餐支付订单
func (c *ControllerV1) CreateSubscriptionPay(ctx context.Context, req *v1.CreateSubscriptionPayReq) (res *v1.CreateSubscriptionPayRes, err error) {
	return service.Alipay().CreateSubscriptionPay(ctx, req)
}

// CreatePay 创建支付订单
func (c *ControllerV1) CreatePay(ctx context.Context, req *v1.CreatePayReq) (res *v1.CreatePayRes, err error) {
	orderString, err := service.Alipay().CreatePay(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.CreatePayRes{
		OrderString: orderString,
		OutTradeNo:  req.OutTradeNo,
	}, nil
}

//// Notify 支付回调通知
//func (c *ControllerV1) Notify(ctx context.Context, req *v1.NotifyReq) (res *v1.NotifyRes, err error) {
//	g.Log().Info(ctx, "支付回调通知:", req)
//	r := g.RequestFromCtx(ctx)
//
//	// 获取所有POST参数
//	params := r.GetFormMap()
//
//	// 验证并处理回调
//	err = service.Alipay().HandleNotify(ctx, params)
//	if err != nil {
//		g.Log().Error(ctx, "支付回调处理失败:", err)
//		r.Response.Write("fail")
//		return nil, nil
//	}
//
//	// 返回success给支付宝,否则支付宝会持续通知
//	r.Response.Write("success")
//	return nil, nil
//}

// QueryOrder 查询订单
func (c *ControllerV1) QueryOrder(ctx context.Context, req *v1.QueryOrderReq) (res *v1.QueryOrderRes, err error) {
	return service.Alipay().QueryOrder(ctx, req.OutTradeNo)
}

// Refund 申请退款
func (c *ControllerV1) Refund(ctx context.Context, req *v1.RefundReq) (res *v1.RefundRes, err error) {
	return service.Alipay().Refund(ctx, req)
}

// CreateBillingPay 根据账单创建支付订单
func (c *ControllerV1) CreateBillingPay(ctx context.Context, req *v1.CreateBillingPayReq) (res *v1.CreateBillingPayRes, err error) {
	return service.Alipay().CreateBillingPay(ctx, req)
}
