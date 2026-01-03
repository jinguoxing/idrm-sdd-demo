# 用户认证功能快速开始

> **创建日期**: 2026-01-03

---

## 前置条件

1. Go 1.24+ 已安装
2. MySQL 8.0 已安装并运行
3. Redis 7.0 已安装并运行
4. Go-Zero 工具已安装：`go install github.com/zeromicro/go-zero/tools/goctl@latest`

---

## 开发步骤

### 1. 数据库迁移

```bash
# 执行数据库迁移
mysql -u root -p idrm_demo < migrations/auth/users.sql
mysql -u root -p idrm_demo < migrations/auth/login_history.sql
```

### 2. 配置 API 文件

在 `api/doc/api.api` 中引入认证模块：

```api
import "auth/user_auth.api"
```

### 3. 生成代码

```bash
cd api
goctl api go -api doc/api.api -dir . --style=go_zero --type-group
```

### 4. 实现 Model 层

按照 `model/auth/user/` 和 `model/auth/login_history/` 目录结构实现：
- `interface.go`: 接口定义
- `types.go`: 数据结构
- `vars.go`: 常量和错误定义
- `factory.go`: ORM 工厂函数
- `gorm_dao.go`: GORM 实现
- `sqlx_model.go`: SQLx 实现（可选）

### 5. 实现 Logic 层

在 `api/internal/logic/auth/` 目录下实现各业务逻辑：
- `sendcode_logic.go`: 发送验证码
- `register_logic.go`: 用户注册
- `login_logic.go`: 用户登录
- `logout_logic.go`: 用户登出
- `refresh_logic.go`: 刷新令牌
- `updatepassword_logic.go`: 修改密码
- `sendresetcode_logic.go`: 发送重置密码验证码
- `resetpassword_logic.go`: 重置密码
- `loginhistory_logic.go`: 查询登录历史

### 6. 配置服务上下文

在 `api/internal/svc/servicecontext.go` 中添加 Model 依赖：

```go
type ServiceContext struct {
    Config      config.Config
    UserModel   user.UserModel
    LoginHistoryModel loginhistory.LoginHistoryModel
    Redis       *redis.Client
    // ... 其他依赖
}
```

### 7. 配置 Redis

在 `api/etc/api.yaml` 中添加 Redis 配置：

```yaml
Redis:
  Host: localhost
  Port: 6379
  DB: 0
  Password: ""
```

### 8. 配置短信服务

根据选择的短信服务提供商，配置 AccessKey/SecretKey：

```yaml
SMS:
  Provider: aliyun  # 或 tencent
  AccessKey: your-access-key
  SecretKey: your-secret-key
  SignName: your-sign-name
  TemplateCode: your-template-code
```

### 9. 运行服务

```bash
cd api
go run idrm-sdd-demo-api.go -f etc/api.yaml
```

---

## API 测试示例

### 1. 发送注册验证码

```bash
curl -X POST http://localhost:8888/api/v1/auth/send-code \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "13800138000",
    "type": "register"
  }'
```

### 2. 用户注册

```bash
curl -X POST http://localhost:8888/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "13800138000",
    "code": "123456",
    "password": "Test@1234"
  }'
```

### 3. 用户登录

```bash
curl -X POST http://localhost:8888/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "13800138000",
    "password": "Test@1234"
  }'
```

### 4. 刷新令牌

```bash
curl -X POST http://localhost:8888/api/v1/auth/refresh \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "your-refresh-token"
  }'
```

### 5. 查询登录历史（需要认证）

```bash
curl -X GET "http://localhost:8888/api/v1/auth/login-history?offset=1&limit=20" \
  -H "Authorization: Bearer your-access-token"
```

---

## 测试

### 运行单元测试

```bash
go test ./api/internal/logic/auth/... -v -cover
go test ./model/auth/... -v -cover
```

### 运行集成测试

```bash
go test ./api/internal/integration/... -v
```

---

## 注意事项

1. **密码强度**: 确保密码符合要求（至少8位，包含大小写字母、数字和特殊字符）
2. **验证码有效期**: 验证码5分钟有效，60秒内只能发送一次
3. **令牌有效期**: access_token 2小时，refresh_token 7天
4. **账户锁定**: 连续5次登录失败锁定30分钟
5. **登录历史**: 保留最近90天记录，查询默认显示最近30天

---

## 常见问题

### Q: 验证码发送失败？
A: 检查 Redis 是否正常运行，短信服务配置是否正确。

### Q: 登录返回 401？
A: 检查手机号和密码是否正确，账户是否被锁定。

### Q: 刷新令牌失败？
A: 检查 refresh_token 是否有效，是否已过期。

### Q: 查询登录历史为空？
A: 检查时间范围是否正确，默认查询最近30天。

