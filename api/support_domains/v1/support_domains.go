package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"kgplatform-backend/internal/model/entity"
	"kgplatform-backend/internal/utils"
)

// 创建领域
type CreateDomainReq struct {
	g.Meta      `path:"/support_domains/create" method:"post" tags:"配置-领域" sm:"新增目标领域"`
	DisplayName string  `json:"displayName" v:"required|length:1,255" dc:"展示名称"`
	Status      *int    `json:"status;omitempty" dc:"状态：1启用，0禁用"`
	Desc        *string `json:"description;omitempty" dc:"说明"`
}
type CreateDomainRes struct {
	Id        int         `json:"id"`
	CreatedAt *gtime.Time `json:"createdAt"`
}

// 领域列表（分页）
type ListDomainsReq struct {
	g.Meta `path:"/support_domains/list" method:"post" tags:"配置-领域" sm:"目标领域列表"`
	Page   utils.IPage `json:"page"`
	Status *int        `json:"status;omitempty"`
}
type ListDomainsRes struct {
	Total   int                     `json:"total"`
	Domains []entity.SupportDomains `json:"support_domains"`
}
