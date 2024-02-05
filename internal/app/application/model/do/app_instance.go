// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AppInstance is the golang structure of table app_instance for DAO operations like Where/Data.
type AppInstance struct {
	g.Meta    `orm:"table:app_instance, do:true"`
	Id        interface{} //
	Name      interface{} // 应用实例名称
	Code      interface{} // 应用实例编码
	Package   interface{} // 应用的安装包
	Icon      interface{} // 应用图标
	Address   interface{} // 访问应用的url或路由
	Type      interface{} // 应用类型（1 内置应用 2 用户自定义应用）
	Installer interface{} // 安装用户
	Status    interface{} // 应用的状态（1 运行中2 已停止 3已卸载）
}
