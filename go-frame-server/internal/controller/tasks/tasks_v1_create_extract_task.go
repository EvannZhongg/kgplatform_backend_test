package tasks

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"kgplatform-backend/api/tasks/v1"
	"kgplatform-backend/external/py_service"
	"kgplatform-backend/internal/consts"
	"kgplatform-backend/internal/logic/projects"
	"kgplatform-backend/internal/logic/tasks"
)

func (c *ControllerV1) CreateExtractTask(ctx context.Context, req *v1.CreateExtractTaskReq) (res *v1.CreateExtractTaskRes, err error) {
	// 保存抽取配置
	projectLogic := projects.NewProjects()
	extractConfig := py_service.ExtractConfig{
		Prompt:                 req.Prompt,
		ModelId:                *req.ModelId,
		Method:                 req.Method,
		SchemaURL:              req.SchemaURL,
		SampleTextURL:          req.SampleTextURL,
		SampleXLSXURL:          req.SampleXLSXURL,
		TargetDomain:           req.TargetDomain,
		DictionaryId:           req.DictionaryId,
		DictionaryURL:          req.DictionaryURL,
		PriorityExtractions:    req.PriorityExtractions,
		ExtractionRequirements: req.ExtractionRequirements,
	}
	err = projectLogic.UpdateProjectByGMap(ctx, req.ProjectId, g.Map{
		"extract_config": gjson.New(extractConfig),
	})
	if err != nil {
		return nil, gerror.New("抽取配置保存失败")
	}

	// 进行抽取
	taskLogic := tasks.New()
	input := &tasks.CreateTaskInput{
		Type:           consts.TaskTypeExtract,
		PipelineId:     req.PipelineId,
		ProjectId:      req.ProjectId,
		MaterialIdList: req.MaterialIDList,
		Status:         consts.TaskStatusPending,
		UpdatedAt:      gtime.Now(),
		CreatedAt:      gtime.Now(),
		Prompt:         req.Prompt,
		Method:         req.Method,
		ModelId:        req.ModelId,
	}
	taskEntity, err := taskLogic.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	res = &v1.CreateExtractTaskRes{
		Status: taskEntity.Status,
		TaskId: taskEntity.Id,
	}
	return res, nil
}
