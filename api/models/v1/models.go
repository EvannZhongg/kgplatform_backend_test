package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"kgplatform-backend/internal/model/entity"
	"kgplatform-backend/internal/utils"
)

// 创建模型
type CreateModelReq struct {
	g.Meta    `path:"/models/create" method:"post" tags:"配置-模型" sm:"新增支持模型"`
	Provider  string  `json:"provider"  v:"required|length:1,100" dc:"提供方，如 OpenAI/DeepSeek"`
	ModelCode string  `json:"modelCode" v:"required|length:1,150" dc:"模型编码，如 gpt-4o 或 DeepSeek-V3-0324"`
	Name      string  `json:"name"      v:"required|length:1,255" dc:"展示名称"`
	Status    *int    `json:"status;omitempty" dc:"状态：1启用，0禁用"`
	Desc      *string `json:"description;omitempty" dc:"说明"`
}
type CreateModelRes struct {
	Id        int         `json:"id"`
	CreatedAt *gtime.Time `json:"createdAt"`
}

// 模型列表（分页）
type ListModelsReq struct {
	g.Meta   `path:"/models/list" method:"post" tags:"配置-模型" sm:"支持模型列表"`
	Page     utils.IPage `json:"page"`
	Provider *string     `json:"provider;omitempty"`
	Status   *int        `json:"status;omitempty"`
}
type ListModelsRes struct {
	Total  int             `json:"total"`
	Models []entity.Models `json:"models"`
}
