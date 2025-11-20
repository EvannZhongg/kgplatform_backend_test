package main

import (
	"fmt"
	_ "kgplatform-backend/internal/logic"
	"kgplatform-backend/internal/neo4j"

	_ "kgplatform-backend/internal/packed"
	"time"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"kgplatform-backend/internal/cmd"
	corn "kgplatform-backend/internal/corn"
)

func main() {
	var err error

	// 全局设置i18n
	g.I18n().SetLanguage("zh-CN")
	err = connDb()
	if err != nil {
		panic(err)
	}

	// 初始化Neo4j驱动
	err = neo4j.InitDriver()
	if err != nil {
		panic(fmt.Errorf("初始化Neo4j驱动失败: %v", err))
	}
	defer neo4j.CloseDriver()

	cmd.AsynQCmd.Run(gctx.GetInitCtx())
	cmd.Main.Run(gctx.GetInitCtx())

	err = corn.RegisterCronJobs(gctx.GetInitCtx())
	if err != nil {
		panic(fmt.Errorf("定时任务加载失败: %v", err))
	}
}

func connDb() error {
	// 重试连接数据库
	maxRetries := 10
	retryDelay := 2 * time.Second

	for i := 0; i < maxRetries; i++ {
		err := g.DB().PingMaster()
		if err == nil {
			return nil
		}

		if i < maxRetries-1 {
			fmt.Printf("数据库连接失败，%v后重试... (尝试 %d/%d)\n", retryDelay, i+1, maxRetries)
			time.Sleep(retryDelay)
		}
	}

	return fmt.Errorf("数据库连接失败，已重试%d次", maxRetries)
}
