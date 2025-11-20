// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BillingPaymentsDao is the data access object for the table billing_payments.
type BillingPaymentsDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  BillingPaymentsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// BillingPaymentsColumns defines and stores column names for the table billing_payments.
type BillingPaymentsColumns struct {
	Id                   string //
	BillingId            string //
	PaymentStatus        string // 支付状态
	PaymentMethod        string // 支付方式: alipay/wechat/credit_card
	PaymentAmount        string // 本次支付金额
	PaymentTransactionId string // 支付平台流水号（第三方返回）
	PaymentChannel       string // 支付渠道: 微信支付/支付宝/银联
	PaidAt               string // 支付成功时间
	FailureReason        string // 支付失败原因
	RefundedAt           string // 退款时间（如有）
	CreatedAt            string // 创建时间（发起支付时）
	UpdatedAt            string // 更新时间（状态变更时）
}

// billingPaymentsColumns holds the columns for the table billing_payments.
var billingPaymentsColumns = BillingPaymentsColumns{
	Id:                   "id",
	BillingId:            "billing_id",
	PaymentStatus:        "payment_status",
	PaymentMethod:        "payment_method",
	PaymentAmount:        "payment_amount",
	PaymentTransactionId: "payment_transaction_id",
	PaymentChannel:       "payment_channel",
	PaidAt:               "paid_at",
	FailureReason:        "failure_reason",
	RefundedAt:           "refunded_at",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
}

// NewBillingPaymentsDao creates and returns a new DAO object for table data access.
func NewBillingPaymentsDao(handlers ...gdb.ModelHandler) *BillingPaymentsDao {
	return &BillingPaymentsDao{
		group:    "default",
		table:    "billing_payments",
		columns:  billingPaymentsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BillingPaymentsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BillingPaymentsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BillingPaymentsDao) Columns() BillingPaymentsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BillingPaymentsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BillingPaymentsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BillingPaymentsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
