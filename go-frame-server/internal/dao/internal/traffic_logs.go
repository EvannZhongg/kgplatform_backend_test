// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TrafficLogsDao is the data access object for the table traffic_logs.
type TrafficLogsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TrafficLogsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TrafficLogsColumns defines and stores column names for the table traffic_logs.
type TrafficLogsColumns struct {
	Id          string // 日志ID
	UserId      string // 用户ID
	TeamId      string // 团队ID(如果是团队使用)
	TrafficType string // 流量类型: graph_query/api_call/file_transfer等
	DataSize    string // 数据大小(字节)
	TrafficKb   string //
	Endpoint    string // 请求端点
	IpAddress   string // IP地址
	CreatedAt   string // 创建时间
}

// trafficLogsColumns holds the columns for the table traffic_logs.
var trafficLogsColumns = TrafficLogsColumns{
	Id:          "id",
	UserId:      "user_id",
	TeamId:      "team_id",
	TrafficType: "traffic_type",
	DataSize:    "data_size",
	TrafficKb:   "traffic_kb",
	Endpoint:    "endpoint",
	IpAddress:   "ip_address",
	CreatedAt:   "created_at",
}

// NewTrafficLogsDao creates and returns a new DAO object for table data access.
func NewTrafficLogsDao(handlers ...gdb.ModelHandler) *TrafficLogsDao {
	return &TrafficLogsDao{
		group:    "default",
		table:    "traffic_logs",
		columns:  trafficLogsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TrafficLogsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TrafficLogsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TrafficLogsDao) Columns() TrafficLogsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TrafficLogsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TrafficLogsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TrafficLogsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
