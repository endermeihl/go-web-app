package main

import (
	"go-web-app/config"
	"go-web-app/logger"
	"go-web-app/middlewares"
	"go-web-app/routes"
	"go-web-app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 初始化配置
	config.LoadConfig()

	// 初始化日志
	logger.InitLogger()
	defer logger.SyncLogger()

	// 连接数据库
	utils.ConnectDatabase()
	utils.ConnectRedis()

	// 初始化 Gin 引擎
	r := gin.New()

	// 使用中间件
	r.Use(middlewares.EncryptionMiddleware())
	r.Use(middlewares.AuthMiddleware())

	// 添加恢复中间件
	r.Use(gin.Recovery())

	// 注册路由
	routes.RegisterRoutes(r)

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		logger.Logger.Fatal("Failed to start server", zap.Error(err))
	}
}
