package projects

import (
	"context"
	"kgplatform-backend/internal/logic/projects"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"kgplatform-backend/api/projects/v1"
)

func (c *ControllerV1) PublishProject(ctx context.Context, req *v1.PublishProjectReq) (res *v1.PublishProjectRes, err error) {
	projectLogic := projects.NewProjects()
	err = projectLogic.UpdateProjectByGMap(ctx, req.ProjectId, g.Map{
		"visibility":         req.Visibility,
		"description":        req.Description,
		"tags":               req.Tags,
		"name":               req.ProjectName,
		"snapshot_photo_url": req.SnapshotPhotoURL,
		"buy_price_cent":     req.BuyPriceCent,
		"read_price_cent":    req.ReadPriceCent,
		"updated_at":         gtime.Now(),
	})

	if err != nil {
		return nil, gerror.New("发布失败")
	}
	return &v1.PublishProjectRes{}, nil
}
