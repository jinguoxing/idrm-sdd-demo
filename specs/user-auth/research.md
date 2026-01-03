# 用户认证功能技术研究

> **创建日期**: 2026-01-03  
> **状态**: 已完成

---

## 技术决策

### 1. JWT 令牌机制

**决策**: 使用双令牌机制（access_token + refresh_token）

**Rationale**:
- access_token 短期有效（2小时），降低令牌泄露风险
- refresh_token 长期有效（7天），减少用户频繁登录
- 刷新时仅刷新 access_token，简化实现和用户体验

**Alternatives considered**:
- 单令牌机制：安全性较低，令牌过期需要重新登录
- Session 机制：需要服务器端存储，不利于微服务架构

---

### 2. 密码加密算法

**决策**: 使用 bcrypt，cost factor >= 10

**Rationale**:
- bcrypt 是行业标准的密码哈希算法，抗暴力破解能力强
- cost factor >= 10 在安全性和性能之间取得平衡
- Go 标准库 `golang.org/x/crypto/bcrypt` 成熟稳定

**Alternatives considered**:
- SHA-256/SHA-512：单向哈希，但速度过快，不适合密码存储
- Argon2：更安全但性能开销较大，对当前场景过度设计

---

### 3. 验证码存储机制

**决策**: 使用 Redis 存储验证码，TTL = 5 分钟

**Rationale**:
- Redis 高性能，支持 TTL 自动过期
- 避免数据库压力，验证码属于临时数据
- 支持频率限制（60秒内只能发送一次）

**存储格式**:
- Key: `sms:code:{phone}:{type}` (type: register/reset)
- Value: 6位数字验证码
- TTL: 300秒（5分钟）

**频率限制**:
- Key: `sms:limit:{phone}`
- TTL: 60秒

---

### 4. 令牌存储机制

**决策**: 
- access_token: 不存储，JWT 自包含
- refresh_token: Redis 存储，用于令牌失效管理

**Rationale**:
- JWT 自包含特性，access_token 无需存储
- refresh_token 存储在 Redis 中，支持以下场景：
  - 密码修改/重置后使所有令牌失效
  - 用户主动登出
  - 令牌黑名单（可选）

**存储格式**:
- Key: `token:refresh:{userId}:{deviceId}`
- Value: refresh_token
- TTL: 7天

---

### 5. 账户锁定机制

**决策**: 连续5次登录失败锁定30分钟

**Rationale**:
- 平衡安全性和用户体验
- 5次失败后锁定可防止暴力破解
- 30分钟自动解锁，避免用户永久无法登录

**实现方式**:
- 使用数据库字段 `failed_attempts` 计数
- 使用 `locked_until` 记录锁定截止时间
- 成功登录后重置计数器

---

### 6. 短信服务集成

**决策**: 待确定具体短信服务提供商（阿里云/腾讯云/其他）

**Rationale**:
- 需要根据实际项目环境选择
- 建议选择支持国内手机号的云服务商
- 需要支持发送频率控制

**集成要求**:
- 支持发送6位数字验证码
- 支持发送频率限制
- 提供 SDK 或 HTTP API
- 需要配置 AccessKey/SecretKey

**Alternatives considered**:
- 第三方短信网关：成本可能较高
- 自建短信服务：开发维护成本高，不推荐

---

## 依赖库

| 库 | 版本 | 用途 |
|----|------|------|
| github.com/golang-jwt/jwt/v5 | latest | JWT 令牌生成和验证 |
| golang.org/x/crypto/bcrypt | latest | 密码加密 |
| gorm.io/gorm | latest | ORM（复杂查询） |
| github.com/jmoiron/sqlx | latest | SQLx（高性能查询） |
| github.com/redis/go-redis/v9 | latest | Redis 客户端 |

---

## 性能考虑

1. **密码加密**: bcrypt cost=10，单次加密耗时约 100-200ms，可接受
2. **验证码存储**: Redis 读写延迟 < 1ms，满足高频访问需求
3. **令牌验证**: JWT 无需数据库查询，性能优异
4. **登录历史**: 使用索引优化查询性能（user_id + login_time）

---

## 安全考虑

1. **密码安全**: bcrypt 加密，cost >= 10
2. **令牌安全**: access_token 短期有效，refresh_token 可撤销
3. **验证码安全**: 5分钟有效期，60秒发送频率限制
4. **账户安全**: 5次失败锁定，防止暴力破解
5. **SQL注入防护**: 使用 ORM/参数化查询

---

## 后续优化方向

1. **短信服务**: 确定具体服务提供商并集成
2. **令牌黑名单**: 考虑实现 refresh_token 黑名单机制
3. **设备管理**: 未来可扩展设备管理和强制下线功能
4. **多因素认证**: 后续版本可考虑添加 MFA 支持

