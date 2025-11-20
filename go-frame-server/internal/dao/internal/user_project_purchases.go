// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserProjectPurchasesDao is the data access object for the table user_project_purchases.
type UserProjectPurchasesDao struct {
	table    string                      // table is the underlying table name of the DAO.
	group    string                      // group is the database configuration group name of the current DAO.
	columns  UserProjectPurchasesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler          // handlers for customized model modification.
}

// UserProjectPurchasesColumns defines and stores column names for the table user_project_purchases.
type UserProjectPurchasesColumns struct {
	Id            string //
	UserId        string //
	ProjectId     string //
	BillingId     string //
	PaymentId     string //
	PurchasePrice string //
	Status        string //
	CreatedAt     string //
	UpdatedAt     string //
}

// userProjectPurchasesColumns holds the columns for the table user_project_purchases.
var userProjectPurchasesColumns = UserProjectPurchasesColumns{
	Id:            "id",
	UserId:        "user_id",
	ProjectId:     "project_id",
	BillingId:     "billing_id",
	PaymentId:     "payment_id",
	PurchasePrice: "purchase_price",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewUserProjectPurchasesDao creates and returns a new DAO object for table data access.
func NewUserProjectPurchasesDao(handlers ...gdb.ModelHandler) *UserProjectPurchasesDao {
	return &UserProjectPurchasesDao{
		group:    "default",
		table:    "user_project_purchases",
		columns:  userProjectPurchasesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserProjectPurchasesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserProjectPurchasesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserProjectPurchasesDao) Columns() UserProjectPurchasesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserProjectPurchasesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserProjectPurchasesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserProjectPurchasesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
