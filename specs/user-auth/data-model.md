# 用户认证功能数据模型

> **创建日期**: 2026-01-03  
> **状态**: Draft

---

## 实体定义

### 1. User（用户表）

**表名**: `users`

**字段定义**:

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | BIGINT UNSIGNED | PRIMARY KEY, AUTO_INCREMENT | 用户ID |
| phone | VARCHAR(11) | NOT NULL, UNIQUE | 手机号（11位数字） |
| password_hash | VARCHAR(255) | NOT NULL | 密码哈希（bcrypt） |
| nickname | VARCHAR(50) | NULL | 昵称（可选） |
| status | TINYINT | NOT NULL, DEFAULT 0 | 账户状态：0=正常，1=锁定，2=禁用 |
| failed_attempts | INT | NOT NULL, DEFAULT 0 | 连续登录失败次数 |
| locked_until | DATETIME | NULL | 锁定截止时间 |
| created_at | DATETIME | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| updated_at | DATETIME | NOT NULL, DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | 更新时间 |
| deleted_at | DATETIME | NULL | 软删除时间 |

**索引**:
- PRIMARY KEY (`id`)
- UNIQUE KEY `uk_phone` (`phone`)
- KEY `idx_status` (`status`)
- KEY `idx_deleted_at` (`deleted_at`)

**状态枚举**:
- `0`: 正常
- `1`: 锁定（30分钟内自动解锁）
- `2`: 禁用（需要管理员操作）

---

### 2. LoginHistory（登录历史表）

**表名**: `login_history`

**字段定义**:

| 字段名 | 类型 | 约束 | 说明 |
|--------|------|------|------|
| id | BIGINT UNSIGNED | PRIMARY KEY, AUTO_INCREMENT | 记录ID |
| user_id | BIGINT UNSIGNED | NOT NULL | 用户ID（外键） |
| login_time | DATETIME | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 登录时间 |
| ip_address | VARCHAR(45) | NOT NULL | IP地址（支持IPv6） |
| user_agent | VARCHAR(500) | NOT NULL | 浏览器/设备信息 |
| login_result | TINYINT | NOT NULL, DEFAULT 1 | 登录结果：1=成功，0=失败 |
| fail_reason | VARCHAR(255) | NULL | 失败原因 |
| created_at | DATETIME | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 创建时间 |

**索引**:
- PRIMARY KEY (`id`)
- KEY `idx_user_id_login_time` (`user_id`, `login_time`)
- KEY `idx_login_time` (`login_time`)

**登录结果枚举**:
- `1`: 成功
- `0`: 失败

**失败原因示例**:
- "手机号或密码错误"
- "账户已锁定"
- "账户已禁用"

---

## 关系定义

```
User (1) ──< (N) LoginHistory
```

- 一个用户可以有多次登录历史记录
- 登录历史通过 `user_id` 关联用户
- 登录历史保留90天（通过定时任务清理）

---

## Go Struct 定义

### User Struct

```go
type User struct {
    Id            int64          `gorm:"primaryKey" json:"id"`
    Phone         string         `gorm:"size:11;uniqueIndex;not null" json:"phone"`
    PasswordHash  string         `gorm:"size:255;not null" json:"-"`
    Nickname      *string        `gorm:"size:50" json:"nickname"`
    Status        int            `gorm:"type:tinyint;not null;default:0" json:"status"`
    FailedAttempts int           `gorm:"type:int;not null;default:0" json:"failed_attempts"`
    LockedUntil   *time.Time     `gorm:"type:datetime" json:"locked_until"`
    CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
    DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
```

### LoginHistory Struct

```go
type LoginHistory struct {
    Id          int64     `gorm:"primaryKey" json:"id"`
    UserId      int64     `gorm:"not null;index:idx_user_id_login_time" json:"user_id"`
    LoginTime   time.Time `gorm:"autoCreateTime;index:idx_user_id_login_time,idx_login_time" json:"login_time"`
    IpAddress   string    `gorm:"size:45;not null" json:"ip_address"`
    UserAgent   string    `gorm:"size:500;not null" json:"user_agent"`
    LoginResult int       `gorm:"type:tinyint;not null;default:1" json:"login_result"`
    FailReason  *string   `gorm:"size:255" json:"fail_reason,omitempty"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
```

---

## 数据验证规则

### User 验证规则

1. **手机号**: 
   - 必填
   - 长度 = 11
   - 格式：纯数字
   - 唯一性

2. **密码哈希**:
   - 必填
   - bcrypt 加密后长度约 60 字符

3. **昵称**:
   - 可选
   - 长度：1-50 字符

4. **状态**:
   - 默认值：0（正常）
   - 枚举值：0, 1, 2

### LoginHistory 验证规则

1. **user_id**: 
   - 必填
   - 必须存在对应的用户

2. **ip_address**:
   - 必填
   - 格式：IPv4 或 IPv6

3. **user_agent**:
   - 必填
   - 最大长度：500 字符

4. **login_result**:
   - 必填
   - 枚举值：0（失败），1（成功）

5. **fail_reason**:
   - 可选
   - 仅当 login_result = 0 时使用

---

## 数据迁移策略

1. **初始迁移**: 创建 users 和 login_history 表
2. **数据清理**: 定期清理90天前的登录历史记录（通过定时任务）
3. **索引优化**: 根据查询模式优化索引

---

## 数据访问模式

### User 访问模式

1. **按手机号查询**: 高频（登录、注册验证）
   - 使用索引: `uk_phone`

2. **按ID查询**: 高频（令牌验证、用户信息）
   - 使用索引: PRIMARY KEY

3. **更新密码**: 中频（密码修改、重置）
   - 使用索引: PRIMARY KEY

4. **更新锁定状态**: 低频（登录失败时）
   - 使用索引: PRIMARY KEY

### LoginHistory 访问模式

1. **按用户ID和时间范围查询**: 高频（登录历史列表）
   - 使用索引: `idx_user_id_login_time`

2. **插入记录**: 高频（每次登录）
   - 无索引依赖

3. **按时间清理**: 定时任务（每天执行）
   - 使用索引: `idx_login_time`

