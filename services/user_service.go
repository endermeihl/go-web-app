package services

import (
	"go-web-app/db"
	"go-web-app/models"
)

// CreateUserService 创建用户
func CreateUserService(user *models.User) error {
	return db.CreateUser(user)
}

// GetUserByIDService 根据ID获取用户
func GetUserByIDService(id uint) (*models.User, error) {
	return db.GetUserByID(id)
}

// GetUsersService 获取分页用户
func GetUsersService(page int, pageSize int) ([]models.User, error) {
	offset := (page - 1) * pageSize
	return db.GetUsers(offset, pageSize)
}
