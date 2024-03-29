// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAuthRule is the golang structure for table sys_auth_rule.
type SysAuthRule struct {
	Id         uint        `json:"id"         description:""`
	Pid        uint        `json:"pid"        description:"父ID"`
	Name       string      `json:"name"       description:"规则名称"`
	Title      string      `json:"title"      description:"规则名称"`
	Icon       string      `json:"icon"       description:"图标"`
	Condition  string      `json:"condition"  description:"条件"`
	Remark     string      `json:"remark"     description:"备注"`
	MenuType   uint        `json:"menuType"   description:"类型 0目录 1菜单 2按钮"`
	Weigh      int         `json:"weigh"      description:"权重"`
	IsHide     uint        `json:"isHide"     description:"显示状态"`
	Path       string      `json:"path"       description:"路由地址"`
	Component  string      `json:"component"  description:"组件路径"`
	IsLink     uint        `json:"isLink"     description:"是否外链 1是 0否"`
	ModuleType string      `json:"moduleType" description:"所属模块"`
	ModelId    uint        `json:"modelId"    description:"模型ID"`
	IsIframe   uint        `json:"isIframe"   description:"是否内嵌iframe"`
	IsCached   uint        `json:"isCached"   description:"是否缓存"`
	Redirect   string      `json:"redirect"   description:"路由重定向地址"`
	IsAffix    uint        `json:"isAffix"    description:"是否固定"`
	LinkUrl    string      `json:"linkUrl"    description:"链接地址"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建日期"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:"修改日期"`
}
