// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TasksDao is the data access object for the table tasks.
type TasksDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TasksColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TasksColumns defines and stores column names for the table tasks.
type TasksColumns struct {
	Id             string //
	Type           string // 任务类型
	PipelineId     string //
	CreatedAt      string //
	UpdatedAt      string //
	Status         string // 任务状态, pending-待处理, processing-处理中, completed-完成, failed-失败
	StartTime      string //
	FinishTime     string //
	MaterialIdList string // 任务需处理的材料
	ErrorMessage   string // 任务失败日志
	ProjectId      string //
}

// tasksColumns holds the columns for the table tasks.
var tasksColumns = TasksColumns{
	Id:             "id",
	Type:           "type",
	PipelineId:     "pipeline_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	Status:         "status",
	StartTime:      "start_time",
	FinishTime:     "finish_time",
	MaterialIdList: "material_id_list",
	ErrorMessage:   "error_message",
	ProjectId:      "project_id",
}

// NewTasksDao creates and returns a new DAO object for table data access.
func NewTasksDao(handlers ...gdb.ModelHandler) *TasksDao {
	return &TasksDao{
		group:    "default",
		table:    "tasks",
		columns:  tasksColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TasksDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TasksDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TasksDao) Columns() TasksColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TasksDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TasksDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TasksDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
