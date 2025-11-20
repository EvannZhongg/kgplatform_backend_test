package cron

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

// NewTmpFileCleanJob 初始化临时文件清理任务
func NewTmpFileCleanJob(ctx context.Context) *CronJob {
	// 从配置中读取清理设置
	enabled := g.Cfg().MustGet(ctx, "download.cleanup.enabled", true).Bool()
	if !enabled {
		glog.Info(ctx, "临时文件清理任务未启用")
		return nil
	}

	// 获取下载路径
	downloadPath := g.Cfg().MustGet(ctx, "download.path", "./resource/public/downloads").String()

	// 获取清理时间表达式
	cronExpr := g.Cfg().MustGet(ctx, "download.cleanup.cron", "0 0 3 * * *").String()

	// 获取文件最大保留时间（秒）
	maxAge := g.Cfg().MustGet(ctx, "download.cleanup.maxAge", 86400).Int() // 默认24小时

	// 添加定时任务
	return &CronJob{
		Name:    "tmp_file_clean",
		Pattern: cronExpr,
		Function: func(ctx context.Context) {
			cleanTmpFiles(ctx, downloadPath, maxAge)
		},
	}
}

// cleanTmpFiles 清理临时文件
func cleanTmpFiles(ctx context.Context, downloadPath string, maxAge int) {
	logger := g.Log()

	// 确保目录存在
	if _, err := os.Stat(downloadPath); os.IsNotExist(err) {
		logger.Info(ctx, "下载目录不存在，无需清理:", downloadPath)
		return
	}

	// 遍历目录中的文件
	err := filepath.Walk(downloadPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.Error(ctx, "访问文件时出错:", err)
			return nil
		}

		// 跳过目录
		if info.IsDir() {
			return nil
		}

		// 检查文件是否超过最大保留时间
		if time.Since(info.ModTime()) > time.Duration(maxAge)*time.Second {
			// 删除文件
			if err := os.Remove(path); err != nil {
				logger.Error(ctx, "删除文件失败:", path, err)
			} else {
				logger.Info(ctx, "已删除过期文件:", path)
			}
		}

		return nil
	})

	if err != nil {
		logger.Error(ctx, "遍历目录时出错:", err)
	} else {
		logger.Info(ctx, "临时文件清理完成")
	}
}
