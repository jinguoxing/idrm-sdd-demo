package user

import (
	"context"
	"time"
)

// UserModel 用户数据访问接口
type UserModel interface {
	// Insert 插入新用户
	Insert(ctx context.Context, data *User) (*User, error)

	// FindOne 根据ID查询用户
	FindOne(ctx context.Context, id int64) (*User, error)

	// FindOneByPhone 根据手机号查询用户
	FindOneByPhone(ctx context.Context, phone string) (*User, error)

	// Update 更新用户信息
	Update(ctx context.Context, data *User) error

	// UpdatePassword 更新用户密码
	UpdatePassword(ctx context.Context, id int64, passwordHash string) error

	// UpdateLockStatus 更新账户锁定状态
	UpdateLockStatus(ctx context.Context, id int64, lockedUntil *time.Time, failedAttempts int) error

	// WithTx 使用事务
	WithTx(tx interface{}) UserModel

	// Trans 执行事务操作
	Trans(ctx context.Context, fn func(ctx context.Context, model UserModel) error) error
}
