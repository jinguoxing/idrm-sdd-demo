package user

import (
	"gorm.io/gorm"
)

// NewUserModel 创建用户模型实例（GORM实现）
func NewUserModel(db *gorm.DB) UserModel {
	return &gormUserModel{
		db: db,
	}
}

// gormUserModel GORM实现的用户模型
type gormUserModel struct {
	db *gorm.DB
}
