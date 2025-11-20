// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TeamsDao is the data access object for the table teams.
type TeamsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TeamsColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TeamsColumns defines and stores column names for the table teams.
type TeamsColumns struct {
	Id                string //
	TeamName          string //
	OwnerId           string //
	TeamCode          string //
	InviteCode        string //
	TotalWordsQuota   string //
	TotalStorageQuota string //
	TotalCuQuota      string //
	TotalTrafficQuota string //
	WordsUsed         string //
	StorageUsed       string //
	CuUsed            string //
	TrafficUsed       string //
	MemberCount       string //
	Status            string //
	CreatedAt         string //
	UpdatedAt         string //
}

// teamsColumns holds the columns for the table teams.
var teamsColumns = TeamsColumns{
	Id:                "id",
	TeamName:          "team_name",
	OwnerId:           "owner_id",
	TeamCode:          "team_code",
	InviteCode:        "invite_code",
	TotalWordsQuota:   "total_words_quota",
	TotalStorageQuota: "total_storage_quota",
	TotalCuQuota:      "total_cu_quota",
	TotalTrafficQuota: "total_traffic_quota",
	WordsUsed:         "words_used",
	StorageUsed:       "storage_used",
	CuUsed:            "cu_used",
	TrafficUsed:       "traffic_used",
	MemberCount:       "member_count",
	Status:            "status",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// NewTeamsDao creates and returns a new DAO object for table data access.
func NewTeamsDao(handlers ...gdb.ModelHandler) *TeamsDao {
	return &TeamsDao{
		group:    "default",
		table:    "teams",
		columns:  teamsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TeamsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TeamsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TeamsDao) Columns() TeamsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TeamsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TeamsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TeamsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
