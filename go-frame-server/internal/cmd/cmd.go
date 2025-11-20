package cmd

import (
	"context"
	"kgplatform-backend/internal/controller/account"
	"kgplatform-backend/internal/controller/alipay"
	"kgplatform-backend/internal/controller/chat"
	"kgplatform-backend/internal/controller/comments"
	"kgplatform-backend/internal/controller/email"
	"kgplatform-backend/internal/controller/graphs"
	"kgplatform-backend/internal/controller/likes"
	"kgplatform-backend/internal/controller/materials"
	"kgplatform-backend/internal/controller/models"
	"kgplatform-backend/internal/controller/pipelines"
	"kgplatform-backend/internal/controller/professional_dictionary"
	"kgplatform-backend/internal/controller/projects"
	"kgplatform-backend/internal/controller/sms"
	"kgplatform-backend/internal/controller/sse"
	"kgplatform-backend/internal/controller/support_domains"
	"kgplatform-backend/internal/controller/tasks"
	"kgplatform-backend/internal/controller/teams"
	"kgplatform-backend/internal/controller/upload"
	"kgplatform-backend/internal/controller/users"
	"kgplatform-backend/internal/logic/middleware"
	"kgplatform-backend/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	tasks2 "kgplatform-backend/internal/logic/tasks"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化配额重置定时任务
			service.UserSubscription.InitCronTask(ctx)

			s := g.Server()

			// 配置静态文件服务
			s.AddStaticPath("/uploads", "./resource/public/uploads")

			s.Group("/", func(group *ghttp.RouterGroup) {
				// 添加CORS中间件
				group.Middleware(func(r *ghttp.Request) {
					r.Response.CORSDefault()
					origin := r.Header.Get("Origin")
					if origin != "" {
						r.Response.Header().Set("Access-Control-Allow-Origin", origin)
					}
					r.Response.Header().Set("Access-Control-Allow-Credentials", "true")
					r.Response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
					r.Response.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")

					r.Middleware.Next()
				})
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/v1", func(group *ghttp.RouterGroup) {
					group.Bind(
						users.NewV1(),
						sms.NewV1(),
						email.NewV1(),
						alipay.NewV1Public(),
						projects.NewV1Public(),
					)
					// SSE发送流量预警通知，供前端调用
					group.Group("/usage", func(group *ghttp.RouterGroup) {
						group.GET("/stream", sse.NewV1().HandleSSE)
					})
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.Auth)
						group.Bind(
							account.NewV1(),
							upload.NewV1(),
							materials.NewV1(),
							projects.NewV1(),
							comments.NewV1(),
							likes.NewV1(),
							teams.NewV1(),
							alipay.NewV1(),
							tasks.NewV1(),
							chat.NewV1(),
							pipelines.NewV1(),
							models.NewV1(),
							support_domains.NewV1(),
							professional_dictionary.NewV1(),
						)
						group.Group("/", func(graphGroup *ghttp.RouterGroup) {
							graphGroup.Middleware(middleware.TrafficStats("graph_query"))
							graphGroup.Bind(graphs.NewV1())
						})
					})
				})
			})

			s.Run()

			return nil
		},
	}
	AsynQCmd = gcmd.Command{
		Name:  "asynq",
		Usage: "asynq",
		Brief: "start asynq server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			tasks2.GetTaskManager()
			return nil
		},
	}
)
