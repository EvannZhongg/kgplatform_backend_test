// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserSubscriptionsDao is the data access object for the table user_subscriptions.
type UserSubscriptionsDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  UserSubscriptionsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// UserSubscriptionsColumns defines and stores column names for the table user_subscriptions.
type UserSubscriptionsColumns struct {
	Id                    string //
	UserId                string //
	TeamId                string //
	UserPlan              string //
	SubscriptionStatus    string //
	WordsUsed             string //
	StorageUsed           string //
	CuUsed                string //
	TrafficUsed           string //
	QuotaResetDate        string //
	WordsWarning80Sent    string //
	WordsWarning100Sent   string //
	StorageWarning80Sent  string //
	StorageWarning100Sent string //
	CuWarning80Sent       string //
	CuWarning100Sent      string //
	TrafficWarning80Sent  string //
	TrafficWarning100Sent string //
	OverageWordsFee       string //
	OverageStorageFee     string //
	OverageTrafficFee     string //
	OverageCuFee          string //
	TotalOverageFee       string //
	SelectedAiModel       string //
	CreatedAt             string //
	UpdatedAt             string //
}

// userSubscriptionsColumns holds the columns for the table user_subscriptions.
var userSubscriptionsColumns = UserSubscriptionsColumns{
	Id:                    "id",
	UserId:                "user_id",
	TeamId:                "team_id",
	UserPlan:              "user_plan",
	SubscriptionStatus:    "subscription_status",
	WordsUsed:             "words_used",
	StorageUsed:           "storage_used",
	CuUsed:                "cu_used",
	TrafficUsed:           "traffic_used",
	QuotaResetDate:        "quota_reset_date",
	WordsWarning80Sent:    "words_warning_80_sent",
	WordsWarning100Sent:   "words_warning_100_sent",
	StorageWarning80Sent:  "storage_warning_80_sent",
	StorageWarning100Sent: "storage_warning_100_sent",
	CuWarning80Sent:       "cu_warning_80_sent",
	CuWarning100Sent:      "cu_warning_100_sent",
	TrafficWarning80Sent:  "traffic_warning_80_sent",
	TrafficWarning100Sent: "traffic_warning_100_sent",
	OverageWordsFee:       "overage_words_fee",
	OverageStorageFee:     "overage_storage_fee",
	OverageTrafficFee:     "overage_traffic_fee",
	OverageCuFee:          "overage_cu_fee",
	TotalOverageFee:       "total_overage_fee",
	SelectedAiModel:       "selected_ai_model",
	CreatedAt:             "created_at",
	UpdatedAt:             "updated_at",
}

// NewUserSubscriptionsDao creates and returns a new DAO object for table data access.
func NewUserSubscriptionsDao(handlers ...gdb.ModelHandler) *UserSubscriptionsDao {
	return &UserSubscriptionsDao{
		group:    "default",
		table:    "user_subscriptions",
		columns:  userSubscriptionsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserSubscriptionsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserSubscriptionsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserSubscriptionsDao) Columns() UserSubscriptionsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserSubscriptionsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserSubscriptionsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserSubscriptionsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
