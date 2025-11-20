package projects

import (
	"context"
	v1 "kgplatform-backend/api/projects/v1"
)

type IProjectsV1 interface {
	CreateProject(ctx context.Context, req *v1.CreateProjectReq) (res *v1.CreateProjectRes, err error)
	GetProject(ctx context.Context, req *v1.GetProjectReq) (res *v1.GetProjectRes, err error)
	UpdateProject(ctx context.Context, req *v1.UpdateProjectReq) (res *v1.UpdateProjectRes, err error)
	UploadSchema(ctx context.Context, req *v1.UploadSchemaReq) (res *v1.UploadSchemaRes, err error)
	UploadSampleText(ctx context.Context, req *v1.UploadSampleTextReq) (res *v1.UploadSampleTextRes, err error)
	UploadSampleXLSX(ctx context.Context, req *v1.UploadSampleXLSXReq) (res *v1.UploadSampleXLSXRes, err error)

	ListProject(ctx context.Context, req *v1.ListProjectReq) (res *v1.ListProjectRes, err error)

	UploadProjectSnapshotPhoto(ctx context.Context, req *v1.UploadProjectSnapshotPhotoReq) (res *v1.UploadProjectSnapshotPhotoRes, err error)

	GetTriplesTypeByProject(ctx context.Context, req *v1.GetTriplesTypeByProjectReq) (res *v1.GetTriplesTypeByProjectRes, err error)

	GetTriplesByType(ctx context.Context, req *v1.GetTriplesByTypeReq) (res *v1.GetTriplesByTypeRes, err error)

	GetProjectEntities(ctx context.Context, req *v1.GetProjectEntitiesReq) (res *v1.GetProjectEntitiesRes, err error)

	ExportTriplesToZip(ctx context.Context, req *v1.ExportTriplesToZipReq) (res *v1.ExportTriplesToZipRes, err error)

	PublishProject(ctx context.Context, req *v1.PublishProjectReq) (res *v1.PublishProjectRes, err error)

	// ListExtractPriority 获取抽取意向优先级
	ListExtractPriority(ctx context.Context, req *v1.ListExtractPriorityReq) (res *v1.ListExtractPriorityRes, err error)

	// UpdateTripleToProject 修改 or 新增三元组
	AddTripleToProject(ctx context.Context, req *v1.AddTripleToProjectReq) (res *v1.AddTripleToProjectRes, err error)

	// 购买项目
	PurchaseProject(ctx context.Context, req *v1.PurchaseProjectReq) (res *v1.PurchaseProjectRes, err error)

	// 查询项目购买状态
	QueryProjectPurchaseStatus(ctx context.Context, req *v1.QueryProjectPurchaseStatusReq) (res *v1.QueryProjectPurchaseStatusRes, err error)

	// 获取项目收藏数
	GetProjectFavorCount(ctx context.Context, req *v1.GetProjectFavorCountReq) (res *v1.GetProjectFavorCountRes, err error)

	// 下载抽取示例文件
	DownloadExtractExample(ctx context.Context, req *v1.DownloadExtractExampleReq) (res *v1.DownloadExtractExampleRes, err error)

	// 获取三元组来源信息
	GetTripleSourceInfo(ctx context.Context, req *v1.GetTripleSourceInfoReq) (res *v1.GetTripleSourceInfoRes, err error)
}
