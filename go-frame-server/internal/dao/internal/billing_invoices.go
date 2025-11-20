// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BillingInvoicesDao is the data access object for the table billing_invoices.
type BillingInvoicesDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  BillingInvoicesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// BillingInvoicesColumns defines and stores column names for the table billing_invoices.
type BillingInvoicesColumns struct {
	Id              string //
	BillingId       string //
	InvoiceRequired string // 是否需要发票
	InvoiceTitle    string // 发票抬头
	InvoiceTaxId    string // 税号
	InvoiceUrl      string // 发票下载链接
	InvoiceIssuedAt string // 发票开具时间
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
}

// billingInvoicesColumns holds the columns for the table billing_invoices.
var billingInvoicesColumns = BillingInvoicesColumns{
	Id:              "id",
	BillingId:       "billing_id",
	InvoiceRequired: "invoice_required",
	InvoiceTitle:    "invoice_title",
	InvoiceTaxId:    "invoice_tax_id",
	InvoiceUrl:      "invoice_url",
	InvoiceIssuedAt: "invoice_issued_at",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewBillingInvoicesDao creates and returns a new DAO object for table data access.
func NewBillingInvoicesDao(handlers ...gdb.ModelHandler) *BillingInvoicesDao {
	return &BillingInvoicesDao{
		group:    "default",
		table:    "billing_invoices",
		columns:  billingInvoicesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BillingInvoicesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BillingInvoicesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BillingInvoicesDao) Columns() BillingInvoicesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BillingInvoicesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BillingInvoicesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BillingInvoicesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
