// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/common/model/entity"
)

type ISysDictType interface {
	List(ctx context.Context, req *system.DictTypeSearchReq) (res *system.DictTypeSearchRes, err error)
	Add(ctx context.Context, req *system.DictTypeAddReq, userId uint64) (err error)
	Edit(ctx context.Context, req *system.DictTypeEditReq, userId uint64) (err error)
	Get(ctx context.Context, req *system.DictTypeGetReq) (dictType *entity.SysDictType, err error)
	ExistsDictType(ctx context.Context, dictType string, dictId ...int64) (err error)
	Delete(ctx context.Context, dictIds []int) (err error)
	GetAllDictType(ctx context.Context) (list []*entity.SysDictType, err error)
}

var localSysDictType ISysDictType

func SysDictType() ISysDictType {
	if localSysDictType == nil {
		panic("implement not found for interface ISysDictType, forgot register?")
	}
	return localSysDictType
}

func RegisterSysDictType(i ISysDictType) {
	localSysDictType = i
}
