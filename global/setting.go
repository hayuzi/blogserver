package global

import (
	"github.com/hayuzi/blogserver/pkg/logger"
	"github.com/hayuzi/blogserver/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	JWTSetting      *setting.JWTSetting
	JaegerSetting   *setting.JaegerSetting
	Logger          *logger.Logger
)
