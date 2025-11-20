# 充值PRD更新数据表
## 概述
本文档以及本文件夹是对充值PRD文档以及公司官网相关需求中需要添加的功能和需求进行的数据库表的设计和设计说明。

## 数据库结构

### 更新后的表关系图
```
users (用户表)
  ├── user_subscriptions (订阅配额表) [user_id] 1:1
  ├── teams (团队表) [owner_id] 1:n
  │   └── team_members (团队成员表) [team_id, user_id] n:n
  ├── billing_records (账单记录表) [user_id] 1:n
  ├── traffic_logs (流量日志表) [user_id] 1:n
  ├── projects (项目表) [user_id] 1:n
  │   ├── materials (材料表) [project_id] 1:n
  │   ├── pipelines (流水线表) [project_id] 1:n
  │   │   └── tasks (任务表) [pipeline_id] 1:n
  │   ├── views (收藏表) [user_id, project_id] n:n
  │   ├── comments (评论表) [user_id, project_id] n:n
  │   │   └── comment_likes (评论点赞表) [user_id, comment_id] n:n
  │   └── likes (点赞表) [user_id, project_id] n:n
  └── graphs (图谱表) [独立表，被projects引用]
```

### 表结构说明
**1. 用户订阅表 user_subscriptions**
```sql
-- 创建枚举类型
CREATE TYPE user_plan_enum AS ENUM ('free', 'professional', 'team');
CREATE TYPE subscription_status_enum AS ENUM ('active', 'expired', 'cancelled', 'suspended');

CREATE TABLE user_subscriptions
(
    id                       BIGSERIAL PRIMARY KEY, -- 自增ID
    user_id                  BIGINT                   NOT NULL UNIQUE,
    team_id                  BIGINT                            DEFAULT NULL,

    -- 套餐类型和订阅状态
    user_plan                user_plan_enum           NOT NULL DEFAULT 'free',
    subscription_status      subscription_status_enum NOT NULL DEFAULT 'active',

    -- 用量统计（流量的单位为MB）
    words_used               INT                      NOT NULL DEFAULT 0,
    storage_used             INT                      NOT NULL DEFAULT 0,
    cu_used                  INT                      NOT NULL DEFAULT 0,
    traffic_used             NUMERIC(10, 5)           NOT NULL DEFAULT 0.00000,

    -- 统一重置日期
    quota_reset_date         DATE                     NOT NULL,

    -- 用量预警
    words_warning_80_sent    BOOLEAN                           DEFAULT FALSE,
    words_warning_100_sent   BOOLEAN                           DEFAULT FALSE,
    storage_warning_80_sent  BOOLEAN                           DEFAULT FALSE,
    storage_warning_100_sent BOOLEAN                           DEFAULT FALSE,
    cu_warning_80_sent       BOOLEAN                           DEFAULT FALSE,
    cu_warning_100_sent      BOOLEAN                           DEFAULT FALSE,
    traffic_warning_80_sent  BOOLEAN                           DEFAULT FALSE,
    traffic_warning_100_sent BOOLEAN                           DEFAULT FALSE,

    -- 超额费用
    overage_words_fee        NUMERIC(10, 2)           NOT NULL DEFAULT 0.00,
    overage_storage_fee      NUMERIC(10, 2)           NOT NULL DEFAULT 0.00,
    overage_traffic_fee      NUMERIC(10, 2)           NOT NULL DEFAULT 0.00,
    overage_cu_fee           NUMERIC(10, 2)           NOT NULL DEFAULT 0.00,
    total_overage_fee        NUMERIC(10, 2)           NOT NULL DEFAULT 0.00,

    -- 用户选择项
    selected_ai_model        VARCHAR(50)                       DEFAULT 'GPT-4o',

    -- 时间戳
    created_at               TIMESTAMPTZ                       DEFAULT CURRENT_TIMESTAMP,
    updated_at               TIMESTAMPTZ                       DEFAULT CURRENT_TIMESTAMP,

    -- 外键
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_team FOREIGN KEY (team_id) REFERENCES teams (id) ON DELETE SET NULL
);

-- 索引
CREATE INDEX idx_plan_status ON user_subscriptions (user_plan, subscription_status);
CREATE INDEX idx_user_reset_date ON user_subscriptions (user_id, quota_reset_date);
-- next_billing_date 字段在原表未定义，如果需要请先添加
-- CREATE INDEX idx_billing_date ON user_subscriptions(next_billing_date);
CREATE INDEX idx_team_id ON user_subscriptions (team_id);

-- 添加表注释
COMMENT
ON TABLE user_subscriptions IS '用户订阅配额表';

```

**2. 团队表 teams**
```sql
-- 创建枚举类型
CREATE TYPE team_status_enum AS ENUM ('active', 'suspended', 'deleted');

CREATE TABLE teams
(
    id                  BIGSERIAL PRIMARY KEY,                    -- 团队ID
    team_name           VARCHAR(100)   NOT NULL,                  -- 团队名称
    owner_id            BIGINT         NOT NULL,                  -- 团队所有者用户ID
    team_code           VARCHAR(50) UNIQUE,                       -- 团队唯一编码
    invite_code         VARCHAR(6) UNIQUE NOT NULL DEFAULT (
        substring(
                upper(
                        translate(md5(random()::text), 'abcdefghijklmnopqrstuvwxyz', '01234567890123456789012345')
                )
                from 1 for 6
        )
        ),

    -- ==================== 团队配额（共享池） ====================
    total_words_quota   INT            NOT NULL DEFAULT 900000,   -- 团队总字数配额(默认3人*30万)
    total_storage_quota INT            NOT NULL DEFAULT 512000,   -- 团队总存储(500GB=512000MB)
    total_cu_quota      INT            NOT NULL DEFAULT 2000,     -- 团队总CU配额
    total_traffic_quota INT            NOT NULL DEFAULT 200,      -- 团队总流量

    -- ==================== 团队使用量 ====================
    words_used          INT            NOT NULL DEFAULT 0,        -- 团队已用字数
    storage_used        INT            NOT NULL DEFAULT 0,        -- 团队已用存储
    cu_used             INT            NOT NULL DEFAULT 0,        -- 团队已用CU
    traffic_used        NUMERIC(10, 3) NOT NULL DEFAULT 0.000,    -- 团队已用流量

    -- ==================== 成员管理 ====================
    member_count        INT            NOT NULL DEFAULT 3,        -- 当前成员数

    -- ==================== 团队状态 ====================
    status              team_status_enum        DEFAULT 'suspended', -- 团队状态，支付后变成active

    -- ==================== 时间戳 ====================
    created_at          TIMESTAMPTZ             DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ             DEFAULT CURRENT_TIMESTAMP,

    -- ==================== 外键 ====================
    CONSTRAINT fk_owner FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE
);

-- ==================== 索引 ====================
CREATE INDEX idx_owner ON teams (owner_id);
CREATE INDEX idx_team_code ON teams (team_code);
CREATE INDEX idx_status ON teams (status);
CREATE INDEX idx_team_invite_code ON teams (invite_code);

-- ==================== 表注释 ====================
COMMENT
ON TABLE teams IS '团队表';

```

**3. 团队成员表 team_members**
```sql
-- ==================== 创建枚举类型 ====================
CREATE TYPE role_enum AS ENUM ('owner', 'admin', 'member');
CREATE TYPE status_enum AS ENUM ('active', 'pending', 'removed');

-- ==================== 创建表 ====================
CREATE TABLE team_members
(
    id                    BIGSERIAL PRIMARY KEY,   -- 成员关系ID
    team_id               BIGINT         NOT NULL, -- 团队ID
    user_id               BIGINT         NOT NULL, -- 用户ID

    -- 成员角色
    role                  role_enum      NOT NULL DEFAULT 'member' COMMENT '角色：所有者/管理员/成员',

    -- 个人分配配额（管理员可调整）
    allocated_words_quota INT,                     -- NULL表示不限制

    -- 个人使用统计
    personal_words_used   INT            NOT NULL DEFAULT 0,
    personal_storage_used INT            NOT NULL DEFAULT 0,
    personal_cu_used      INT            NOT NULL DEFAULT 0,
    personal_traffic_used NUMERIC(10, 3) NOT NULL DEFAULT 0.000,

    -- 成员状态
    status                status_enum             DEFAULT 'active',
    invite_code           VARCHAR(100),
    invited_by            BIGINT,

    -- 时间戳
    joined_at             TIMESTAMPTZ             DEFAULT CURRENT_TIMESTAMP,
    removed_at            TIMESTAMPTZ,

    -- 外键约束
    CONSTRAINT fk_team FOREIGN KEY (team_id) REFERENCES teams (id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_invited_by FOREIGN KEY (invited_by) REFERENCES users (id) ON DELETE SET NULL,

    -- 唯一约束
    CONSTRAINT uk_team_user UNIQUE (team_id, user_id)
);

-- ==================== 索引 ====================
CREATE INDEX idx_user_status ON team_members (user_id, status);
CREATE INDEX idx_role ON team_members (role);
CREATE INDEX idx_invite_code ON team_members (invite_code);

```

**4. 核心账单记录表 billing_records**
```sql
-- ==================== 创建枚举类型 ====================
CREATE TYPE billing_type_enum AS ENUM ('subscription', 'overage', 'refund', 'month');
CREATE TYPE billing_status_enum AS ENUM ('unpaid', 'paid');
-- ==================== 创建表 ====================
CREATE TABLE billing_records
(
    id                    BIGSERIAL PRIMARY KEY,                             -- 账单ID
    user_id               BIGINT         NOT NULL,                           -- 用户ID
    team_id               BIGINT,                                            -- 团队ID（团队账单）

    -- 账期信息
    billing_period        VARCHAR(7)     NOT NULL,                           -- 账期 YYYY-MM
    billing_date          DATE           NOT NULL,                           -- 结算日期
    billing_type          billing_type_enum       DEFAULT 'subscription',    -- 账单类型

    -- 费用信息
    base_subscription_fee NUMERIC(10, 2) NOT NULL DEFAULT 0.00,              -- 基础订阅费
    overage_fee           NUMERIC(10, 2) NOT NULL DEFAULT 0.00,              -- 超额费用总额
    subtotal              NUMERIC(10, 2) NOT NULL,                           -- 小计（基础费+超额费）
    discount_amount       NUMERIC(10, 2)          DEFAULT 0.00,              -- 折扣金额
    total_amount          NUMERIC(10, 2) NOT NULL,                           -- 应付总金额

    -- 支付状态
    status                billing_status_enum NOT NULL DEFAULT 'unpaid' COMMENT '账单状态：未支付/已支付',

    -- 备注与时间戳
    remark                TEXT,                                              -- 备注说明
    created_at            TIMESTAMPTZ             DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    updated_at            TIMESTAMPTZ             DEFAULT CURRENT_TIMESTAMP, -- 更新时间

-- ==================== 表与列注释（可选） ====================
COMMENT
ON TABLE billing_records IS '账单核心信息表';
COMMENT
ON COLUMN billing_records.billing_period IS '账期 YYYY-MM';
COMMENT
ON COLUMN billing_records.billing_date IS '结算日期';
COMMENT
ON COLUMN billing_records.billing_type IS '账单类型';
COMMENT
ON COLUMN billing_records.base_subscription_fee IS '基础订阅费';
COMMENT
ON COLUMN billing_records.overage_fee IS '超额费用总额';
COMMENT
ON COLUMN billing_records.subtotal IS '小计（基础费+超额费）';
COMMENT
ON COLUMN billing_records.discount_amount IS '折扣金额';
COMMENT
ON COLUMN billing_records.total_amount IS '应付总金额';
COMMENT
ON COLUMN billing_records.remark IS '备注说明';
COMMENT
ON COLUMN billing_records.created_at IS '创建时间';
COMMENT
ON COLUMN billing_records.updated_at IS '更新时间';

```

**5. 支付记录表 billing_payments**
```sql
-- ==================== 创建枚举类型 ====================
CREATE TYPE payment_status_enum AS ENUM ('pending', 'paid', 'failed', 'refunded', 'cancelled');

-- ==================== 创建表 ====================
CREATE TABLE billing_payments
(
    id                     BIGSERIAL PRIMARY KEY,                         -- 支付记录ID
    billing_id             BIGINT         NOT NULL,                       -- 关联账单ID

    -- 支付信息
    payment_status         payment_status_enum DEFAULT 'pending',         -- 支付状态
    payment_method         VARCHAR(50)    NOT NULL,                       -- 支付方式: alipay/wechat/credit_card
    payment_amount         NUMERIC(10, 2) NOT NULL,                       -- 本次支付金额
    payment_transaction_id VARCHAR(100),                                  -- 支付平台流水号（第三方返回）
    payment_channel        VARCHAR(50),                                   -- 支付渠道: 微信支付/支付宝/银联

    -- 支付结果信息
    paid_at                TIMESTAMPTZ,                                   -- 支付成功时间
    failure_reason         VARCHAR(500),                                  -- 支付失败原因
    refunded_at            TIMESTAMPTZ,                                   -- 退款时间（如有）

    -- 时间戳
    created_at             TIMESTAMPTZ         DEFAULT CURRENT_TIMESTAMP, -- 创建时间（发起支付时）
    updated_at             TIMESTAMPTZ         DEFAULT CURRENT_TIMESTAMP, -- 更新时间（状态变更时）

    -- 外键约束
    CONSTRAINT fk_billing FOREIGN KEY (billing_id) REFERENCES billing_records (id) ON DELETE CASCADE
);

-- ==================== 索引 ====================
CREATE INDEX idx_billing ON billing_payments (billing_id);
CREATE INDEX idx_status_time ON billing_payments (payment_status, created_at);
CREATE INDEX idx_transaction_id ON billing_payments (payment_transaction_id);

-- ==================== 表与列注释 ====================
COMMENT
ON TABLE billing_payments IS '账单支付记录表';
COMMENT
ON COLUMN billing_payments.payment_status IS '支付状态';
COMMENT
ON COLUMN billing_payments.payment_method IS '支付方式: alipay/wechat/credit_card';
COMMENT
ON COLUMN billing_payments.payment_amount IS '本次支付金额';
COMMENT
ON COLUMN billing_payments.payment_transaction_id IS '支付平台流水号（第三方返回）';
COMMENT
ON COLUMN billing_payments.payment_channel IS '支付渠道: 微信支付/支付宝/银联';
COMMENT
ON COLUMN billing_payments.paid_at IS '支付成功时间';
COMMENT
ON COLUMN billing_payments.failure_reason IS '支付失败原因';
COMMENT
ON COLUMN billing_payments.refunded_at IS '退款时间（如有）';
COMMENT
ON COLUMN billing_payments.created_at IS '创建时间（发起支付时）';
COMMENT
ON COLUMN billing_payments.updated_at IS '更新时间（状态变更时）';

```

**6. 超额费用明细表 billing_overages**
```sql
-- ==================== 创建表 ====================
CREATE TABLE billing_overages
(
    id                     BIGSERIAL PRIMARY KEY,                    -- 明细ID
    billing_id             BIGINT NOT NULL,                          -- 关联账单ID

    -- 四种资源的超额用量和费用
    words_overage_amount   NUMERIC(10, 3) DEFAULT 0.000,             -- 字数超额量（单位：千字）
    words_overage_fee      NUMERIC(10, 2) DEFAULT 0.00,              -- 字数超额费用

    storage_overage_amount NUMERIC(10, 3) DEFAULT 0.000,             -- 存储超额量（单位：GB）
    storage_overage_fee    NUMERIC(10, 2) DEFAULT 0.00,              -- 存储超额费用

    traffic_overage_amount NUMERIC(10, 3) DEFAULT 0.000,             -- 流量超额量（单位：GB）
    traffic_overage_fee    NUMERIC(10, 2) DEFAULT 0.00,              -- 流量超额费用

    cu_overage_amount      NUMERIC(10, 3) DEFAULT 0.000,             -- 计算单元超额量（单位：CU）
    cu_overage_fee         NUMERIC(10, 2) DEFAULT 0.00,              -- 计算单元超额费用

    -- 汇总信息（冗余字段）
    total_overage_fee      NUMERIC(10, 2) DEFAULT 0.00,              -- 超额费用汇总

    -- 时间戳
    created_at             TIMESTAMPTZ    DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    updated_at             TIMESTAMPTZ    DEFAULT CURRENT_TIMESTAMP, -- 更新时间

    -- 外键约束
    CONSTRAINT fk_billing FOREIGN KEY (billing_id) REFERENCES billing_records (id) ON DELETE CASCADE,

    -- 唯一约束
    CONSTRAINT uk_billing UNIQUE (billing_id)
);

-- ==================== 索引 ====================
CREATE INDEX idx_billing ON billing_overages (billing_id);

-- ==================== 表与列注释 ====================
COMMENT
ON TABLE billing_overages IS '账单超额费用明细表';
COMMENT
ON COLUMN billing_overages.words_overage_amount IS '字数超额量（单位：千字）';
COMMENT
ON COLUMN billing_overages.words_overage_fee IS '字数超额费用';
COMMENT
ON COLUMN billing_overages.storage_overage_amount IS '存储超额量（单位：GB）';
COMMENT
ON COLUMN billing_overages.storage_overage_fee IS '存储超额费用';
COMMENT
ON COLUMN billing_overages.traffic_overage_amount IS '流量超额量（单位：GB）';
COMMENT
ON COLUMN billing_overages.traffic_overage_fee IS '流量超额费用';
COMMENT
ON COLUMN billing_overages.cu_overage_amount IS '计算单元超额量（单位：CU）';
COMMENT
ON COLUMN billing_overages.cu_overage_fee IS '计算单元超额费用';
COMMENT
ON COLUMN billing_overages.total_overage_fee IS '超额费用汇总';
COMMENT
ON COLUMN billing_overages.created_at IS '创建时间';
COMMENT
ON COLUMN billing_overages.updated_at IS '更新时间';

```

**7. 发票信息表 billing_invoices**
```sql
-- ==================== 创建表 ====================
CREATE TABLE billing_invoices
(
    id                BIGSERIAL PRIMARY KEY,                 -- 发票ID
    billing_id        BIGINT NOT NULL UNIQUE,                -- 关联账单ID（一对一）

    -- 发票信息
    invoice_required  BOOLEAN     DEFAULT FALSE,             -- 是否需要发票
    invoice_title     VARCHAR(200),                          -- 发票抬头
    invoice_tax_id    VARCHAR(50),                           -- 税号
    invoice_url       VARCHAR(500),                          -- 发票下载链接
    invoice_issued_at TIMESTAMPTZ,                           -- 发票开具时间

    -- 时间戳
    created_at        TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    updated_at        TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, -- 更新时间

    -- 外键约束
    CONSTRAINT fk_billing FOREIGN KEY (billing_id) REFERENCES billing_records (id) ON DELETE CASCADE
);

-- ==================== 索引 ====================
CREATE INDEX idx_invoice_issued ON billing_invoices (invoice_issued_at);
CREATE INDEX idx_invoice_title ON billing_invoices (invoice_title);

-- ==================== 表与列注释 ====================
COMMENT
ON TABLE billing_invoices IS '账单发票信息表';
COMMENT
ON COLUMN billing_invoices.invoice_required IS '是否需要发票';
COMMENT
ON COLUMN billing_invoices.invoice_title IS '发票抬头';
COMMENT
ON COLUMN billing_invoices.invoice_tax_id IS '税号';
COMMENT
ON COLUMN billing_invoices.invoice_url IS '发票下载链接';
COMMENT
ON COLUMN billing_invoices.invoice_issued_at IS '发票开具时间';
COMMENT
ON COLUMN billing_invoices.created_at IS '创建时间';
COMMENT
ON COLUMN billing_invoices.updated_at IS '更新时间';

```

### 添加存储过程
**1. 创建更新账单总额的函数**
```sql
-- 创建更新账单总额的 PL/pgSQL 函数
CREATE OR REPLACE FUNCTION update_billing_totals()
RETURNS TRIGGER AS $$
DECLARE
    total_overage DECIMAL(10,2);
    billing_record RECORD;
BEGIN
    -- 确定受影响的 billing_id
    DECLARE
        affected_billing_id BIGINT;
    BEGIN
        IF TG_OP = 'DELETE' THEN
            affected_billing_id := OLD.billing_id;
        ELSE
            affected_billing_id := NEW.billing_id;
        END IF;
        
        -- 计算超额费用总额
        SELECT COALESCE(SUM(
            COALESCE(words_overage_fee, 0) + 
            COALESCE(storage_overage_fee, 0) + 
            COALESCE(traffic_overage_fee, 0) + 
            COALESCE(cu_overage_fee, 0)
        ), 0)
        INTO total_overage
        FROM billing_overages 
        WHERE billing_id = affected_billing_id;
        
        -- 获取账单记录
        SELECT * INTO billing_record 
        FROM billing_records 
        WHERE id = affected_billing_id;
        
        IF FOUND THEN
            -- 更新账单记录
            UPDATE billing_records 
            SET overage_fee = total_overage,
                subtotal = billing_record.base_subscription_fee + total_overage,
                total_amount = billing_record.base_subscription_fee + total_overage - COALESCE(billing_record.discount_amount, 0),
                updated_at = CURRENT_TIMESTAMP
            WHERE id = affected_billing_id;
        END IF;
        
    END;
    
    -- 对于触发器，返回适当的记录
    IF TG_OP = 'DELETE' THEN
        RETURN OLD;
    ELSE
        RETURN NEW;
    END IF;
END;
$$ LANGUAGE plpgsql;
```

**2. 为超额明细表创建触发器**
```sql
-- 为插入、更新、删除操作创建触发器
CREATE TRIGGER trigger_update_billing_totals_insert
    AFTER INSERT ON billing_overages
    FOR EACH ROW
    EXECUTE FUNCTION update_billing_totals();

CREATE TRIGGER trigger_update_billing_totals_update
    AFTER UPDATE ON billing_overages
    FOR EACH ROW
    EXECUTE FUNCTION update_billing_totals();

CREATE TRIGGER trigger_update_billing_totals_delete
    AFTER DELETE ON billing_overages
    FOR EACH ROW
    EXECUTE FUNCTION update_billing_totals();
```

3. **为超额明细表添加计算字段的触发器**
```sql
-- 创建计算汇总字段的函数
CREATE OR REPLACE FUNCTION calculate_overage_totals()
RETURNS TRIGGER AS $$
BEGIN
    -- 自动计算汇总字段
    NEW.total_overage_amount := 
        COALESCE(NEW.words_overage_amount, 0) + 
        COALESCE(NEW.storage_overage_amount, 0) + 
        COALESCE(NEW.traffic_overage_amount, 0) + 
        COALESCE(NEW.cu_overage_amount, 0);
        
    NEW.total_overage_fee := 
        COALESCE(NEW.words_overage_fee, 0) + 
        COALESCE(NEW.storage_overage_fee, 0) + 
        COALESCE(NEW.traffic_overage_fee, 0) + 
        COALESCE(NEW.cu_overage_fee, 0);
        
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 为插入和更新创建触发器
CREATE TRIGGER trigger_calculate_overage_totals_insert
    BEFORE INSERT ON billing_overages
    FOR EACH ROW
    EXECUTE FUNCTION calculate_overage_totals();

CREATE TRIGGER trigger_calculate_overage_totals_update
    BEFORE UPDATE ON billing_overages
    FOR EACH ROW
    EXECUTE FUNCTION calculate_overage_totals();
```

### 更新后的所有表结构
| 模块    | 主要表                                                                           | 功能说明            |
| ----- |-------------------------------------------------------------------------------| --------------- |
| 用户管理  | `users`                                                                       | 管理用户账号信息        |
| 项目系统  | `projects`, `graphs`, `materials`, `pipelines`, `tasks`                       | 用户项目、图谱、任务与材料关联 |
| 社交互动  | `views`, `likes`, `comments`, `comment_likes`                                 | 收藏、点赞与评论功能      |
| 团队系统  | `teams`, `team_members`                                                       | 团队、成员与配额控制      |
| 订阅与配额 | `user_subscriptions` ,`traffic_logs`                                         | 用户套餐与资源用量统计     |
| 计费与发票 | `billing_records`, `billing_payments`, `billing_overages`, `billing_invoices` | 账单、支付、超额与发票管理   |


## 安装和使用

### 前置条件
- PostgreSQL 12+ 已安装并运行
- 具有创建数据库权限的用户账户

### 快速开始

#### Linux/macOS
```bash
# 给脚本执行权限
chmod +x database/init.sh

# 运行初始化脚本
./database/init.sh
```

#### Windows
```cmd
# 双击运行或在命令行执行
database\init.bat
```

#### 手动执行
```bash
# 创建数据库
createdb -U postgres kg

# 执行schema
psql -U postgres -d kg -f database/schema.sql
```

### 连接信息
- **数据库名称**: kg
- **默认用户**: postgres
- **默认密码**: 12345678
- **默认主机**: localhost
- **默认端口**: 5432
- **连接字符串**: `postgres://postgres:12345678@localhost:5432/kg?sslmode=disable`

## 说明

1. 涉及到的新的数据库表仅为初步设计，后续可能进一步调整；
2. 此文档为初步计划，仍有许多后端接口仍在开发，待完善。
3. 此文件夹下的config.yaml是一些固定的内容（不会经常变更），包括套餐计划、指数计量等等，将其放在yaml中进行读取比从数据库中读取更快。