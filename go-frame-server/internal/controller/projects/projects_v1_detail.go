package projects

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"kgplatform-backend/internal/logic/projects"

	v1 "kgplatform-backend/api/projects/v1"
)

func (c *ControllerV1) GetProject(ctx context.Context, req *v1.GetProjectReq) (res *v1.GetProjectRes, err error) {
	projectDetail, err := c.projects.GetProjectDetail(ctx, &projects.GetProjectDetailInput{
		ProjectId: req.Id,
	})
	if err != nil {
		return nil, gerror.New("获取项目详情失败")
	}

	return &v1.GetProjectRes{
		Id:              projectDetail.Id,
		UserId:          projectDetail.UserId,
		ProjectName:     projectDetail.ProjectName,
		ProjectProgress: projectDetail.ProjectProgress,
		GraphId:         projectDetail.GraphId,
		Materials:       projectDetail.Materials,
		ExtractType:     projectDetail.ExtractType,
		Schema:          projectDetail.Schema,
		SampleTextUrl:   projectDetail.SampleTextUrl,
		SampleXlsxUrl:   projectDetail.SampleXlsxUrl,
		CreatedAt:       projectDetail.CreatedAt,
		UpdatedAt:       projectDetail.UpdatedAt,
		TripleList:      projectDetail.TripleList,
		PipelineId:      projectDetail.PipelineId,
		ExtractConfig:   projectDetail.ExtractConfig,
	}, nil
}
