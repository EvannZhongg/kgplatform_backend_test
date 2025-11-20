package projects

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"

	"kgplatform-backend/api/projects/v1"
)

func (c *ControllerV1) UploadSampleXLSX(ctx context.Context, req *v1.UploadSampleXLSXReq) (res *v1.UploadSampleXLSXRes, err error) {
	err = c.projects.UpdateProjectByGMap(ctx, req.ProjectId, g.Map{
		"sample_xlsx_url": req.FilePath,
	})
	if err != nil {
		return nil, gerror.New("上传失败")
	}
	return &v1.UploadSampleXLSXRes{}, nil
}
