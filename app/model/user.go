package model

import (
	"time"
)

// AppUser 应用用户结构体
type User struct {
	ID            int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Username      string    `json:"username" gorm:"size:50;not null" validate:"required,min=3,max=50"`
	Password      string    `json:"password,omitempty" gorm:"size:255;not null" validate:"required,min=6,max=255"`
	Phone         string    `json:"phone" gorm:"size:20;not null;default:''" validate:"omitempty,numeric,len=11"`
	Email         string    `json:"email" gorm:"size:100;not null;default:''" validate:"omitempty,email"`
	Nickname      string    `json:"nickname" gorm:"size:50;not null;default:''"`
	Avatar        string    `json:"avatar" gorm:"size:255;not null;default:''"`
	RealName      string    `json:"realName" gorm:"size:50;not null;default:''" validate:"omitempty,max=50"`
	Gender        int       `json:"gender" gorm:"type:tinyint;not null;default:0" validate:"oneof=0 1 2"`
	Status        int       `json:"status" gorm:"type:tinyint;not null;default:1" validate:"oneof=0 1"`
	LastLoginTime time.Time `json:"last_login_time" gorm:"column:last_login_time"`
	LastLoginIP   string    `json:"last_login_ip" gorm:"column:last_login_ip"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:create_time;autoCreateTime"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:update_time;autoUpdateTime"`
}

// TableName 设置表名
func (User) TableName() string {
	return "user"
}
