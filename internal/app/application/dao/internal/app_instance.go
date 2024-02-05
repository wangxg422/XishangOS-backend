// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppInstanceDao is the data access object for table app_instance.
type AppInstanceDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns AppInstanceColumns // columns contains all the column names of Table for convenient usage.
}

// AppInstanceColumns defines and stores column names for table app_instance.
type AppInstanceColumns struct {
	Id        string //
	Name      string // 应用实例名称
	Code      string // 应用实例编码
	Package   string // 应用的安装包
	Icon      string // 应用图标
	Address   string // 访问应用的url或路由
	Type      string // 应用类型（1 内置应用 2 用户自定义应用）
	Installer string // 安装用户
	Status    string // 应用的状态（1 运行中2 已停止 3已卸载）
}

// appInstanceColumns holds the columns for table app_instance.
var appInstanceColumns = AppInstanceColumns{
	Id:        "id",
	Name:      "name",
	Code:      "code",
	Package:   "package",
	Icon:      "icon",
	Address:   "address",
	Type:      "type",
	Installer: "installer",
	Status:    "status",
}

// NewAppInstanceDao creates and returns a new DAO object for table data access.
func NewAppInstanceDao() *AppInstanceDao {
	return &AppInstanceDao{
		group:   "default",
		table:   "app_instance",
		columns: appInstanceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AppInstanceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AppInstanceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AppInstanceDao) Columns() AppInstanceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AppInstanceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AppInstanceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AppInstanceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
