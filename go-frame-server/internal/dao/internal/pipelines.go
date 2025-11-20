// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PipelinesDao is the data access object for the table pipelines.
type PipelinesDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PipelinesColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PipelinesColumns defines and stores column names for the table pipelines.
type PipelinesColumns struct {
	Id          string //
	StartStep   string // 起始步骤
	ProjectId   string //
	CreatedAt   string //
	UpdatedAt   string //
	LikesCount  string //
	FavorsCount string //
	ViewsCount  string //
}

// pipelinesColumns holds the columns for the table pipelines.
var pipelinesColumns = PipelinesColumns{
	Id:          "id",
	StartStep:   "start_step",
	ProjectId:   "project_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	LikesCount:  "likes_count",
	FavorsCount: "favors_count",
	ViewsCount:  "views_count",
}

// NewPipelinesDao creates and returns a new DAO object for table data access.
func NewPipelinesDao(handlers ...gdb.ModelHandler) *PipelinesDao {
	return &PipelinesDao{
		group:    "default",
		table:    "pipelines",
		columns:  pipelinesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PipelinesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PipelinesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PipelinesDao) Columns() PipelinesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PipelinesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PipelinesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PipelinesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
