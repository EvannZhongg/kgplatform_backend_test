package cron

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

// NewSyncViewCountJob 定义一个“浏览量同步任务”
func NewSyncViewCountJob(ctx context.Context) CronJob {
	return CronJob{
		Name:    "SyncViewCountJob",
		Pattern: "*/5 * * * *", // 每5分钟执行一次
		Function: func(ctx context.Context) {
			SyncViewCountToDB(ctx)
		},
	}
}

// SyncViewCountToDB 同步Redis中的浏览量到数据库
func SyncViewCountToDB(ctx context.Context) {
	g.Log().Info(ctx, "开始同步浏览量到数据库...")

	pattern := "project:view:count:*"
	keys, err := g.Redis().Keys(ctx, pattern)
	if err != nil {
		g.Log().Error(ctx, "获取Redis keys失败:", err)
		return
	}

	if len(keys) == 0 {
		g.Log().Info(ctx, "没有需要同步的浏览量数据")
		return
	}

	syncCount := 0
	for _, key := range keys {
		var projectId int
		_, err := fmt.Sscanf(key, "project:view:count:%d", &projectId)
		if err != nil {
			continue
		}

		// 使用 GetDel 保证原子操作
		count, err := g.Redis().GetDel(ctx, key)
		if err != nil || count.Int() == 0 {
			continue
		}

		// 增量更新数据库中的 view_count 字段
		_, err = g.Model("projects").
			Where("id", projectId).
			Increment("view_count", count.Int())
		if err != nil {
			g.Log().Errorf(ctx, "更新项目 %d 浏览量失败: %v", projectId, err)
			// 如果失败，放回 Redis，防止数据丢失
			g.Redis().IncrBy(ctx, key, count.Int64())
		} else {
			syncCount++
			g.Log().Debugf(ctx, "项目 %d 浏览量已同步: +%d", projectId, count.Int())
		}
	}

	g.Log().Infof(ctx, "浏览量同步完成，共同步 %d 个项目", syncCount)

	// 同步浏览日志
	syncViewLogs(ctx)
}

// syncViewLogs 同步浏览日志到数据库
func syncViewLogs(ctx context.Context) {
	viewLogKey := "project:view:logs"
	batchSize := 1000

	for {
		logs, err := g.Redis().LRange(ctx, viewLogKey, 0, int64(batchSize-1))
		if err != nil || len(logs) == 0 {
			break
		}

		var records []g.Map
		for _, log := range logs {
			var logData g.Map
			if err := gjson.Unmarshal([]byte(log.String()), &logData); err == nil {
				records = append(records, logData)
			}
		}

		if len(records) > 0 {
			_, err = g.Model("project_views").Data(records).Insert()
			if err != nil {
				g.Log().Error(ctx, "批量插入浏览日志失败:", err)
				break
			}

			g.Redis().LTrim(ctx, viewLogKey, int64(len(logs)), -1)
			g.Log().Infof(ctx, "已同步 %d 条浏览日志", len(records))
		}

		if len(logs) < batchSize {
			break
		}
	}
}
