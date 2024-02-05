package application

 import (
	 "github.com/gogf/gf/v2/frame/g"
 )
 
 type AppInstanceSearchReq struct {
	 g.Meta   `path:"/app/instance/list" tags:"应用实例管理" method:"get" summary:"应用实例列表"`
	 appInstanceName string `p:"appInstanceName"`
	 appInstanceCode string `p:appInstanceCode`
	 Status   string `p:"status"`
 }
 
 type AppInstanceSearchRes struct {
	 g.Meta   `mime:"application/json"`
	 AppInstanceList []*entity.AppInstance `json:"appInstanceList"`
 }
 
 type AppInstanceAddReq struct {
	 g.Meta   `path:"/app/instance/install" tags:"应用实例管理" method:"post" summary:"安装应用"`
	 Name string    `p:"name"  v:"required#应用实例名称不能为空"`
	 Code string    `p:"code"  v:"required#应用实例编码不能为空"`
	 package string `p:"package"  v:"required#应用安装包不能为空"`
	 icon string `p:"icon"`
	 address string `p:"address"`
	 type int `p:"type"  v:"required#应用类型不能为空"`
	 installer string `p:"installer"`
	 status int `p:"package"`
 }
 
 type AppInstanceAddRes struct {
 }
 
 type AppInstanceEditReq struct {
	 g.Meta   `path:"/app/instance/update" tags:"应用实例管理" method:"put" summary:"修改应用实例"`
	 Name string    `p:"name"  v:"required#应用实例名称不能为空"`
	 Code string    `p:"code"  v:"required#应用实例编码不能为空"`
	 package string `p:"package"  v:"required#应用安装包不能为空"`
	 icon string `p:"icon"`
	 address string `p:"address"`
	 type int `p:"type"  v:"required#应用类型不能为空"`
	 installer string `p:"installer"`
	 status int `p:"package"`
 }
 
 type AppInstanceEditRes struct {
 }
 
 type AppInstanceDeleteReq struct {
	 g.Meta `path:"/app/instance/uninstall" tags:"应用实例管理" method:"PUT" summary:"卸载应用"`
	 Id     uint64 `p:"id" v:"required#id不能为空"`
 }
 
 type AppInstanceDeleteRes struct {
 }