package user

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

// Insert 插入新用户
func (m *gormUserModel) Insert(ctx context.Context, data *User) (*User, error) {
	if err := m.db.WithContext(ctx).Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// FindOne 根据ID查询用户
func (m *gormUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	var user User
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(ErrMsgUserNotFound)
		}
		return nil, err
	}
	return &user, nil
}

// FindOneByPhone 根据手机号查询用户
func (m *gormUserModel) FindOneByPhone(ctx context.Context, phone string) (*User, error) {
	var user User
	err := m.db.WithContext(ctx).Where("phone = ?", phone).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(ErrMsgUserNotFound)
		}
		return nil, err
	}
	return &user, nil
}

// Update 更新用户信息
func (m *gormUserModel) Update(ctx context.Context, data *User) error {
	return m.db.WithContext(ctx).Model(data).Updates(data).Error
}

// UpdatePassword 更新用户密码
func (m *gormUserModel) UpdatePassword(ctx context.Context, id int64, passwordHash string) error {
	return m.db.WithContext(ctx).Model(&User{}).
		Where("id = ?", id).
		Update("password_hash", passwordHash).Error
}

// UpdateLockStatus 更新账户锁定状态
func (m *gormUserModel) UpdateLockStatus(ctx context.Context, id int64, lockedUntil *time.Time, failedAttempts int) error {
	updates := map[string]interface{}{
		"failed_attempts": failedAttempts,
	}
	if lockedUntil != nil {
		updates["locked_until"] = lockedUntil
	} else {
		updates["locked_until"] = nil
	}
	return m.db.WithContext(ctx).Model(&User{}).
		Where("id = ?", id).
		Updates(updates).Error
}

// WithTx 使用事务
func (m *gormUserModel) WithTx(tx interface{}) UserModel {
	if gormTx, ok := tx.(*gorm.DB); ok {
		return &gormUserModel{db: gormTx}
	}
	return m
}

// Trans 执行事务操作
func (m *gormUserModel) Trans(ctx context.Context, fn func(ctx context.Context, model UserModel) error) error {
	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txModel := &gormUserModel{db: tx}
		return fn(ctx, txModel)
	})
}
