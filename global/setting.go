package global

import (
	"github.com/hayuzi/blogserver/pkg/logger"
	"github.com/hayuzi/blogserver/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	Logger          *logger.Logger
)
