package pipelines

import (
	"context"
	"kgplatform-backend/internal/consts"
	"kgplatform-backend/internal/logic/pipelines"
	"kgplatform-backend/internal/logic/tasks"
	"kgplatform-backend/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"

	"kgplatform-backend/api/pipelines/v1"
)

func (c *ControllerV1) GetPipeline(ctx context.Context, req *v1.GetPipelineReq) (res *v1.GetPipelineRes, err error) {
	pipelineLogic := pipelines.NewPipelines()
	pipeline, err := pipelineLogic.GetPipeline(ctx, &pipelines.GetPipelineInput{
		ProjectId: req.ProjectId,
	})
	if err != nil {
		return nil, gerror.New("获取工作流失败")
	}
	taskLogic := tasks.New()
	taskListOutput, err := taskLogic.ListTask(ctx, &tasks.ListTaskInput{
		PipelineId: &pipeline.Id,
	})
	if err != nil {
		return nil, gerror.New("获取任务列表失败")
	}
	res = &v1.GetPipelineRes{
		Id:        pipeline.Id,
		ProjectId: pipeline.ProjectId,
		CreatedAt: pipeline.CreatedAt,
		UpdatedAt: pipeline.UpdatedAt,
	}
	// 为每种任务类型保留最新创建的任务
	taskMap := make(map[string]*entity.Tasks)
	for _, task := range taskListOutput.Tasks {
		if existingTask, exists := taskMap[task.Type]; !exists || task.CreatedAt.After(existingTask.CreatedAt) {
			taskMap[task.Type] = task
		}
	}

	// 将最新任务分配给对应的字段
	for _, task := range taskMap {
		if task.Type == consts.TaskTypeOCR {
			res.OCRTask = task
		} else if task.Type == consts.TaskTypeExtract {
			res.ExtractTask = task
		} else if task.Type == consts.TaskTypeGraph {
			res.GraphTask = task
		}
	}
	return res, nil
}
