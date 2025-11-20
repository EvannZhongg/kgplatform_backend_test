package v1

import "github.com/gogf/gf/v2/frame/g"

// 创建支付订单请求
type CreatePayReq struct {
	g.Meta      `path:"/alipay/create" method:"post" tags:"支付宝支付" summary:"创建支付订单"`
	Subject     string  `json:"subject" v:"required#商品名称不能为空"`
	OutTradeNo  string  `json:"out_trade_no" v:"required#订单号不能为空"`
	TotalAmount float64 `json:"total_amount" v:"required|min:0.01#金额不能为空|金额必须大于0.01"`
	Body        string  `json:"body"`
	PayType     string  `json:"pay_type" v:"required|in:web,wap,app#支付类型不能为空|支付类型错误"` // web:电脑网站, wap:手机网站, app:APP
}

type CreatePayRes struct {
	OrderString string `json:"order_string"` // 支付宝订单信息(HTML或URL)
	OutTradeNo  string `json:"out_trade_no"` // 商户订单号
}

// 支付回调通知
type NotifyReq struct {
	g.Meta `path:"/alipay/notify" method:"post" tags:"支付宝支付" summary:"支付回调通知"`
}

type NotifyRes struct {
	Result string `json:"result"`
}

// 查询订单
type QueryOrderReq struct {
	g.Meta     `path:"/alipay/query" method:"get" tags:"支付宝支付" summary:"查询订单"`
	OutTradeNo string `json:"out_trade_no" v:"required#订单号不能为空"`
}

type QueryOrderRes struct {
	TradeNo     string `json:"trade_no"`      // 支付宝交易号
	OutTradeNo  string `json:"out_trade_no"`  // 商户订单号
	TradeStatus string `json:"trade_status"`  // 交易状态
	TotalAmount string `json:"total_amount"`  // 订单金额
	BuyerUserId string `json:"buyer_user_id"` // 买家用户ID
}

// 退款请求
type RefundReq struct {
	g.Meta       `path:"/alipay/refund" method:"post" tags:"支付宝支付" summary:"申请退款"`
	OutTradeNo   string  `json:"out_trade_no" v:"required#订单号不能为空"`
	RefundAmount float64 `json:"refund_amount" v:"required|min:0.01#退款金额不能为空|退款金额必须大于0.01"`
	RefundReason string  `json:"refund_reason"`
}

type RefundRes struct {
	TradeNo      string `json:"trade_no"`      // 支付宝交易号
	OutTradeNo   string `json:"out_trade_no"`  // 商户订单号
	RefundAmount string `json:"refund_amount"` // 退款金额
	RefundStatus string `json:"refund_status"` // 退款状态
}

// 创建基于账单的支付订单
type CreateBillingPayReq struct {
	g.Meta    `path:"/alipay/create-billing-pay" method:"post" tags:"支付宝支付" summary:"根据账单创建支付订单"`
	BillingId int64  `json:"billing_id" v:"required#账单ID不能为空"`
	PayType   string `json:"pay_type" v:"required|in:web,wap,app#支付类型不能为空|支付类型错误"` // web:电脑网站, wap:手机网站, app:APP
	Remark    string `json:"remark"`
}

type CreateBillingPayRes struct {
	OrderString string `json:"order_string"` // 支付宝订单信息(HTML或URL)
	OutTradeNo  string `json:"out_trade_no"` // 商户订单号
	TotalAmount string `json:"total_amount"` // 支付金额
}

// 创建订阅套餐支付订单
type CreateSubscriptionPayReq struct {
	g.Meta      `path:"/alipay/create-subscription-pay" method:"post" tags:"支付宝支付" summary:"创建订阅套餐支付订单"`
	UserPlan    string `json:"user_plan" v:"required|in:professional,team#套餐类型不能为空|套餐类型错误"`
	MemberCount int    `json:"member_count" v:"required-if:UserPlan,team|min:3|max:100#团队版需要指定人数|团队人数不能少于3人|团队人数不能超过100人"`
	PayType     string `json:"pay_type" v:"required|in:web,wap,app#支付类型不能为空|支付类型错误"` // web:电脑网站, wap:手机网站, app:APP
}

type CreateSubscriptionPayRes struct {
	OrderString string `json:"order_string"` // 支付宝订单信息(HTML或URL)
	OutTradeNo  string `json:"out_trade_no"` // 商户订单号
	TotalAmount string `json:"total_amount"` // 支付金额
}
