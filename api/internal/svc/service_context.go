// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"fmt"
	"github.com/jinguoxing/idrm-sdd-demo/api/internal/config"
	"github.com/jinguoxing/idrm-sdd-demo/model/auth/user"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	UserModel user.UserModel
	Redis     *redis.Client
	DB        *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化 GORM 数据库连接
	dsn := buildDSN(c.DB.Default)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// 初始化 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     buildRedisAddr(c.Redis),
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})

	// 初始化 UserModel
	userModel := user.NewUserModel(db)

	return &ServiceContext{
		Config:    c,
		UserModel: userModel,
		Redis:     rdb,
		DB:        db,
	}
}

// buildDSN 构建 MySQL DSN
func buildDSN(cfg config.Config.DB.Default) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Charset)
}

// buildRedisAddr 构建 Redis 地址
func buildRedisAddr(cfg config.Config.Redis) string {
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}
