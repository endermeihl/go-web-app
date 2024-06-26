package db

import (
	"go-web-app/models"
	"go-web-app/utils"
)

// CreateUser 在数据库中创建用户
func CreateUser(user *models.User) error {
	return utils.DB.Create(user).Error
}

// GetUserByID 根据ID获取用户
func GetUserByID(id uint) (*models.User, error) {
	var user models.User

	err := utils.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUsers 获取分页用户
func GetUsers(offset int, limit int) ([]models.User, error) {
	var users []models.User
	err := utils.DB.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
