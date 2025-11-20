package projects

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"kgplatform-backend/internal/logic/projects"

	"github.com/gogf/gf/v2/errors/gerror"

	"kgplatform-backend/api/projects/v1"
)

func (c *ControllerV1) UploadProjectSnapshotPhoto(ctx context.Context, req *v1.UploadProjectSnapshotPhotoReq) (res *v1.UploadProjectSnapshotPhotoRes, err error) {
	projectLogic := projects.NewProjects()
	project, err := projectLogic.GetProject(ctx, req.ProjectId)
	if err != nil {
		g.Log().Errorf(ctx, "项目不存在: %v", err)
		return nil, gerror.New("项目不存在")
	}
	err = projectLogic.UpdateProjectByGMap(ctx, project.Id, g.Map{
		"snapshot_photo": req.FilePath,
	})
	if err != nil {
		g.Log().Errorf(ctx, "上传失败: %v", err)
		return nil, gerror.New("上传失败")
	}
	return &v1.UploadProjectSnapshotPhotoRes{}, nil
}
