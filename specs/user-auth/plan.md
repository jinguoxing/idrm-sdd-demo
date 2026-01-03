# 用户认证功能 Technical Plan

> **Branch**: `20260103-user-auth`  
> **Spec Path**: `specs/user-auth/`  
> **Created**: 2026-01-03  
> **Status**: Draft

---

## Summary

基于 Go-Zero 框架实现用户认证功能，包括手机号注册（含短信验证码）、登录认证、密码管理和登录历史记录。采用 JWT 双令牌机制（access_token + refresh_token），使用 bcrypt 加密密码，支持账户锁定和登录失败计数等安全机制。

---

## Technical Context

| Item | Value |
|------|-------|
| **Language** | Go 1.24+ |
| **Framework** | Go-Zero v1.9+ |
| **Storage** | MySQL 8.0 |
| **Cache** | Redis 7.0 (用于验证码和令牌管理) |
| **ORM** | GORM / SQLx |
| **Testing** | go test |
| **密码加密** | bcrypt (cost >= 10) |
| **JWT 库** | github.com/golang-jwt/jwt/v5 |
| **短信服务** | 需要集成第三方短信服务（待确定具体提供商） |

---

## Go-Zero 开发流程

按以下顺序完成技术设计和代码生成：

| Step | 任务 | 方式 | 产出 |
|------|------|------|------|
| 1 | 定义 API 文件 | AI 实现 | `api/doc/auth/user_auth.api` |
| 2 | 生成 Handler/Types | goctl 生成 | `api/internal/handler/auth/`, `types/` |
| 3 | 定义 DDL 文件 | AI 手写 | `migrations/auth/users.sql`, `migrations/auth/login_history.sql` |
| 4 | 实现 Model 接口 | AI 手写 | `model/auth/user/`, `model/auth/login_history/` |
| 5 | 实现 Logic 层 | AI 实现 | `api/internal/logic/auth/` |

> ⚠️ **重要**：goctl 必须在 `api/doc/api.api` 入口文件上执行，不能针对单个功能文件！

**goctl 命令**:
```bash
# 步骤1：在 api/doc/api.api 中 import 新模块
# 步骤2：执行 goctl 生成代码（针对整个项目）
goctl api go -api api/doc/api.api -dir api/ --style=go_zero --type-group
```

---

## File Structure

### 文件产出清单

| 序号 | 文件 | 生成方式 | 位置 |
|------|------|----------|------|
| 1 | API 文件 | AI 实现 | `api/doc/auth/user_auth.api` |
| 2 | DDL 文件 | AI 实现 | `migrations/auth/users.sql`, `migrations/auth/login_history.sql` |
| 3 | Handler | goctl 生成 | `api/internal/handler/auth/` |
| 4 | Types | goctl 生成 | `api/internal/types/` |
| 5 | Logic | AI 实现 | `api/internal/logic/auth/` |
| 6 | Model | AI 实现 | `model/auth/user/`, `model/auth/login_history/` |

### 代码结构

```
api/internal/
├── handler/auth/
│   ├── sendcode_handler.go          # goctl 生成
│   ├── register_handler.go
│   ├── login_handler.go
│   ├── logout_handler.go
│   ├── refresh_handler.go
│   ├── updatepassword_handler.go
│   ├── sendresetcode_handler.go
│   ├── resetpassword_handler.go
│   ├── loginhistory_handler.go
│   └── routes.go
├── logic/auth/
│   ├── sendcode_logic.go            # AI 实现
│   ├── register_logic.go
│   ├── login_logic.go
│   ├── logout_logic.go
│   ├── refresh_logic.go
│   ├── updatepassword_logic.go
│   ├── sendresetcode_logic.go
│   ├── resetpassword_logic.go
│   └── loginhistory_logic.go
├── types/
│   └── types.go                     # goctl 生成
└── svc/
    └── servicecontext.go            # 手动维护

model/auth/user/
├── interface.go                     # 接口定义
├── types.go                         # 数据结构
├── vars.go                          # 常量/错误
├── factory.go                       # ORM 工厂
├── gorm_dao.go                      # GORM 实现
└── sqlx_model.go                    # SQLx 实现

model/auth/login_history/
├── interface.go
├── types.go
├── vars.go
├── factory.go
├── gorm_dao.go
└── sqlx_model.go
```

---

## Architecture Overview

遵循 IDRM 分层架构：

```
HTTP Request → Handler → Logic → Model → Database
```

| 层级 | 职责 | 最大行数 |
|------|------|----------|
| Handler | 解析参数、格式化响应 | 30 |
| Logic | 业务逻辑实现 | 50 |
| Model | 数据访问 | 50 |

---

## Interface Definitions

### User Model Interface

```go
type UserModel interface {
    Insert(ctx context.Context, data *User) (*User, error)
    FindOne(ctx context.Context, id int64) (*User, error)
    FindOneByPhone(ctx context.Context, phone string) (*User, error)
    Update(ctx context.Context, data *User) error
    UpdatePassword(ctx context.Context, id int64, passwordHash string) error
    UpdateLockStatus(ctx context.Context, id int64, lockedUntil *time.Time, failedAttempts int) error
    WithTx(tx interface{}) UserModel
    Trans(ctx context.Context, fn func(ctx context.Context, model UserModel) error) error
}
```

### LoginHistory Model Interface

```go
type LoginHistoryModel interface {
    Insert(ctx context.Context, data *LoginHistory) (*LoginHistory, error)
    FindByUserId(ctx context.Context, userId int64, startTime, endTime time.Time, offset, limit int) ([]*LoginHistory, int64, error)
    WithTx(tx interface{}) LoginHistoryModel
}
```

---

## Data Model

详细数据模型定义见 `data-model.md`

---

## API Contract

详细 API 合约定义见 `contracts/user_auth.api`

---

## Testing Strategy

| 类型 | 方法 | 覆盖率 |
|------|------|--------|
| 单元测试 | 表驱动测试，Mock Model | > 80% |
| 集成测试 | 测试数据库 | 核心流程（注册、登录、密码重置） |

---

## Revision History

| Version | Date | Author | Changes |
|---------|------|--------|---------|
| 1.0 | 2026-01-03 | AI | 初始版本 |

