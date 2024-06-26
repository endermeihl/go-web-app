package controllers

import (
	"go-web-app/logger"
	"go-web-app/models"
	"go-web-app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateUser 创建用户
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Logger.Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateUserService(&user); err != nil {
		logger.Logger.Error("Failed to insert user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Logger.Info("User created", zap.Uint("userID", uint(user.ID)))
	c.JSON(http.StatusOK, user)
}

// GetUser 获取用户
func GetUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Logger.Error("Invalid user ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := services.GetUserByIDService(uint(userID))
	if err != nil {
		logger.Logger.Error("Failed to query user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Logger.Info("User fetched", zap.Uint("userID", uint(userID)))
	c.JSON(http.StatusOK, user)
}

// GetUsers 分页查询用户
func GetUsers(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	users, err := services.GetUsersService(page, pageSize)
	if err != nil {
		logger.Logger.Error("Failed to fetch users", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, users)
}
