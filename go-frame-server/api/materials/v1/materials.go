package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"kgplatform-backend/internal/logic/materials"
	"kgplatform-backend/internal/utils"
)

type CreateMaterialReq struct {
	g.Meta    `path:"materials/create" method:"post" sm:"创建材料" tags:"材料管理"`
	FilePath  string `json:"filePath" v:"required" dc:"文件路径（存储桶中的位置）"`
	ProjectId int    `json:"projectId" v:"required|min:1" dc:"项目ID"`
}

type CreateMaterialRes struct {
	MaterialId int    `json:"materialId" dc:"材料ID"`
	Status     string `json:"status" dc:"状态"`
}

type GetMaterialReq struct {
	g.Meta `path:"materials/{id}" method:"get" sm:"获取材料详情" tags:"材料管理"`
	Id     int `json:"id" v:"required|min:1" dc:"材料ID"`
}

type GetMaterialRes struct {
	Id        int    `json:"id" dc:"材料ID"`
	FileUrl   string `json:"fileUrl" dc:"文件下载链接"`
	OcrUrl    string `json:"ocrUrl" dc:"文字化结果下载链接"`
	TripleUrl string `json:"tripleUrl" dc:"三元组抽取结果下载链接"`
	ProjectId int    `json:"projectId" dc:"项目ID"`
	CreatedAt string `json:"createdAt" dc:"创建时间"`
	UpdatedAt string `json:"updatedAt" dc:"更新时间"`
}

type UpdateMaterialReq struct {
	g.Meta     `path:"materials/{id}" method:"put" sm:"更新材料" tags:"材料管理"`
	Id         int    `json:"id" v:"required|min:1" dc:"材料ID"`
	FilePath   string `json:"filePath" dc:"文件路径（存储桶中的位置）"`
	OcrPath    string `json:"ocrPath" dc:"OCR文字化结果文件路径"`
	TriplePath string `json:"triplePath" dc:"三元组抽取结果文件路径"`
	ProjectId  int    `json:"projectId" v:"min:1" dc:"项目ID"`
}

type UpdateMaterialRes struct {
	Success bool `json:"success" dc:"更新是否成功"`
}

type DeleteMaterialReq struct {
	g.Meta `path:"materials/{id}" method:"delete" sm:"删除材料" tags:"材料管理"`
	Id     int `json:"id" v:"required|min:1" dc:"材料ID"`
}

type DeleteMaterialRes struct {
	Success bool `json:"success" dc:"删除是否成功"`
}

type ListMaterialReq struct {
	g.Meta    `path:"materials" method:"get" sm:"获取材料列表" tags:"材料管理"`
	Page      utils.IPage `json:"page" dc:"分页"`
	ProjectId int         `json:"projectId" v:"min:1" dc:"项目ID"`
}

type ListMaterialRes struct {
	Total    int                         `json:"total" dc:"总记录数"`
	Material []*materials.MaterialDetail `json:"material" dc:"材料列表"`
}

type UpdateMaterialsOCRTextReq struct {
	g.Meta   `path:"materials/updateOCRText" method:"post" sm:"更新材料OCR文字" tags:"材料管理"`
	Id       int    `json:"id" v:"required|min:1" dc:"材料ID"`
	FilePath string `json:"filePath" v:"required" dc:"文件名称"`
}
type UpdateMaterialsOCRTextRes struct {
}

// SwitchMaterialStatusReq 切换材料状态
type SwitchMaterialStatusReq struct {
	g.Meta `path:"materials/enableOrUnable" method:"post" sm:"切换材料状态" tags:"材料管理"`
	Id     int `json:"id" v:"required|min:1" dc:"材料ID"`
}
type SwitchMaterialStatusRes struct {
}
