package projects

import (
	"context"
	v1 "kgplatform-backend/api/projects/v1"
	"kgplatform-backend/internal/logic/projects"
)

// ControllerV1Public 项目公开控制器，不需要登录认证
// 用于提供不需要登录即可访问的项目相关接口

type ControllerV1Public struct {
	projects *projects.Projects
}

// NewV1Public 创建项目公开控制器实例
func NewV1Public() *ControllerV1Public {
	return &ControllerV1Public{
		projects: projects.NewProjects(),
	}
}

// GetPublicProjects 获取所有公开项目列表
// 这个接口不需要用户登录认证
func (c *ControllerV1Public) GetPublicProjects(ctx context.Context, req *v1.GetPublicProjectsReq) (res *v1.GetPublicProjectsRes, err error) {
	projectOutput, err := c.projects.GetPublicProjects(ctx, &projects.GetPublicProjectsInput{
		Page: req.Page,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GetPublicProjectsRes{
		Total:    projectOutput.Total,
		Projects: projectOutput.Projects,
	}, nil
}
