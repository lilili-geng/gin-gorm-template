package service

import (
	"template-project/app/dto"
	"template-project/app/model"
	"template-project/framework/dal"
)

type UserService struct{}

// 新增用户
func (s *UserService) CreateUser(param dto.User) error {
	tx := dal.Gorm.Begin()

	user := model.User{
		Username: param.Username,
		Password: param.Password,
		Phone:    param.Phone,
		Nickname: param.Nickname,
		Status:   param.Status,
	}

	if err := tx.Model(model.User{}).Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 根据用户名查询用户信息
func (s *UserService) GetUserByUsername(userName string) *model.User {
	var user model.User
	if err := dal.Gorm.Where("phone = ?", userName).First(&user).Error; err != nil {
		return nil
	}
	return &user
}

// 更新用户
func (s *UserService) UpdateUser(param dto.User) error {
	tx := dal.Gorm.Begin()

	if err := tx.Model(&model.User{}).Where("id = ?", param.ID).Updates(param).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
