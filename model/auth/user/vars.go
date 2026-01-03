package user

// 用户状态枚举
const (
	UserStatusNormal   = 0 // 正常
	UserStatusLocked   = 1 // 锁定
	UserStatusDisabled = 2 // 禁用
)

// 错误码定义（20000-29999 为业务错误）
const (
	ErrUserNotFound      = 20001 // 用户不存在
	ErrUserAlreadyExists = 20002 // 用户已存在
	ErrUserLocked        = 20003 // 用户已锁定
	ErrUserDisabled      = 20004 // 用户已禁用
)

// 错误消息
const (
	ErrMsgUserNotFound      = "用户不存在"
	ErrMsgUserAlreadyExists = "用户已存在"
	ErrMsgUserLocked        = "账户已锁定"
	ErrMsgUserDisabled      = "账户已禁用"
)
