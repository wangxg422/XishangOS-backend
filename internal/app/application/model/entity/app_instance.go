// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AppInstance is the golang structure for table app_instance.
type AppInstance struct {
	Id        int64  `json:"id"        description:""`
	Name      string `json:"name"      description:"应用实例名称"`
	Code      string `json:"code"      description:"应用实例编码"`
	Package   int64  `json:"package"   description:"应用的安装包"`
	Icon      string `json:"icon"      description:"应用图标"`
	Address   string `json:"address"   description:"访问应用的url或路由"`
	Type      uint   `json:"type"      description:"应用类型（1 内置应用 2 用户自定义应用）"`
	Installer int64  `json:"installer" description:"安装用户"`
	Status    uint   `json:"status"    description:"应用的状态（1 运行中2 已停止 3已卸载）"`
}
