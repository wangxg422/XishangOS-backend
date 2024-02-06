// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/application"
	"github.com/tiger1103/gfast/v3/internal/app/application/model/entity"
)

type (
	IAppInstance interface {
		GetList(ctx context.Context, req *application.AppInstanceSearchReq) (list []*entity.AppInstance, err error)
		Add(ctx context.Context, req *application.AppInstanceAddReq) (err error)
		Edit(ctx context.Context, req *application.AppInstanceEditReq) (err error)
		Delete(ctx context.Context, id uint64) (err error)
		GetByAppInstanceId(ctx context.Context, AppInstanceId uint64) (AppInstance *entity.AppInstance, err error)
	}
)

var (
	localAppInstance IAppInstance
)

func AppInstance() IAppInstance {
	if localAppInstance == nil {
		panic("implement not found for interface IAppInstance, forgot register?")
	}
	return localAppInstance
}

func RegisterAppInstance(i IAppInstance) {
	localAppInstance = i
}
