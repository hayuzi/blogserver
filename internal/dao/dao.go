package dao

import (
	"context"
	"gorm.io/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(ctx context.Context, engine *gorm.DB) *Dao {
	return &Dao{
		engine: engine.WithContext(ctx),
	}
}
