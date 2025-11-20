// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ModelsDao is the data access object for the table models.
type ModelsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ModelsColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ModelsColumns defines and stores column names for the table models.
type ModelsColumns struct {
	Id          string //
	Provider    string // 模型提供方
	ModelCode   string // 模型编码
	Name        string // 展示名称
	Status      string // 状态：1启用，0禁用
	Description string //
	CreatedAt   string //
	UpdatedAt   string //
}

// modelsColumns holds the columns for the table models.
var modelsColumns = ModelsColumns{
	Id:          "id",
	Provider:    "provider",
	ModelCode:   "model_code",
	Name:        "name",
	Status:      "status",
	Description: "description",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewModelsDao creates and returns a new DAO object for table data access.
func NewModelsDao(handlers ...gdb.ModelHandler) *ModelsDao {
	return &ModelsDao{
		group:    "default",
		table:    "models",
		columns:  modelsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ModelsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ModelsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ModelsDao) Columns() ModelsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ModelsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ModelsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ModelsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
