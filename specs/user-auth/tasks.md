# 用户认证功能 Tasks

> **Branch**: `20260103-user-auth`  
> **Spec Path**: `specs/user-auth/`  
> **Created**: 2026-01-03  
> **Input**: spec.md, plan.md, data-model.md, contracts/user_auth.api

---

## Task Format

```
- [ ] [TaskID] [P?] [Story?] Description with file path
```

| 标记 | 含义 |
|------|------|
| `T001` | 任务 ID |
| `[P]` | 可并行执行（不同文件，无依赖） |
| `[US1]` | 关联 User Story 1 |

---

## Task Overview

| ID | Task | Story | Status | Parallel | Est. Lines |
|----|------|-------|--------|----------|------------|
| T001-T005 | Setup & Foundation | Setup | ⏸️ | - | - |
| T006-T025 | User Story 1: 手机号注册 | US1 | ⏸️ | Yes | ~400 |
| T026-T045 | User Story 2: 登录认证 | US2 | ⏸️ | Yes | ~400 |
| T046-T055 | User Story 3: 密码管理 | US3 | ⏸️ | - | ~150 |
| T056-T065 | User Story 5: 密码重置 | US5 | ⏸️ | - | ~200 |
| T066-T075 | User Story 4: 登录历史记录 | US4 | ⏸️ | - | ~150 |
| T076-T080 | Polish & Cross-cutting | Polish | ⏸️ | - | - |

**总计**: 80 个任务

---

## Phase 1: Setup

**目的**: 项目初始化和基础配置

- [x] T001 确认 Go 1.24+ 和 Go-Zero v1.9+ 已安装
- [x] T002 [P] 确认 MySQL 8.0 和 Redis 7.0 已安装并运行
- [x] T003 [P] 确认 goctl 工具已安装：`go install github.com/zeromicro/go-zero/tools/goctl@latest`
- [x] T004 确认项目结构符合 Go-Zero 标准（api/, model/, migrations/ 目录存在）
- [x] T005 确认 `api/doc/base.api` 已定义通用类型

**Checkpoint**: ✅ 开发环境就绪

---

## Phase 2: Foundation

**目的**: 必须完成后才能开始 User Story 实现

- [x] T006 在 `api/etc/api.yaml` 中配置 Redis 连接信息
- [ ] T007 在 `api/internal/svc/servicecontext.go` 中添加 Redis 客户端初始化（需要先运行goctl生成基础代码）
- [x] T008 在 `api/etc/api.yaml` 中配置 JWT 密钥（AccessSecret 和 AccessExpire）- 已存在，AccessExpire: 7200秒(2小时)符合需求
- [x] T009 安装依赖包：`go get github.com/golang-jwt/jwt/v5 golang.org/x/crypto/bcrypt github.com/redis/go-redis/v9`
- [ ] T010 创建验证码服务工具类 `api/internal/utils/sms.go`（接口定义，具体实现待后续集成短信服务）

**Checkpoint**: ✅ 基础设施就绪（Redis、JWT配置、依赖包），可开始 User Story 实现

---

## Phase 3: User Story 1 - 手机号注册 (P1) 🎯 MVP

**目标**: 用户可以使用手机号、验证码和密码注册账户

**独立测试**: 使用有效手机号和验证码验证后，设置密码完成注册，可以使用该账户登录系统

### Step 1: 定义 API 文件

- [x] T011 [US1] 将 `contracts/user_auth.api` 复制到 `api/doc/auth/user_auth.api`
- [x] T012 [US1] 在 `api/doc/api.api` 中 import auth 模块：`import "auth/user_auth.api"`
- [x] T013 [US1] 验证 API 文件语法：`goctl api validate --api api/doc/api.api`

### Step 2: 生成代码

- [x] T014 [US1] 运行 goctl 生成 Handler/Types：`goctl api go -api api/doc/api.api -dir api/ --style=go_zero --type-group`
- [x] T015 [US1] 运行 goctl 生成 Swagger 文档：`goctl api swagger --api api/doc/api.api --dir api/swagger --filename api`

### Step 3: 定义 DDL

- [x] T016 [P] [US1] 创建 `migrations/auth/users.sql`，定义 users 表结构（包含 phone, password_hash, nickname, status, failed_attempts, locked_until 等字段）
- [ ] T017 [US1] 执行数据库迁移：`mysql -u root -p idrm_demo < migrations/auth/users.sql`

### Step 4: 实现 Model 层 - User

- [ ] T018 [US1] 创建 `model/auth/user/interface.go`，定义 UserModel 接口（Insert, FindOne, FindOneByPhone, Update, WithTx, Trans 等方法）
- [ ] T019 [P] [US1] 创建 `model/auth/user/types.go`，定义 User 结构体（包含所有字段和 gorm 标签）
- [ ] T020 [P] [US1] 创建 `model/auth/user/vars.go`，定义常量（用户状态枚举、错误码等）
- [ ] T021 [US1] 创建 `model/auth/user/factory.go`，实现 NewUserModel 工厂函数（返回 GORM 实现）
- [ ] T022 [US1] 实现 `model/auth/user/gorm_dao.go`，实现 UserModel 接口的所有方法

### Step 5: 配置 ServiceContext

- [ ] T023 [US1] 在 `api/internal/svc/servicecontext.go` 中添加 UserModel 字段和初始化代码

### Step 6: 实现 Logic 层 - 发送验证码

- [ ] T024 [US1] 实现 `api/internal/logic/auth/sendcode_logic.go`，包含验证码生成、Redis 存储、频率限制、短信发送调用（先实现 Mock）

### Step 7: 实现 Logic 层 - 注册

- [ ] T025 [US1] 实现 `api/internal/logic/auth/register_logic.go`，包含验证码校验、密码加密、用户创建、数据校验

**Checkpoint**: ✅ User Story 1 可独立测试和验证（可以发送验证码并完成注册）

---

## Phase 4: User Story 2 - 登录认证 (P1) 🎯 MVP

**目标**: 用户可以使用手机号和密码登录系统，获得访问令牌

**独立测试**: 使用正确的手机号和密码登录后，返回有效的访问令牌

### Step 1: 定义 DDL - LoginHistory

- [ ] T026 [P] [US2] 创建 `migrations/auth/login_history.sql`，定义 login_history 表结构
- [ ] T027 [US2] 执行数据库迁移：`mysql -u root -p idrm_demo < migrations/auth/login_history.sql`

### Step 2: 实现 Model 层 - LoginHistory

- [ ] T028 [US2] 创建 `model/auth/login_history/interface.go`，定义 LoginHistoryModel 接口（Insert, FindByUserId, WithTx 等方法）
- [ ] T029 [P] [US2] 创建 `model/auth/login_history/types.go`，定义 LoginHistory 结构体
- [ ] T030 [P] [US2] 创建 `model/auth/login_history/vars.go`，定义常量（登录结果枚举等）
- [ ] T031 [US2] 创建 `model/auth/login_history/factory.go`，实现 NewLoginHistoryModel 工厂函数
- [ ] T032 [US2] 实现 `model/auth/login_history/gorm_dao.go`，实现 LoginHistoryModel 接口的所有方法

### Step 3: 配置 ServiceContext

- [ ] T033 [US2] 在 `api/internal/svc/servicecontext.go` 中添加 LoginHistoryModel 字段和初始化代码

### Step 4: 实现 JWT 工具

- [ ] T034 [US2] 创建 `api/internal/utils/jwt.go`，实现 JWT 令牌生成和验证函数（GenerateToken, ParseToken, GenerateRefreshToken）

### Step 5: 实现 Logic 层 - 登录

- [ ] T035 [US2] 实现 `api/internal/logic/auth/login_logic.go`，包含密码验证、账户锁定检查、登录失败计数、JWT 令牌生成、登录历史记录

### Step 6: 实现 Logic 层 - 刷新令牌

- [ ] T036 [US2] 实现 `api/internal/logic/auth/refresh_logic.go`，包含 refresh_token 验证、新 access_token 生成

### Step 7: 实现 Logic 层 - 登出

- [ ] T037 [US2] 实现 `api/internal/logic/auth/logout_logic.go`，包含令牌失效处理（Redis 黑名单或删除 refresh_token）

**Checkpoint**: ✅ User Story 2 可独立测试和验证（可以登录、刷新令牌、登出）

---

## Phase 5: User Story 3 - 密码管理 (P2)

**目标**: 用户可以使用旧密码验证后修改登录密码

**独立测试**: 使用旧密码验证后，成功设置新密码，新密码可用于登录

### Implementation

- [ ] T038 [US3] 在 `model/auth/user/interface.go` 中添加 UpdatePassword 方法定义
- [ ] T039 [US3] 在 `model/auth/user/gorm_dao.go` 中实现 UpdatePassword 方法
- [ ] T040 [US3] 实现 `api/internal/logic/auth/updatepassword_logic.go`，包含旧密码验证、新密码加密、密码更新、所有设备令牌失效处理

**Checkpoint**: ✅ User Story 3 可独立测试和验证（可以修改密码，修改后需要重新登录）

---

## Phase 6: User Story 5 - 密码重置 (P2)

**目标**: 用户可以通过手机验证码重置密码

**独立测试**: 通过手机号和验证码验证身份后，成功设置新密码，新密码可用于登录

### Implementation

- [ ] T041 [US5] 实现 `api/internal/logic/auth/sendresetcode_logic.go`，包含手机号存在性验证、验证码生成和发送（复用发送验证码逻辑）
- [ ] T042 [US5] 实现 `api/internal/logic/auth/resetpassword_logic.go`，包含验证码校验、手机号验证、密码重置、所有设备令牌失效处理

**Checkpoint**: ✅ User Story 5 可独立测试和验证（可以重置密码，重置后需要重新登录）

---

## Phase 7: User Story 4 - 登录历史记录 (P2)

**目标**: 用户可以查看自己的登录历史记录

**独立测试**: 登录后可查看登录时间、IP 地址、设备信息等历史记录

### Implementation

- [ ] T043 [US4] 实现 `api/internal/logic/auth/loginhistory_logic.go`，包含用户ID获取（从JWT）、时间范围处理（默认30天）、分页查询、数据格式化

**Checkpoint**: ✅ User Story 4 可独立测试和验证（可以查询登录历史，支持时间范围和分页）

---

## Phase 8: Polish & Cross-cutting Concerns

**目的**: 收尾工作、代码质量、错误处理完善

### Code Quality

- [ ] T044 运行代码格式化：`go fmt ./...`
- [ ] T045 运行静态检查：`golangci-lint run ./api/... ./model/...`
- [ ] T046 检查所有错误处理是否完善（所有 error 返回值都已处理）
- [ ] T047 检查所有公开函数和类型是否有中文注释

### Testing

- [ ] T048 为 `model/auth/user/gorm_dao.go` 编写单元测试 `model/auth/user/gorm_dao_test.go`，覆盖率 > 80%
- [ ] T049 [P] 为 `model/auth/login_history/gorm_dao.go` 编写单元测试 `model/auth/login_history/gorm_dao_test.go`，覆盖率 > 80%
- [ ] T050 [P] 为核心 Logic 层（login_logic.go, register_logic.go）编写单元测试，覆盖率 > 80%

### Integration & Documentation

- [ ] T051 集成真实的短信服务（替换 Mock 实现），更新 `api/internal/utils/sms.go`
- [ ] T052 更新 `api/etc/api.yaml` 中的短信服务配置（AccessKey, SecretKey 等）
- [ ] T053 编写 API 测试脚本或 Postman 集合，覆盖所有端点
- [ ] T054 更新 README.md，添加认证模块使用说明
- [ ] T055 运行完整测试套件：`go test ./... -cover`，确保覆盖率 > 80%

**Checkpoint**: ✅ 代码质量达标，测试覆盖充分，文档完整

---

## Dependencies

```
Phase 1 (Setup)
    ↓
Phase 2 (Foundation)
    ↓
Phase 3 (US1: 注册) ──┐
    ↓                  │
Phase 4 (US2: 登录) ───┼──→ Phase 5 (US3: 密码管理)
    ↓                  │         ↓
    └──────────────────┼──→ Phase 6 (US5: 密码重置)
                       │         ↓
                       └──→ Phase 7 (US4: 登录历史)
                                  ↓
                            Phase 8 (Polish)
```

### 用户故事依赖关系

- **US1 (注册)** → **US2 (登录)**: 登录需要已注册的用户
- **US2 (登录)** → **US3 (密码管理)**、**US4 (登录历史)**、**US5 (密码重置)**: 这些功能需要用户已登录（JWT认证）
- **US3 (密码管理)** 和 **US5 (密码重置)** 可并行实现（都涉及密码更新）
- **US4 (登录历史)** 依赖 US2（登录时记录历史）

### 并行执行说明

- **Phase 2**: T006-T010 中，T007-T009 可并行（配置不同部分）
- **Phase 3**: 
  - T019-T020 可并行（不同文件）
  - T016 可与 T018-T022 并行（DDL 和 Model 实现无依赖）
- **Phase 4**: 
  - T029-T030 可并行（不同文件）
  - T026-T027 可与 T028-T032 并行（DDL 和 Model 实现无依赖）
- **Phase 8**: T049-T050 可并行（不同测试文件）

---

## Implementation Strategy

### MVP Scope (最小可用产品)

**建议 MVP 包含**:
- ✅ User Story 1: 手机号注册（包含验证码）
- ✅ User Story 2: 登录认证（包含 JWT 令牌、刷新令牌、登出）

**MVP 可独立运行和测试**，用户可以完成完整的注册-登录流程。

### 增量交付

1. **Sprint 1 (MVP)**: Phase 1-4
   - 完成注册和登录功能
   - 用户可以注册账户并登录系统
   
2. **Sprint 2**: Phase 5-6
   - 完成密码管理和密码重置
   - 用户可以修改密码和找回密码
   
3. **Sprint 3**: Phase 7-8
   - 完成登录历史记录
   - 代码质量提升和测试完善

---

## Notes

- 每个 Task 完成后提交代码
- 每个 Checkpoint 进行验证
- 遇到问题及时记录到 Open Questions
- 短信服务可以先使用 Mock 实现，后续再集成真实服务
- JWT 配置中的 AccessSecret 和 AccessExpire 需要与业务规则一致（2小时、7天）

