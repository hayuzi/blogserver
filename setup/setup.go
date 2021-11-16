package setup

import (
	"github.com/hayuzi/blogserver/global"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/logger"
	"github.com/hayuzi/blogserver/pkg/setting"
	"github.com/hayuzi/blogserver/pkg/tracer"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"strings"
	"time"
)

func Init(config string) {
	err := setupSetting(config)
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

func setupSetting(configPath string) error {
	st, err := setting.NewSetting(strings.Split(configPath, ",")...)
	if err != nil {
		return err
	}
	err = st.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = st.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = st.ReadSection("DataBase", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = st.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Minute

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	ljLogger := &lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600, //最大占用空间为600MB
		MaxAge:    10,  //最大生存周期为10天
		LocalTime: true,
	}
	// WithCaller(3) 回溯到入口文件
	//global.Logger = logger.NewLogger(ljLogger, "", log.LstdFlags).WithCaller(3)
	global.Logger = logger.NewLogger(ljLogger, "", log.LstdFlags)
	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blogserver", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
