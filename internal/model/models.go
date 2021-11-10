package model

import (
	"fmt"
	"github.com/hayuzi/blogserver/global"
	"github.com/hayuzi/blogserver/pkg/gormplugins"
	"github.com/hayuzi/blogserver/pkg/setting"
	"github.com/hayuzi/blogserver/pkg/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type Model struct {
	Id        int            `json:"id" gorm:"primary_key"`
	CreatedAt util.JSONTime  `json:"createdAt" gorm:"autoCreateTime:milli"` // 创建时间 datetime
	UpdatedAt util.JSONTime  `json:"updatedAt" gorm:"autoUpdateTime:milli"` // 更新时间 datetime
	DeletedAt gorm.DeletedAt `json:"-"`                                     // 软删除字段(可以为NULL)  datetime
}

func NewDBEngine(dbSetting *setting.DatabaseSetting) (*gorm.DB, error) {
	var err error
	// %s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Asia%%2FShanghai  使用上海东八区，但是在scratch的docker镜像中不可用
	// %s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local， 服务器本地时区未设置
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		dbSetting.UserName,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.DBName,
		dbSetting.Charset,
		dbSetting.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatalf("model.NewDBEngine err: %v", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	if global.ServerSetting.RunMode == "debug" {
		db.Logger.LogMode(logger.Info)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	gormplugins.AddGormCallbacks(db)
	return db, nil
}

func (model *Model) BeforeCreate(db *gorm.DB) error {
	return nil
}

func (model *Model) BeforeUpdate(db *gorm.DB) error {
	return nil
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
