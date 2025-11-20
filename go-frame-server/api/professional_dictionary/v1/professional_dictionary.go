package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"kgplatform-backend/internal/model/entity"
	"kgplatform-backend/internal/utils"
)

// 创建专业词典
type CreateProfessionalDictionaryReq struct {
	g.Meta `path:"/professional_dictionary/create" method:"post" tags:"配置-专业词典" sm:"新增专业词典"`
	Name   string `json:"name" v:"required|length:1,255" dc:"词典名称"`
	Url    string `json:"url"  v:"required|length:1,2000" dc:"词典下载地址或存储路径"`
}
type CreateProfessionalDictionaryRes struct {
	Id        int         `json:"id"`
	CreatedAt *gtime.Time `json:"createdAt"`
}

// 专业词典列表（分页）
type ListProfessionalDictionaryReq struct {
	g.Meta `path:"/professional_dictionary/list" method:"post" tags:"配置-专业词典" sm:"专业词典列表"`
	Page   utils.IPage `json:"page"`
}
type ListProfessionalDictionaryRes struct {
	Total        int                             `json:"total"`
	Dictionaries []entity.ProfessionalDictionary `json:"dictionaries"`
}
