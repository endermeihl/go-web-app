package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Helloweb(c *gin.Context) {

	c.JSON(http.StatusOK, "Hello World!")
}
