package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hayuzi/blogserver/global"
	"github.com/hayuzi/blogserver/internal/router"
	"github.com/hayuzi/blogserver/setup"
	"net/http"
)

func main() {
	// init(根据配置启动注册所有全局变量)
	setup.Init()

	global.Logger.Infof("%s: hayuzi/%s", "blogserver", "Infof")

	// 启动路由
	routers := router.NewRouter()
	// 开启服务
	gin.SetMode(global.ServerSetting.RunMode)
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        routers,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
