package cmd

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hayuzi/blogserver/global"
	"github.com/hayuzi/blogserver/internal/router"
	"github.com/hayuzi/blogserver/setup"
	"github.com/spf13/cobra"
	"net/http"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "启动HTTP服务",
	Long:  "使用Gin框架启动HTTP服务并常驻运行",
	Run:   startGinServer,
}

func startGinServer(cmd *cobra.Command, args []string) {
	// init ( 根据配置启动注册所有全局变量, 供全部业务使用 )
	setup.Init(ConfigPath)
	// 启动日志
	global.Logger.Infof(context.Background(), "%s: hayuzi/%s", "blogserver", "start")
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

func init() {
	RootCmd.AddCommand(serverCmd)
}
