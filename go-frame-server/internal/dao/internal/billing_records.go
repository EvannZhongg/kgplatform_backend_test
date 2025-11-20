// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BillingRecordsDao is the data access object for the table billing_records.
type BillingRecordsDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  BillingRecordsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// BillingRecordsColumns defines and stores column names for the table billing_records.
type BillingRecordsColumns struct {
	Id                  string //
	UserId              string //
	TeamId              string //
	BillingPeriod       string // 账期 YYYY-MM
	BillingDate         string // 结算日期
	BillingType         string // 账单类型
	BaseSubscriptionFee string // 基础订阅费
	OverageFee          string // 超额费用总额
	Subtotal            string // 小计（基础费+超额费）
	DiscountAmount      string // 折扣金额
	TotalAmount         string // 应付总金额
	Status              string //
	Remark              string // 备注说明
	CreatedAt           string // 创建时间
	UpdatedAt           string // 更新时间
	AlipayPaymentUrl    string //
	WechatPaymentUrl    string //
}

// billingRecordsColumns holds the columns for the table billing_records.
var billingRecordsColumns = BillingRecordsColumns{
	Id:                  "id",
	UserId:              "user_id",
	TeamId:              "team_id",
	BillingPeriod:       "billing_period",
	BillingDate:         "billing_date",
	BillingType:         "billing_type",
	BaseSubscriptionFee: "base_subscription_fee",
	OverageFee:          "overage_fee",
	Subtotal:            "subtotal",
	DiscountAmount:      "discount_amount",
	TotalAmount:         "total_amount",
	Status:              "status",
	Remark:              "remark",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	AlipayPaymentUrl:    "alipay_payment_url",
	WechatPaymentUrl:    "wechat_payment_url",
}

// NewBillingRecordsDao creates and returns a new DAO object for table data access.
func NewBillingRecordsDao(handlers ...gdb.ModelHandler) *BillingRecordsDao {
	return &BillingRecordsDao{
		group:    "default",
		table:    "billing_records",
		columns:  billingRecordsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BillingRecordsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BillingRecordsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BillingRecordsDao) Columns() BillingRecordsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BillingRecordsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BillingRecordsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BillingRecordsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
