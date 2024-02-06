package sysAppInstance

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/application"
	"github.com/tiger1103/gfast/v3/internal/app/application/model/entity"
)
 
 func init() {
	 service.RegisterAppInstance(New())
 }
 
 func New() *sAppInstance {
	 return &sAppInstance{}
 }
 
 type sAppInstance struct {
 }
 
 func (s *sAppInstance) GetList(ctx context.Context, req *application.AppInstanceSearchReq) (list []*entity.AppInstance, err error) {
	 return
 }
 
 func (s *sAppInstance) Add(ctx context.Context, req *application.AppInstanceAddReq) (err error) {
	 
	 return
 }
 
 func (s *sAppInstance) Edit(ctx context.Context, req *application.AppInstanceEditReq) (err error) {
	 return
 }
 
 func (s *sAppInstance) Delete(ctx context.Context, id uint64) (err error) {
	 return
 }
 
 func (s *sAppInstance) GetByAppInstanceId(ctx context.Context, AppInstanceId uint64) (AppInstance *entity.AppInstance, err error) {
	 return
 }
 