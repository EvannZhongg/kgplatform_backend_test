// alipay_v1_public.go - 不需要认证的接口（支付回调）
package alipay

import (
	"context"
	v1 "kgplatform-backend/api/alipay/v1"
	"kgplatform-backend/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type ControllerV1Public struct{}

func NewV1Public() *ControllerV1Public {
	return &ControllerV1Public{}
}

// Notify 支付回调通知（不需要认证）
func (c *ControllerV1Public) Notify(ctx context.Context, req *v1.NotifyReq) (res *v1.NotifyRes, err error) {
	g.Log().Info(ctx, "支付回调通知:", req)
	r := g.RequestFromCtx(ctx)

	// 获取所有POST参数
	params := r.GetFormMap()

	// 验证并处理回调
	err = service.Alipay().HandleNotify(ctx, params)
	if err != nil {
		g.Log().Error(ctx, "支付回调处理失败:", err)
		r.Response.Write("fail")
		return nil, nil
	}

	// 返回success给支付宝,否则支付宝会持续通知
	r.Response.Write("success")
	return nil, nil
}
