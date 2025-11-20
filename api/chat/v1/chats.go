package v1

import "github.com/gogf/gf/v2/frame/g"

type GeneratePromptReq struct {
	g.Meta                 `path:"chat/generatePrompt" method:"post" sm:"生成提示" tags:"任务管理"`
	ProjectId              int       `json:"projectId" v:"required|min:1" dc:"项目ID"`
	DictionaryId           *int      `json:"dictionaryId;omitempty" dc:"使用的专业词典"`
	TargetDomain           *string   `json:"targetDomain;omitempty" dc:"目标领域, 如建筑学/医学.."`
	PriorityExtractions    *[]string `json:"priorityExtractions;omitempty" dc:"抽取意向优先级"`
	ExtractionRequirements *string   `json:"extractionRequirements;omitempty" dc:"抽取要求描述文本"`
	BaseInstruction        *string   `json:"baseInstruction;omitempty" dc:"自定义基础指导语"`
}

type GeneratePromptRes struct {
	Prompt                 string   `json:"prompt" dc:"提示"`
	SchemaURL              string   `json:"schemaUrl"`
	SampleTextURL          string   `json:"sampleTextUrl"`
	SampleXLSXURL          string   `json:"sampleXlsxUrl"`
	TargetDomain           string   `json:"targetDomain"`
	DictionaryId           *int     `json:"dictionaryId"`
	DictionaryURL          string   `json:"dictionaryUrl"`
	PriorityExtractions    []string `json:"priorityExtractions"`
	ExtractionRequirements string   `json:"extractionRequirements"`
}
