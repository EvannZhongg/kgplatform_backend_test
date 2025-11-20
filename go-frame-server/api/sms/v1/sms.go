package v1

import "github.com/gogf/gf/v2/frame/g"

type SendSmsReq struct {
	g.Meta `path:"sms/send" method:"post" sm:"发送短信验证码" tags:"短信"`
	Phone  string `json:"phone" v:"required|phone" dc:"手机号码"`
}

type SendSmsRes struct {
	Message string `json:"message" dc:"响应消息"`
}

type VerifySmsReq struct {
	g.Meta `path:"sms/verify" method:"post" sm:"验证短信验证码" tags:"短信"`
	Phone  string `json:"phone" v:"required|phone" dc:"手机号码"`
	Code   string `json:"code" v:"required|length:6,6" dc:"验证码"`
}

type VerifySmsRes struct {
	Valid   bool   `json:"valid" dc:"验证是否通过"`
	Message string `json:"message" dc:"响应消息"`
}
