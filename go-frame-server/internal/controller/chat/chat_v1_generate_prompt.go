package chat

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"kgplatform-backend/external/py_service"
	"kgplatform-backend/internal/logic/chat"
	"kgplatform-backend/internal/logic/projects"

	"github.com/gogf/gf/v2/errors/gerror"

	"kgplatform-backend/api/chat/v1"
)

func (c *ControllerV1) GeneratePrompt(ctx context.Context, req *v1.GeneratePromptReq) (res *v1.GeneratePromptRes, err error) {
	// 检查权限
	projectLogic := projects.NewProjects()
	_, err = projectLogic.GetProject(ctx, req.ProjectId)
	if err != nil {
		g.Log().Errorf(ctx, "项目不存在: %v", err)
		return nil, gerror.New("项目不存在")
	}

	// 生成prompt
	chatLogic := chat.New(py_service.GetPythonClient())
	chatOutput, err := chatLogic.Chat(ctx, &chat.ChatInput{
		ProjectId:              req.ProjectId,
		DictionaryId:           req.DictionaryId,
		TargetDomain:           req.TargetDomain,
		PriorityExtractions:    req.PriorityExtractions,
		ExtractionRequirements: req.ExtractionRequirements,
		BaseInstruction:        req.BaseInstruction,
	})
	if err != nil {
		g.Log().Errorf(ctx, "prompt生成失败: %v", err)
		return nil, err
	}
	return &v1.GeneratePromptRes{
		Prompt:                 chatOutput.Prompt,
		SchemaURL:              chatOutput.SchemaURL,
		SampleXLSXURL:          chatOutput.SampleXLSXURL,
		SampleTextURL:          chatOutput.SampleTextURL,
		DictionaryId:           chatOutput.DictionaryId,
		DictionaryURL:          chatOutput.DictionaryURL,
		ExtractionRequirements: chatOutput.ExtractionRequirements,
		PriorityExtractions:    chatOutput.PriorityExtractions,
		TargetDomain:           chatOutput.TargetDomain,
	}, nil
}
