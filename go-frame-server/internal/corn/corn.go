package cron

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/glog"
)

// CronJob 定时任务结构体
type CronJob struct {
	Name     string
	Pattern  string
	Function func(ctx context.Context)
}

// RegisterCronJobs 注册所有定时任务到系统
func RegisterCronJobs(ctx context.Context) error {
	fileCleanJob := NewTmpFileCleanJob(ctx)
	//syncViewCountJob := NewSyncViewCountJob(ctx)
	_, err := gcron.AddSingleton(ctx, fileCleanJob.Pattern, fileCleanJob.Function, fileCleanJob.Name)
	if err != nil {
		glog.Error(ctx, "添加临时文件清理任务失败:", err)
		return err
	}

	//_, err = gcron.AddSingleton(ctx, syncViewCountJob.Pattern, syncViewCountJob.Function, syncViewCountJob.Name)
	//if err != nil {
	//	glog.Error(ctx, "添加浏览量同步任务失败:", err)
	//	return err
	//}

	return nil
}
