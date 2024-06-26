package controllers

import (
	"go-web-app/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Helloweb(c *gin.Context) {
	logger.Logger.Info("Hello World", zap.String("id", c.Param("id")))

	c.JSON(http.StatusOK, "Hello World!")
}
