// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BillingOveragesDao is the data access object for the table billing_overages.
type BillingOveragesDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  BillingOveragesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// BillingOveragesColumns defines and stores column names for the table billing_overages.
type BillingOveragesColumns struct {
	Id                   string //
	BillingId            string //
	WordsOverageAmount   string // 字数超额量（单位：千字）
	WordsOverageFee      string // 字数超额费用
	StorageOverageAmount string // 存储超额量（单位：GB）
	StorageOverageFee    string // 存储超额费用
	TrafficOverageAmount string // 流量超额量（单位：GB）
	TrafficOverageFee    string // 流量超额费用
	CuOverageAmount      string // 计算单元超额量（单位：CU）
	CuOverageFee         string // 计算单元超额费用
	TotalOverageFee      string // 超额费用汇总
	CreatedAt            string // 创建时间
	UpdatedAt            string // 更新时间
}

// billingOveragesColumns holds the columns for the table billing_overages.
var billingOveragesColumns = BillingOveragesColumns{
	Id:                   "id",
	BillingId:            "billing_id",
	WordsOverageAmount:   "words_overage_amount",
	WordsOverageFee:      "words_overage_fee",
	StorageOverageAmount: "storage_overage_amount",
	StorageOverageFee:    "storage_overage_fee",
	TrafficOverageAmount: "traffic_overage_amount",
	TrafficOverageFee:    "traffic_overage_fee",
	CuOverageAmount:      "cu_overage_amount",
	CuOverageFee:         "cu_overage_fee",
	TotalOverageFee:      "total_overage_fee",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
}

// NewBillingOveragesDao creates and returns a new DAO object for table data access.
func NewBillingOveragesDao(handlers ...gdb.ModelHandler) *BillingOveragesDao {
	return &BillingOveragesDao{
		group:    "default",
		table:    "billing_overages",
		columns:  billingOveragesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BillingOveragesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BillingOveragesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BillingOveragesDao) Columns() BillingOveragesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BillingOveragesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BillingOveragesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BillingOveragesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
