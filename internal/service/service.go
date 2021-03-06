package service

import (
	"context"
	"github.com/hayuzi/blogserver/global"
	"github.com/hayuzi/blogserver/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(ctx, global.DBEngine)
	return svc
}
