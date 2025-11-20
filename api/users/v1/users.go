package v1

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta   `path:"users/register" method:"post" sm:"注册用户" tags:"用户"`
	Username string `json:"username" v:"required|length:3,12" dc:"用户名"`
	Password string `json:"password" v:"required|length:6,16" dc:"密码"`
	Phone    string `json:"phone" v:"required|phone" dc:"手机号码"`
	Email    string `json:"email" v:"required|email" dc:"邮箱"`
}

type RegisterRes struct {
}

type LoginReq struct {
	g.Meta     `path:"users/login" method:"post" sm:"登录" tags:"用户"`
	Username   string `json:"username" v:"required|length:3,12"`
	Password   string `json:"password" v:"required|length:6,16"`
	RememberMe bool   `json:"remember_me" dc:"记住我"`
}

type LoginRes struct {
	Token string `json:"token" dc:"在需要鉴权的接口中header加入Authorization: token"`
}

type LoginByPhoneReq struct {
	g.Meta     `path:"users/login/phone" method:"post" sm:"通过手机号验证码登录" tags:"用户"`
	Phone      string `json:"phone" v:"required|phone" dc:"手机号码"`
	Code       string `json:"code" v:"required|length:6,6" dc:"验证码"`
	RememberMe bool   `json:"remember_me" dc:"记住我"`
}

type LoginByPhoneRes struct {
	Token string `json:"token" dc:"在需要鉴权的接口中header加入Authorization: token"`
}

type LoginByEmailReq struct {
	g.Meta     `path:"users/login/email" method:"post" sm:"通过邮箱验证码登录" tags:"用户"`
	Email      string `json:"email" v:"required|email" dc:"邮箱地址"`
	Code       string `json:"code" v:"required|length:6,6" dc:"验证码"`
	RememberMe bool   `json:"remember_me" dc:"记住我"`
}

type LoginByEmailRes struct {
	Token string `json:"token" dc:"在需要鉴权的接口中header加入Authorization: token"`
}
