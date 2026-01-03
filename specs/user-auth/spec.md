# 用户认证功能 Specification

> **Branch**: `feature/user-auth`  
> **Spec Path**: `specs/user-auth/`  
> **Created**: 2026-01-03  

> **Status**: Draft

---

## Overview

提供完整的用户认证功能，包括手机号注册、登录认证、密码管理和登录历史记录，确保用户账户安全和操作可追溯。

---

## Clarifications

### Session 2026-01-03

- Q: 注册时是否需要手机验证码验证? → A: 需要验证码：注册时发送短信验证码，用户输入验证码后才能完成注册
- Q: 是否支持密码重置/找回密码功能? → A: 支持密码重置：提供忘记密码功能，通过手机验证码验证身份后重置密码
- Q: 密码强度的具体要求是什么? → A: 高强度：至少8位，必须包含大小写字母、数字和特殊字符
- Q: 登录历史记录的查询范围和分页规则是什么? → A: 默认显示最近30天，支持自定义时间范围筛选，每页20条记录
- Q: 刷新令牌时，refresh_token 是否也需要刷新? → A: 仅刷新 access_token：refresh_token 保持不变，继续使用直到过期（7天）

---

## User Stories

### Story 1: 手机号注册 (P1)

AS a 新用户
I WANT 使用手机号、验证码和密码注册账户
SO THAT 获得系统访问权限

**独立测试**: 使用有效手机号和验证码验证后，设置密码完成注册，可以使用该账户登录系统

### Story 2: 登录认证 (P1)

AS a 注册用户
I WANT 使用手机号和密码登录系统
SO THAT 访问我的个人数据和功能

**独立测试**: 使用正确的手机号和密码登录后，返回有效的访问令牌

### Story 3: 密码管理 (P2)

AS a 注册用户
I WANT 修改我的登录密码
SO THAT 保障账户安全

**独立测试**: 使用旧密码验证后，成功设置新密码，新密码可用于登录

### Story 5: 密码重置 (P2)

AS a 忘记密码的用户
I WANT 通过手机验证码重置密码
SO THAT 在忘记旧密码时也能重新获得账户访问权限

**独立测试**: 通过手机号和验证码验证身份后，成功设置新密码，新密码可用于登录

### Story 4: 登录历史记录 (P2)

AS a 注册用户
I WANT 查看我的登录历史记录
SO THAT 发现异常登录行为

**独立测试**: 登录后可查看登录时间、IP 地址、设备信息等历史记录

---

## Acceptance Criteria (EARS)

### 正常流程

| ID | Scenario | Trigger | Expected Behavior |
|----|----------|---------|-------------------|
| AC-01 | 注册成功 | WHEN 用户提交有效手机号、验证码和密码 | THE SYSTEM SHALL 创建用户并返回 201 |
| AC-02 | 发送验证码 | WHEN 用户请求发送注册验证码 | THE SYSTEM SHALL 发送短信验证码并返回 200 |
| AC-03 | 验证码校验 | WHEN 用户提交正确的验证码 | THE SYSTEM SHALL 验证通过并允许继续注册流程 |
| AC-04 | 登录成功 | WHEN 用户提交正确的手机号和密码 | THE SYSTEM SHALL 返回 access_token 和 refresh_token |
| AC-05 | 刷新令牌 | WHEN 用户使用有效的 refresh_token | THE SYSTEM SHALL 返回新的 access_token，refresh_token 保持不变 |
| AC-06 | 修改密码成功 | WHEN 用户提交正确的旧密码和新密码 | THE SYSTEM SHALL 更新密码并返回 200 |
| AC-07 | 发送重置密码验证码 | WHEN 用户请求发送重置密码验证码 | THE SYSTEM SHALL 发送短信验证码并返回 200 |
| AC-08 | 重置密码成功 | WHEN 用户提交手机号、验证码和新密码 | THE SYSTEM SHALL 更新密码并返回 200 |
| AC-09 | 查询登录历史 | WHEN 用户请求登录历史列表（默认最近30天） | THE SYSTEM SHALL 返回分页的登录记录（每页20条） |
| AC-09a | 按时间范围查询 | WHEN 用户指定时间范围查询登录历史 | THE SYSTEM SHALL 返回指定时间范围内的分页登录记录 |
| AC-10 | 登出成功 | WHEN 用户请求登出 | THE SYSTEM SHALL 使当前令牌失效并返回 200 |

### 异常处理

| ID | Scenario | Trigger | Expected Behavior |
|----|----------|---------|-------------------|
| AC-11 | 注册验证码错误 | WHEN 用户提交错误的注册验证码 | THE SYSTEM SHALL 返回 400 "验证码错误" |
| AC-12 | 注册验证码过期 | WHEN 用户提交过期的注册验证码（超过有效期） | THE SYSTEM SHALL 返回 400 "验证码已过期" |
| AC-13 | 重置密码验证码错误 | WHEN 用户提交错误的重置密码验证码 | THE SYSTEM SHALL 返回 400 "验证码错误" |
| AC-14 | 重置密码验证码过期 | WHEN 用户提交过期的重置密码验证码 | THE SYSTEM SHALL 返回 400 "验证码已过期" |
| AC-15 | 手机号未注册 | WHEN 重置密码时手机号不存在 | THE SYSTEM SHALL 返回 404 "手机号未注册" |
| AC-16 | 手机号已注册 | WHEN 手机号已存在 | THE SYSTEM SHALL 返回 409 "手机号已注册" |
| AC-17 | 手机号格式错误 | WHEN 手机号不是 11 位数字 | THE SYSTEM SHALL 返回 400 "手机号格式错误" |
| AC-18 | 密码强度不足 | WHEN 密码不符合强度要求（少于8位或缺少大小写字母、数字、特殊字符中的任意一项） | THE SYSTEM SHALL 返回 400 "密码强度不足，必须包含大小写字母、数字和特殊字符，至少8位" |
| AC-19 | 登录失败 | WHEN 手机号或密码错误 | THE SYSTEM SHALL 返回 401 "用户名或密码错误" |
| AC-20 | 账户被锁定 | WHEN 连续 5 次登录失败 | THE SYSTEM SHALL 锁定账户 30 分钟并返回 423 |
| AC-21 | 令牌过期 | WHEN access_token 已过期 | THE SYSTEM SHALL 返回 401 "令牌已过期" |
| AC-22 | 旧密码错误 | WHEN 修改密码时旧密码不正确 | THE SYSTEM SHALL 返回 400 "旧密码错误" |

---

## Edge Cases

| ID | Case | Expected Behavior |
|----|------|-------------------|
| EC-01 | 并发多设备登录 | 允许多设备登录，每个设备有独立的 token |
| EC-02 | 密码修改后其他设备 | 其他设备的 token 立即失效，需要重新登录 |
| EC-08 | 密码重置后其他设备 | 密码重置后，所有设备的 token 立即失效，需要重新登录 |
| EC-09 | 重置密码验证码发送频率限制 | 同一手机号 60 秒内重复请求重置密码验证码，返回 429 "请求过于频繁" |
| EC-03 | 刷新令牌时 access_token 仍有效 | 返回新的 access_token，旧 token 仍可使用至过期 |
| EC-04 | 账户锁定期间尝试登录 | 返回锁定剩余时间，不重置锁定时间 |
| EC-05 | 连续登录失败计数重置 | 成功登录后重置错误计数器 |
| EC-06 | 验证码发送频率限制 | 同一手机号 60 秒内重复请求发送验证码，返回 429 "请求过于频繁" |
| EC-07 | 验证码过期后使用 | 使用过期的验证码注册，返回 400 "验证码已过期"，需重新获取 |
| EC-10 | 查询超出保留期的登录历史 | 用户查询超过90天的登录历史，返回空结果或提示超出查询范围 |
| EC-11 | 登录历史时间范围边界 | 用户查询的开始时间早于90天前，自动调整为90天前的起始时间 |

---

## Business Rules

| ID | Rule | Description |
|----|------|-------------|
| BR-01 | 手机号唯一 | 每个手机号只能注册一个账户 |
| BR-02 | 密码强度 | 至少 8 位，必须包含大小写字母、数字和特殊字符 |
| BR-03 | 令牌有效期 | access_token 2小时，refresh_token 7天 |
| BR-04 | 登录失败锁定 | 5 次失败锁定 30 分钟 |
| BR-05 | 密码加密存储 | 使用 bcrypt 加密存储密码 |
| BR-06 | 登录历史保留 | 保留最近 90 天的登录记录 |
| BR-10 | 登录历史查询默认范围 | 默认查询最近 30 天的登录记录 |
| BR-11 | 登录历史分页 | 每页显示 20 条记录，支持自定义时间范围筛选（不超过90天） |
| BR-07 | 验证码有效期 | 注册验证码和重置密码验证码有效期 5 分钟 |
| BR-08 | 验证码发送频率 | 同一手机号 60 秒内只能发送一次验证码（注册或重置密码） |
| BR-09 | 密码重置后令牌失效 | 密码重置后，所有设备的访问令牌立即失效 |

---

## Data Considerations

| Field | Description | Constraints |
|-------|-------------|-------------|
| phone | 手机号 | 必填，11 位数字，唯一 |
| password_hash | 密码哈希 | bcrypt 加密 |
| nickname | 昵称 | 可选，1-50 字符 |
| status | 账户状态 | 正常/锁定/禁用 |
| failed_attempts | 登录失败次数 | 整数，默认 0 |
| locked_until | 锁定截止时间 | 可空，datetime |
| created_at | 注册时间 | datetime |
| updated_at | 更新时间 | datetime |

**登录历史**:

| Field | Description | Constraints |
|-------|-------------|-------------|
| user_id | 用户 ID | 必填，外键 |
| login_time | 登录时间 | datetime |
| ip_address | IP 地址 | varchar(45) |
| user_agent | 浏览器/设备信息 | varchar(500) |
| login_result | 登录结果 | 成功/失败 |
| fail_reason | 失败原因 | 可空 |

---

## Success Metrics

| ID | Metric | Target |
|----|--------|--------|
| SC-01 | 注册接口响应时间 | < 500ms (P99) |
| SC-02 | 登录接口响应时间 | < 200ms (P99) |
| SC-03 | 密码校验安全级别 | bcrypt cost >= 10 |
| SC-04 | 测试覆盖率 | > 80% |

---

## API Endpoints (参考)

| Method | Path | Description |
|--------|------|-------------|
| POST | /api/v1/auth/send-code | 发送注册验证码 |
| POST | /api/v1/auth/register | 用户注册 |
| POST | /api/v1/auth/login | 用户登录 |
| POST | /api/v1/auth/send-reset-code | 发送重置密码验证码 |
| POST | /api/v1/auth/reset-password | 重置密码 |
| POST | /api/v1/auth/logout | 用户登出 |
| POST | /api/v1/auth/refresh | 刷新令牌 |
| PUT | /api/v1/auth/password | 修改密码 |
| GET | /api/v1/auth/login-history | 登录历史 |

---

## Open Questions

- [x] 是否支持邮箱登录? → 暂不支持，仅手机号
- [x] 是否需要验证码注册? → 需要验证码：注册时发送短信验证码，用户输入验证码后才能完成注册
- [ ] 是否支持第三方登录 (微信/支付宝)? → 后续版本

---

## Revision History

| Version | Date | Author | Changes |
|---------|------|--------|---------|
| 1.0 | 2026-01-03 | AI | 初始版本 |
