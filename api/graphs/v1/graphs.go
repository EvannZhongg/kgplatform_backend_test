package v1

import "github.com/gogf/gf/v2/frame/g"

type GetGraphReq struct {
	g.Meta    `path:"/graph/get/:projectId" method:"get" tags:"图谱" sm:"获取图谱"`
	ProjectId int `path:"projectId" v:"required#请选择项目"`
}

type GetGraphRes struct {
	Id    int                 `json:"id"`
	Nodes []map[string]string `json:"nodes"`
	Edges []map[string]string `json:"edges"`
}
