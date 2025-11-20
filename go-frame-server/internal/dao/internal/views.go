// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ViewsDao is the data access object for the table views.
type ViewsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ViewsColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ViewsColumns defines and stores column names for the table views.
type ViewsColumns struct {
	TableCatalog            string //
	TableSchema             string //
	Id                      string //
	TableName               string //
	ViewDefinition          string //
	UserId                  string //
	ProjectId               string //
	CheckOption             string //
	IsUpdatable             string //
	IsInsertableInto        string //
	CreatedAt               string //
	IsTriggerUpdatable      string //
	IsTriggerDeletable      string //
	UpdatedAt               string //
	IsTriggerInsertableInto string //
}

// viewsColumns holds the columns for the table views.
var viewsColumns = ViewsColumns{
	TableCatalog:            "table_catalog",
	TableSchema:             "table_schema",
	Id:                      "id",
	TableName:               "table_name",
	ViewDefinition:          "view_definition",
	UserId:                  "user_id",
	ProjectId:               "project_id",
	CheckOption:             "check_option",
	IsUpdatable:             "is_updatable",
	IsInsertableInto:        "is_insertable_into",
	CreatedAt:               "created_at",
	IsTriggerUpdatable:      "is_trigger_updatable",
	IsTriggerDeletable:      "is_trigger_deletable",
	UpdatedAt:               "updated_at",
	IsTriggerInsertableInto: "is_trigger_insertable_into",
}

// NewViewsDao creates and returns a new DAO object for table data access.
func NewViewsDao(handlers ...gdb.ModelHandler) *ViewsDao {
	return &ViewsDao{
		group:    "default",
		table:    "views",
		columns:  viewsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ViewsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ViewsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ViewsDao) Columns() ViewsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ViewsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ViewsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ViewsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
