// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe-shop-v2/internal/model"
)

type (
	IColletion interface {
		AddCollection(ctx context.Context, in model.AddCollectionInput) (res *model.AddCollectionOutput, err error)
		DeleteCollection(ctx context.Context, in model.DeleteCollectionInput) (res *model.DeleteCollectionOutput, err error)
	}
)

var (
	localColletion IColletion
)

func Colletion() IColletion {
	if localColletion == nil {
		panic("implement not found for interface IColletion, forgot register?")
	}
	return localColletion
}

func RegisterColletion(i IColletion) {
	localColletion = i
}
