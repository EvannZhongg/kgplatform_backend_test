// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TeamMembersDao is the data access object for the table team_members.
type TeamMembersDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TeamMembersColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TeamMembersColumns defines and stores column names for the table team_members.
type TeamMembersColumns struct {
	Id                  string //
	TeamId              string //
	UserId              string //
	Role                string //
	AllocatedWordsQuota string //
	PersonalWordsUsed   string //
	PersonalStorageUsed string //
	PersonalCuUsed      string //
	PersonalTrafficUsed string //
	Status              string //
	InviteCode          string //
	InvitedBy           string //
	JoinedAt            string //
	RemovedAt           string //
}

// teamMembersColumns holds the columns for the table team_members.
var teamMembersColumns = TeamMembersColumns{
	Id:                  "id",
	TeamId:              "team_id",
	UserId:              "user_id",
	Role:                "role",
	AllocatedWordsQuota: "allocated_words_quota",
	PersonalWordsUsed:   "personal_words_used",
	PersonalStorageUsed: "personal_storage_used",
	PersonalCuUsed:      "personal_cu_used",
	PersonalTrafficUsed: "personal_traffic_used",
	Status:              "status",
	InviteCode:          "invite_code",
	InvitedBy:           "invited_by",
	JoinedAt:            "joined_at",
	RemovedAt:           "removed_at",
}

// NewTeamMembersDao creates and returns a new DAO object for table data access.
func NewTeamMembersDao(handlers ...gdb.ModelHandler) *TeamMembersDao {
	return &TeamMembersDao{
		group:    "default",
		table:    "team_members",
		columns:  teamMembersColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TeamMembersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TeamMembersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TeamMembersDao) Columns() TeamMembersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TeamMembersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TeamMembersDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TeamMembersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
