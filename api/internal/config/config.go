// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis struct {
		Host     string
		Port     int
		DB       int
		Password string
	}
	DB struct {
		Default struct {
			Host            string
			Port            int
			Database        string
			Username        string
			Password        string
			Charset         string
			MaxIdleConns    int
			MaxOpenConns    int
			ConnMaxLifetime int
			ConnMaxIdleTime int
			LogLevel        string
			SlowThreshold   int
			SkipDefaultTxn  bool
			PrepareStmt     bool
			SingularTable   bool
			DisableForeignKey bool
		}
	}
}
