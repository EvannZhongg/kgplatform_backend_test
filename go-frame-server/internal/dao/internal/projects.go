// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProjectsDao is the data access object for the table projects.
type ProjectsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ProjectsColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ProjectsColumns defines and stores column names for the table projects.
type ProjectsColumns struct {
	Id               string //
	UserId           string //
	ProjectName      string // 项目名称
	ProjectProgress  string // 项目进度(0-100)
	GraphId          string //
	CreatedAt        string //
	UpdatedAt        string //
	SchemaUrl        string // 主体结构url
	SampleTextUrl    string // 示例原文的url
	SampleXlsxUrl    string // 示例抽取结果(三元组)的url
	SnapshotPhotoUrl string // 图谱照片url
	Visibility       string // 可见性, 0-private, 1-public
	BuyPriceCent     string //
	ReadPriceCent    string //
	PurchaseCount    string //
	ViewCount        string //
	TripleTypeUrl    string //
	TripleUrl        string //
	Description      string //
	ExtractConfig    string //
}

// projectsColumns holds the columns for the table projects.
var projectsColumns = ProjectsColumns{
	Id:               "id",
	UserId:           "user_id",
	ProjectName:      "project_name",
	ProjectProgress:  "project_progress",
	GraphId:          "graph_id",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	SchemaUrl:        "schema_url",
	SampleTextUrl:    "sample_text_url",
	SampleXlsxUrl:    "sample_xlsx_url",
	SnapshotPhotoUrl: "snapshot_photo_url",
	Visibility:       "visibility",
	BuyPriceCent:     "buy_price_cent",
	ReadPriceCent:    "read_price_cent",
	PurchaseCount:    "purchase_count",
	ViewCount:        "view_count",
	TripleTypeUrl:    "triple_type_url",
	TripleUrl:        "triple_url",
	Description:      "description",
	ExtractConfig:    "extract_config",
}

// NewProjectsDao creates and returns a new DAO object for table data access.
func NewProjectsDao(handlers ...gdb.ModelHandler) *ProjectsDao {
	return &ProjectsDao{
		group:    "default",
		table:    "projects",
		columns:  projectsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ProjectsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ProjectsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ProjectsDao) Columns() ProjectsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ProjectsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ProjectsDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ProjectsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
