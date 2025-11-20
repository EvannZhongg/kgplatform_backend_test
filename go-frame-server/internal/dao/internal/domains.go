// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DomainsDao is the data access object for the table domains.
type DomainsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  DomainsColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// DomainsColumns defines and stores column names for the table domains.
type DomainsColumns struct {
	DomainCatalog          string //
	Id                     string //
	DisplayName            string // 展示名称
	DomainSchema           string //
	DomainName             string //
	Status                 string // 状态：1启用，0禁用
	DataType               string //
	Description            string //
	CharacterMaximumLength string //
	CreatedAt              string //
	UpdatedAt              string //
	CharacterOctetLength   string //
	CharacterSetCatalog    string //
	CharacterSetSchema     string //
	CharacterSetName       string //
	CollationCatalog       string //
	CollationSchema        string //
	CollationName          string //
	NumericPrecision       string //
	NumericPrecisionRadix  string //
	NumericScale           string //
	DatetimePrecision      string //
	IntervalType           string //
	IntervalPrecision      string //
	DomainDefault          string //
	UdtCatalog             string //
	UdtSchema              string //
	UdtName                string //
	ScopeCatalog           string //
	ScopeSchema            string //
	ScopeName              string //
	MaximumCardinality     string //
	DtdIdentifier          string //
}

// domainsColumns holds the columns for the table domains.
var domainsColumns = DomainsColumns{
	DomainCatalog:          "domain_catalog",
	Id:                     "id",
	DisplayName:            "display_name",
	DomainSchema:           "domain_schema",
	DomainName:             "domain_name",
	Status:                 "status",
	DataType:               "data_type",
	Description:            "description",
	CharacterMaximumLength: "character_maximum_length",
	CreatedAt:              "created_at",
	UpdatedAt:              "updated_at",
	CharacterOctetLength:   "character_octet_length",
	CharacterSetCatalog:    "character_set_catalog",
	CharacterSetSchema:     "character_set_schema",
	CharacterSetName:       "character_set_name",
	CollationCatalog:       "collation_catalog",
	CollationSchema:        "collation_schema",
	CollationName:          "collation_name",
	NumericPrecision:       "numeric_precision",
	NumericPrecisionRadix:  "numeric_precision_radix",
	NumericScale:           "numeric_scale",
	DatetimePrecision:      "datetime_precision",
	IntervalType:           "interval_type",
	IntervalPrecision:      "interval_precision",
	DomainDefault:          "domain_default",
	UdtCatalog:             "udt_catalog",
	UdtSchema:              "udt_schema",
	UdtName:                "udt_name",
	ScopeCatalog:           "scope_catalog",
	ScopeSchema:            "scope_schema",
	ScopeName:              "scope_name",
	MaximumCardinality:     "maximum_cardinality",
	DtdIdentifier:          "dtd_identifier",
}

// NewDomainsDao creates and returns a new DAO object for table data access.
func NewDomainsDao(handlers ...gdb.ModelHandler) *DomainsDao {
	return &DomainsDao{
		group:    "default",
		table:    "domains",
		columns:  domainsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DomainsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DomainsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DomainsDao) Columns() DomainsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DomainsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DomainsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DomainsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
