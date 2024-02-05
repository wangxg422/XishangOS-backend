package sysAppInstance

 import (
	 "context"
 
 func init() {
	 service.RegisterAppInstance(New())
 }
 
 func New() *sAppInstance {
	 return &sAppInstance{}
 }
 
 type sAppInstance struct {
 }
 
 func (s *sAppInstance) GetList(ctx context.Context, req *system.AppInstanceSearchReq) (list []*entity.SysAppInstance, err error) {
	 list, err = s.GetFromCache(ctx)
	 if err != nil {
		 return
	 }
	 rList := make([]*entity.SysAppInstance, 0, len(list))
	 if req.AppInstanceName != "" || req.Status != "" {
		 for _, v := range list {
			 if req.AppInstanceName != "" && !gstr.ContainsI(v.AppInstanceName, req.AppInstanceName) {
				 continue
			 }
			 if req.Status != "" && v.Status != gconv.Uint(req.Status) {
				 continue
			 }
			 rList = append(rList, v)
		 }
		 list = rList
	 }
	 return
 }
 
 // Add 添加部门
 func (s *sSysAppInstance) Add(ctx context.Context, req *system.AppInstanceAddReq) (err error) {
	 err = g.Try(ctx, func(ctx context.Context) {
		 _, err = dao.SysAppInstance.Ctx(ctx).Insert(do.SysAppInstance{
			 ParentId:  req.ParentID,
			 AppInstanceName:  req.AppInstanceName,
			 OrderNum:  req.OrderNum,
			 Leader:    req.Leader,
			 Phone:     req.Phone,
			 Email:     req.Email,
			 Status:    req.Status,
			 CreatedBy: service.Context().GetUserId(ctx),
		 })
		 liberr.ErrIsNil(ctx, err, "添加部门失败")
		 // 删除缓存
		 commonService.Cache().Remove(ctx, consts.CacheSysAppInstance)
	 })
	 return
 }
 
 // Edit 部门修改
 func (s *sSysAppInstance) Edit(ctx context.Context, req *system.AppInstanceEditReq) (err error) {
	 err = g.Try(ctx, func(ctx context.Context) {
		 _, err = dao.SysAppInstance.Ctx(ctx).WherePri(req.AppInstanceId).Update(do.SysAppInstance{
			 ParentId:  req.ParentID,
			 AppInstanceName:  req.AppInstanceName,
			 OrderNum:  req.OrderNum,
			 Leader:    req.Leader,
			 Phone:     req.Phone,
			 Email:     req.Email,
			 Status:    req.Status,
			 UpdatedBy: service.Context().GetUserId(ctx),
		 })
		 liberr.ErrIsNil(ctx, err, "修改部门失败")
		 // 删除缓存
		 commonService.Cache().Remove(ctx, consts.CacheSysAppInstance)
	 })
	 return
 }
 
 func (s *sSysAppInstance) Delete(ctx context.Context, id uint64) (err error) {
	 err = g.Try(ctx, func(ctx context.Context) {
		 var list []*entity.SysAppInstance
		 err = dao.SysAppInstance.Ctx(ctx).Scan(&list)
		 liberr.ErrIsNil(ctx, err, "不存在部门信息")
		 children := s.FindSonByParentId(list, id)
		 ids := make([]uint64, 0, len(list))
		 for _, v := range children {
			 ids = append(ids, v.AppInstanceId)
		 }
		 ids = append(ids, id)
		 _, err = dao.SysAppInstance.Ctx(ctx).Where(dao.SysAppInstance.Columns().AppInstanceId+" in (?)", ids).Delete()
		 liberr.ErrIsNil(ctx, err, "删除部门失败")
		 // 删除缓存
		 commonService.Cache().Remove(ctx, consts.CacheSysAppInstance)
	 })
	 return
 }
 
 
 // GetByAppInstanceId 通过部门id获取部门信息
 func (s *sSysAppInstance) GetByAppInstanceId(ctx context.Context, AppInstanceId uint64) (AppInstance *entity.SysAppInstance, err error) {
	 var AppInstances []*entity.SysAppInstance
	 AppInstances, err = s.GetFromCache(ctx)
	 if err != nil {
		 return
	 }
	 for _, v := range AppInstances {
		 if v.AppInstanceId == AppInstanceId {
			 AppInstance = v
			 break
		 }
	 }
	 return
 }
 