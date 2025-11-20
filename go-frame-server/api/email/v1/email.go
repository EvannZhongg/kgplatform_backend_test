package v1

import "github.com/gogf/gf/v2/frame/g"

type SendEmailReq struct {
	g.Meta `path:"email/send" method:"post" sm:"发送邮箱验证码" tags:"邮箱"`
	Email  string `json:"email" v:"required|email" dc:"邮箱地址"`
}

type SendEmailRes struct {
	Message string `json:"message" dc:"响应消息"`
}

type VerifyEmailReq struct {
	g.Meta `path:"email/verify" method:"post" sm:"验证邮箱验证码" tags:"邮箱"`
	Email  string `json:"email" v:"required|email" dc:"邮箱地址"`
	Code   string `json:"code" v:"required|length:6,6" dc:"验证码"`
}

type VerifyEmailRes struct {
	Valid   bool   `json:"valid" dc:"验证是否通过"`
	Message string `json:"message" dc:"响应消息"`
}
