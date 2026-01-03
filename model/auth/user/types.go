package user

import (
	"time"

	"gorm.io/gorm"
)

// User 用户实体
type User struct {
	Id             int64          `gorm:"primaryKey" json:"id"`
	Phone          string         `gorm:"size:11;uniqueIndex;not null" json:"phone"`
	PasswordHash   string         `gorm:"size:255;not null" json:"-"`
	Nickname       *string        `gorm:"size:50" json:"nickname"`
	Status         int            `gorm:"type:tinyint;not null;default:0" json:"status"`
	FailedAttempts int            `gorm:"type:int;not null;default:0" json:"failed_attempts"`
	LockedUntil    *time.Time     `gorm:"type:datetime" json:"locked_until"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
