package projects

import (
	"context"
	v1 "kgplatform-backend/api/projects/v1"
	projectsLogic "kgplatform-backend/internal/logic/projects"
	"kgplatform-backend/internal/service"
)

// PurchaseProject 购买项目接口
func (c *ControllerV1) PurchaseProject(ctx context.Context, req *v1.PurchaseProjectReq) (res *v1.PurchaseProjectRes, err error) {
	// 调用逻辑层的BuyProject方法创建支付订单
	output, err := c.projects.BuyProject(ctx, &projectsLogic.BuyProjectInput{
		ProjectId: req.ProjectId,
	})
	if err != nil {
		return nil, err
	}

	// 转换逻辑层输出为API响应
	return &v1.PurchaseProjectRes{
		OrderString: output.OrderString,
		OutTradeNo:  output.OutTradeNo,
		TotalAmount: output.TotalAmount,
	}, nil
}

// QueryProjectPurchaseStatus 查询项目购买状态接口
func (c *ControllerV1) QueryProjectPurchaseStatus(ctx context.Context, req *v1.QueryProjectPurchaseStatusReq) (res *v1.QueryProjectPurchaseStatusRes, err error) {
	// 1. 调用支付宝服务查询订单状态
	alipayRes, err := service.Alipay().QueryOrder(ctx, req.OutTradeNo)
	if err != nil {
		return &v1.QueryProjectPurchaseStatusRes{
			IsPurchased: false,
			TradeStatus: "",
			Message:     "查询订单失败: " + err.Error(),
		}, nil
	}

	// 2. 判断交易状态是否为支付成功
	isPurchased := false
	message := "订单未支付"

	// 支付宝交易状态说明：
	// - WAIT_BUYER_PAY：等待买家付款
	// - TRADE_CLOSED：交易关闭或已退款
	// - TRADE_SUCCESS：支付成功
	// - TRADE_FINISHED：交易结束，不可退款

	if alipayRes.TradeStatus == "TRADE_SUCCESS" || alipayRes.TradeStatus == "TRADE_FINISHED" {
		isPurchased = true
		message = "支付成功"
	} else if alipayRes.TradeStatus == "TRADE_CLOSED" {
		message = "交易已关闭"
	} else if alipayRes.TradeStatus == "WAIT_BUYER_PAY" {
		message = "等待支付"
	}

	// 3. 返回查询结果
	return &v1.QueryProjectPurchaseStatusRes{
		IsPurchased: isPurchased,
		TradeStatus: alipayRes.TradeStatus,
		Message:     message,
	}, nil
}
