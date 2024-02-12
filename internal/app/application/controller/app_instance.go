package controller

import "context"

type applicationController struct {
	BaseController
}

var AppInstance = applicationController{}


// List 系统参数列表
func (c *applicationController) List(ctx context.Context, req *application.AppInstanceSearchReq) (res *system.ConfigSearchRes, err error) {
	res, err = appInstanceService.
	commonService.SysConfig().List(ctx, req)
	return
}

// Add 添加系统参数
func (c *applicationController) Add(ctx context.Context, req *system.ConfigAddReq) (res *system.ConfigAddRes, err error) {
	err = commonService.SysConfig().Add(ctx, req, service.Context().GetUserId(ctx))
	return
}

// Get 获取系统参数
func (c *applicationController) Get(ctx context.Context, req *system.ConfigGetReq) (res *system.ConfigGetRes, err error) {
	res, err = commonService.SysConfig().Get(ctx, req.Id)
	return
}

// Edit 修改系统参数
func (c *applicationController) Edit(ctx context.Context, req *system.ConfigEditReq) (res *system.ConfigEditRes, err error) {
	err = commonService.SysConfig().Edit(ctx, req, service.Context().GetUserId(ctx))
	return
}

// Delete 删除系统参数
func (c *applicationController) Delete(ctx context.Context, req *system.ConfigDeleteReq) (res *system.ConfigDeleteRes, err error) {
	err = commonService.SysConfig().Delete(ctx, req.Ids)
	return
}
