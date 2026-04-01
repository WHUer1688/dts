package main

import (
	// module name is travel
	// use other package within the module
	"fmt"
	"travel/internal/api"
	"travel/internal/config"
	"travel/internal/ws"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"time"
)

func main() {
	// 读取配置，按需读取
	config.InitConfig()
	port := fmt.Sprintf(":%d", config.GlobalConfig.Port)

	r := gin.Default()

	// 为了解决本地开发的CORS问题
    r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

	// 注册路由与接口
	// 静态文件托管
	static_file_host_dir := config.GlobalConfig.SFHD
	r.Static("/photos", static_file_host_dir)
	// 接口注册
	r.GET("/hello", api.HttpHello)
	r.POST("/api/v1/photos", api.HttpPostPhotos)
	r.GET("/ws", ws.WsHello)
	
	r.Run(port)
}