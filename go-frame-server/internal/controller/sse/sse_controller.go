package sse

import (
	"fmt"
	"kgplatform-backend/internal/service"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type ControllerV1 struct{}

func NewV1() *ControllerV1 {
	return &ControllerV1{}
}

// HandleSSE 处理用户流量用量SSE连接
func (c *ControllerV1) HandleSSE(r *ghttp.Request) {
	userIdVar := r.Get("user_id")
	userId := userIdVar.Int64() // 转成 int64
	if userId <= 0 {
		r.Response.WriteStatus(400, "user_id 必填或无效")
		return
	}

	//r.Response.CORSDefault()
	ctx := r.GetCtx()
	g.Log().Info(ctx, "设置SSE响应头")
	// 设置SSE响应头
	r.Response.Header().Set("Content-Type", "text/event-stream")
	r.Response.Header().Set("Cache-Control", "no-cache")
	r.Response.Header().Set("Connection", "keep-alive")

	// 注册SSE连接
	g.Log().Info(ctx, "注册SSE连接")
	msgChan := service.Manager.Register(userId)
	defer service.Manager.Unregister(userId)

	// 心跳定时器，防止中间代理关闭连接
	ticker := time.NewTicker(30 * time.Second)
	g.Log().Info(ctx, "开始心脏跳动")
	defer ticker.Stop()

	for {
		select {
		case msg, ok := <-msgChan:
			if !ok {
				return
			}
			// 发送SSE消息
			fmt.Fprintf(r.Response.Writer, "data: %s\n\n", msg)
			r.Response.Flush()
		case <-ticker.C:
			// 发送心跳
			fmt.Fprintf(r.Response.Writer, ": keep-alive\n\n")
			r.Response.Flush()
		case <-r.Context().Done():
			return
		}
	}
}
