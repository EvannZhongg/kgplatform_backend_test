package projects

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"kgplatform-backend/api/projects/v1"
)

func (c *ControllerV1) UploadSampleText(ctx context.Context, req *v1.UploadSampleTextReq) (res *v1.UploadSampleTextRes, err error) {
	err = c.projects.UpdateProjectByGMap(ctx, req.ProjectId, g.Map{
		"sample_text_url": req.FilePath,
	})
	if err != nil {
		return nil, gerror.New("上传失败")
	}
	return &v1.UploadSampleTextRes{}, nil
}
