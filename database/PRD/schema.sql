-- PostgreSQL数据库schema for kg平台
--如果存在则删除
DROP
DATABASE IF EXISTS kg;

\c
kg

-- 创建数据库
CREATE
DATABASE kg;
-- 连接到kg数据库
-- \c kg;

-- 创建枚举类型
CREATE TYPE user_plan_enum AS ENUM ('free', 'professional', 'team');
CREATE TYPE subscription_status_enum AS ENUM ('active', 'expired', 'cancelled', 'suspended');
CREATE TYPE team_status_enum AS ENUM ('active', 'suspended', 'deleted');
CREATE TYPE role_enum AS ENUM ('owner', 'admin', 'member');
CREATE TYPE status_enum AS ENUM ('active', 'pending', 'removed');
CREATE TYPE billing_type_enum AS ENUM ('subscription', 'overage', 'refund', 'month');
CREATE TYPE payment_status_enum AS ENUM ('pending', 'paid', 'failed', 'refunded', 'cancelled');
CREATE TYPE billing_status_enum AS ENUM ('unpaid', 'paid');

ALTER TYPE billing_type_enum ADD VALUE 'project';

-- 创建用户表
CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(50)  NOT NULL UNIQUE,
    password   VARCHAR(100) NOT NULL,
    phone      VARCHAR(20),
    email      VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

comment
on table users is '用户表';
comment
on column users.username is '用户名';
comment
on column users.phone is '手机号';
comment
on column users.email is '邮箱';

-- 创建图谱表
CREATE TABLE graphs
(
    id         SERIAL PRIMARY KEY,
    url        TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

comment
on table graphs is '图谱表';
comment
on column graphs.url is 'graph生成的nodes/edges上传至云端后的url';

alter table graphs
    owner to postgres;

-- 创建项目表
CREATE TABLE projects
(
    id                 integer                            NOT NULL,
    user_id            integer,
    project_name       character varying(255),
    project_progress   integer,
    graph_id           integer,
    created_at         timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at         timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    schema_url         text,
    sample_text_url    text,
    sample_xlsx_url    text,
    visibility         integer                  DEFAULT 0 NOT NULL,
    snapshot_photo_url text,
    triple_type_url    text,
    triple_url         text,
    description        text,
    extract_config     jsonb
);


ALTER TABLE projects
    ADD COLUMN buy_price_cent DECIMAL(10, 2) DEFAULT 0.00 CHECK (buy_price_cent >= 0);

ALTER TABLE projects
    ADD COLUMN read_price_cent DECIMAL(10, 2) DEFAULT 0.00 CHECK (read_price_cent >= 0);

ALTER TABLE projects
    ADD COLUMN purchase_count INTEGER DEFAULT 0 CHECK (purchase_count >= 0);

ALTER TABLE projects
    ADD COLUMN view_count INTEGER DEFAULT 0;

comment
on table projects is '项目表';
comment
on column projects.project_name is '项目名称';
comment
on column projects.project_progress is '项目进度(0-100)';
comment
on column projects.schema_url is '主体结构url';
comment
on column projects.sample_text_url is '示例原文的url';
comment
on column projects.sample_xlsx_url is '示例抽取结果(三元组)的url';
comment
on column projects.visibility is '可见性, 0-private, 1-public';
comment
on column projects.snapshot_photo_url is '图谱照片url';

alter table projects
    owner to postgres;

create index idx_projects_user_id on projects (user_id);
create index idx_projects_graph_id on projects (graph_id);

-- 创建流水线表
CREATE TABLE pipelines
(
    id         SERIAL PRIMARY KEY,
    start_step VARCHAR(100),
    project_id INTEGER NOT NULL REFERENCES projects (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE pipelines
    ADD COLUMN likes_count INT DEFAULT 0;

ALTER TABLE pipelines
    ADD COLUMN favors_count INT DEFAULT 0;

ALTER TABLE pipelines
    ADD COLUMN views_count INT DEFAULT 0;

comment
on table pipelines is '流水线表';
comment
on column pipelines.start_step is '起始步骤';

alter table pipelines
    owner to postgres;

create index idx_pipelines_project_id on pipelines (project_id);

CREATE TABLE pipelines_likes
(
    id          SERIAL PRIMARY KEY,                                 -- 主键 ID
    user_id     INT NOT NULL,                                       -- 点赞用户 ID
    pipeline_id INT NOT NULL,                                       -- 被点赞的工作流 ID
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- 点赞时间
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- 更新时间
    UNIQUE (user_id, pipeline_id),                                  -- 保证每个用户只能点赞一次

    CONSTRAINT fk_user_pipeline_likes
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,

    CONSTRAINT fk_pipeline_pipeline_likes
        FOREIGN KEY (pipeline_id) REFERENCES pipelines (id) ON DELETE CASCADE
);
CREATE INDEX idx_pipeline_likes_pipeline_id ON pipelines_likes (pipeline_id);
CREATE INDEX idx_pipeline_likes_user_id ON pipelines_likes (user_id);

CREATE TABLE pipelines_favorites
(
    id          SERIAL PRIMARY KEY, -- 主键ID
    user_id     INT NOT NULL,       -- 收藏用户ID
    pipeline_id INT NOT NULL,       -- 被收藏的工作流ID
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, pipeline_id),  -- 用户不可重复收藏

    CONSTRAINT fk_user_pipeline_fav
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_pipeline_pipeline_fav
        FOREIGN KEY (pipeline_id) REFERENCES pipelines (id) ON DELETE CASCADE
);

CREATE INDEX idx_pipeline_fav_pipeline_id ON pipelines_favorites (pipeline_id);
CREATE INDEX idx_pipeline_fav_user_id ON pipelines_favorites (user_id);


-- 创建材料表
CREATE TABLE materials
(
    id         SERIAL PRIMARY KEY,
    url        TEXT    NOT NULL,
    project_id INTEGER NOT NULL REFERENCES projects (id) ON DELETE CASCADE,
    text_url   TEXT,
    triple_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

comment
on table materials is '材料表';
comment
on column materials.url is '材料URL';
comment
on column materials.text_url is '文字化结果存储URL';
comment
on column materials.triple_url is '三元组抽取结果存储URL';

alter table materials
    owner to postgres;

create index idx_materials_project_id on materials (project_id);

create table tasks
(
    id               serial
        primary key,
    type             varchar(255),
    pipeline_id      integer,
    created_at       timestamp with time zone,
    updated_at       timestamp with time zone,
    status           varchar(50),
    start_time       timestamp with time zone,
    finish_time      timestamp with time zone,
    material_id_list integer[],
    error_message    text,
    project_id       integer
);

comment
on table tasks is '任务表';
comment
on column tasks.type is '任务类型';
comment
on column tasks.status is '任务状态, pending-待处理, processing-处理中, completed-完成, failed-失败';
comment
on column tasks.material_id_list is '任务需处理的材料';
comment
on column tasks.error_message is '任务失败日志';

alter table tasks
    owner to postgres;

create index idx_tasks_pipeline_id on tasks (pipeline_id);

-- 创建视图表（收藏功能）
create table views
(
    table_catalog              varchar(255),
    id                         serial
        primary key,
    table_schema               varchar(255),
    user_id                    integer,
    project_id                 integer,
    table_name                 varchar(255),
    created_at                 timestamp with time zone,
    view_definition            text,
    updated_at                 timestamp with time zone,
    check_option               varchar(255),
    is_updatable               varchar(255),
    is_insertable_into         varchar(255),
    is_trigger_updatable       varchar(255),
    is_trigger_deletable       varchar(255),
    is_trigger_insertable_into varchar(255)
);

comment
on table views is '视图收藏表';

alter table views
    owner to postgres;

create index idx_views_user_id on views (user_id);
create index idx_views_project_id on views (project_id);

-- 创建点赞表（点赞功能）
CREATE TABLE projects_likes
(
    id         SERIAL PRIMARY KEY,
    user_id    INT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    project_id INT NOT NULL REFERENCES projects (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, project_id)
);

create table models
(
    id          serial primary key,
    provider    varchar(100)                       not null, -- 模型提供方，如 OpenAI/DeepSeek
    model_code  varchar(150)                       not null, -- 模型编码，如 gpt-4o、DeepSeek-V3-0324
    name        varchar(255)                       not null, -- 展示名称
    status      integer                  default 1 not null, -- 状态：1启用，0禁用
    description text,                                        -- 说明
    created_at  timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at  timestamp with time zone default CURRENT_TIMESTAMP
);

comment
on table models is '支持的模型配置';
comment
on column models.provider is '模型提供方';
comment
on column models.model_code is '模型编码';
comment
on column models.name is '展示名称';
comment
on column models.status is '状态：1启用，0禁用';


-- 唯一约束：同一provider下model_code唯一
alter table models
    add constraint uq_models_provider_code unique (provider, model_code);

-- 索引
create index idx_models_status on models (status);

create table domains
(
    id           serial primary key,
    display_name varchar(255)                       not null, -- 展示名称，如 通用领域、建筑学
    status       integer                  default 1 not null, -- 状态：1启用，0禁用
    description  text,                                        -- 说明
    created_at   timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at   timestamp with time zone default CURRENT_TIMESTAMP
);

comment
on table domains is '支持的目标领域配置';
comment
on column domains.display_name is '展示名称';
comment
on column domains.status is '状态：1启用，0禁用';

-- 创建项目评论表
CREATE TABLE comments
(
    id         SERIAL PRIMARY KEY,
    project_id INT  NOT NULL REFERENCES projects (id) ON DELETE CASCADE,
    user_id    INT  NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    parent_id  INT REFERENCES comments (id) ON DELETE CASCADE,
    content    TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 创建评论点赞表
CREATE TABLE comment_likes
(
    id         SERIAL PRIMARY KEY,
    user_id    INT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    comment_id INT NOT NULL REFERENCES comments (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, comment_id)
);

-- 创建团队表
CREATE TABLE teams
(
    id                  BIGSERIAL PRIMARY KEY,
    team_name           VARCHAR(100)      NOT NULL,
    owner_id            BIGINT            NOT NULL,
    team_code           VARCHAR(50) UNIQUE,
    invite_code         VARCHAR(6) UNIQUE NOT NULL DEFAULT (
        substring(
                upper(
                        translate(md5(random()::text), 'abcdefghijklmnopqrstuvwxyz', '01234567890123456789012345')
                )
                from 1 for 6
        )
        ),

    total_words_quota   INT               NOT NULL DEFAULT 900000,
    total_storage_quota INT               NOT NULL DEFAULT 512000,
    total_cu_quota      INT               NOT NULL DEFAULT 2000,
    total_traffic_quota INT               NOT NULL DEFAULT 200,

    words_used          INT               NOT NULL DEFAULT 0,
    storage_used        INT               NOT NULL DEFAULT 0,
    cu_used             INT               NOT NULL DEFAULT 0,
    traffic_used        NUMERIC(10, 3)    NOT NULL DEFAULT 0.000,

    member_count        INT               NOT NULL DEFAULT 3,

    status              team_status_enum           DEFAULT 'active',

    created_at          TIMESTAMPTZ                DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ                DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_owner FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE
);

-- 创建用户订阅表
CREATE TABLE user_subscriptions
(
    id                       BIGSERIAL PRIMARY KEY,
    user_id                  BIGINT                   NOT NULL UNIQUE,
    team_id                  BIGINT                            DEFAULT NULL,

    user_plan                user_plan_enum           NOT NULL DEFAULT 'free',
    subscription_status      subscription_status_enum NOT NULL DEFAULT 'active',

    words_used               INT                      NOT NULL DEFAULT 0,
    storage_used             INT                      NOT NULL DEFAULT 0,
    cu_used                  INT                      NOT NULL DEFAULT 0,
    traffic_used             NUMERIC(10, 5)           NOT NULL DEFAULT 0.00000,

    quota_reset_date         DATE                     NOT NULL,

    words_warning_80_sent    BOOLEAN                           DEFAULT FALSE,
    words_warning_100_sent   BOOLEAN                           DEFAULT FALSE,
    storage_warning_80_sent  BOOLEAN                           DEFAULT FALSE,
    storage_warning_100_sent BOOLEAN                           DEFAULT FALSE,
    cu_warning_80_sent       BOOLEAN                           DEFAULT FALSE,
    cu_warning_100_sent      BOOLEAN                           DEFAULT FALSE,
    traffic_warning_80_sent  BOOLEAN                           DEFAULT FALSE,
    traffic_warning_100_sent BOOLEAN                           DEFAULT FALSE,

    overage_words_fee        NUMERIC(10, 2)           NOT NULL DEFAULT 0.00,
    overage_storage_fee      NUMERIC(10, 2)           NOT NULL DEFAULT 0.00,
    overage_traffic_fee      NUMERIC(10, 2)           NOT NULL DEFAULT 0.00,
    overage_cu_fee           NUMERIC(10, 2)           NOT NULL DEFAULT 0.00,
    total_overage_fee        NUMERIC(10, 2)           NOT NULL DEFAULT 0.00,

    selected_ai_model        VARCHAR(50)                       DEFAULT 'gpt-4o',

    created_at               TIMESTAMPTZ                       DEFAULT CURRENT_TIMESTAMP,
    updated_at               TIMESTAMPTZ                       DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_team FOREIGN KEY (team_id) REFERENCES teams (id) ON DELETE SET NULL
);

-- 创建团队成员表
CREATE TABLE team_members
(
    id                    BIGSERIAL PRIMARY KEY,
    team_id               BIGINT         NOT NULL,
    user_id               BIGINT         NOT NULL,

    role                  role_enum      NOT NULL DEFAULT 'member',

    allocated_words_quota INT,

    personal_words_used   INT            NOT NULL DEFAULT 0,
    personal_storage_used INT            NOT NULL DEFAULT 0,
    personal_cu_used      INT            NOT NULL DEFAULT 0,
    personal_traffic_used NUMERIC(10, 3) NOT NULL DEFAULT 0.000,

    status                status_enum             DEFAULT 'active',
    invite_code           VARCHAR(100),
    invited_by            BIGINT,

    joined_at             TIMESTAMPTZ             DEFAULT CURRENT_TIMESTAMP,
    removed_at            TIMESTAMPTZ,

    CONSTRAINT fk_team FOREIGN KEY (team_id) REFERENCES teams (id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_invited_by FOREIGN KEY (invited_by) REFERENCES users (id) ON DELETE SET NULL,

    CONSTRAINT uk_team_user UNIQUE (team_id, user_id)
);
-- 创建账单记录表
CREATE TABLE billing_records
(
    id                    BIGSERIAL PRIMARY KEY,
    user_id               BIGINT              NOT NULL,
    team_id               BIGINT,

    billing_period        VARCHAR(7)          NOT NULL,
    billing_date          DATE                NOT NULL,
    billing_type          billing_type_enum            DEFAULT 'subscription',

    base_subscription_fee NUMERIC(10, 2)      NOT NULL DEFAULT 0.00,
    overage_fee           NUMERIC(10, 2)      NOT NULL DEFAULT 0.00,
    subtotal              NUMERIC(10, 2)      NOT NULL,
    discount_amount       NUMERIC(10, 2)               DEFAULT 0.00,
    total_amount          NUMERIC(10, 2)      NOT NULL,

    status                billing_status_enum NOT NULL DEFAULT 'unpaid',

    remark                TEXT,
    created_at            TIMESTAMPTZ                  DEFAULT CURRENT_TIMESTAMP,
    updated_at            TIMESTAMPTZ                  DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_team FOREIGN KEY (team_id) REFERENCES teams (id) ON DELETE SET NULL
);

ALTER TABLE billing_records
    ADD COLUMN alipay_payment_url TEXT;

ALTER TABLE billing_records
    ADD COLUMN wechat_payment_url TEXT;

-- 创建账单支付记录表
CREATE TABLE billing_payments
(
    id                     BIGSERIAL PRIMARY KEY,
    billing_id             BIGINT         NOT NULL,

    payment_status         payment_status_enum DEFAULT 'pending',
    payment_method         VARCHAR(50)    NOT NULL,
    payment_amount         NUMERIC(10, 2) NOT NULL,
    payment_transaction_id VARCHAR(100),
    payment_channel        VARCHAR(50),

    paid_at                TIMESTAMPTZ,
    failure_reason         VARCHAR(500),
    refunded_at            TIMESTAMPTZ,

    created_at             TIMESTAMPTZ         DEFAULT CURRENT_TIMESTAMP,
    updated_at             TIMESTAMPTZ         DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_billing FOREIGN KEY (billing_id) REFERENCES billing_records (id) ON DELETE CASCADE
);
-- 创建超额记录表
CREATE TABLE billing_overages
(
    id                     BIGSERIAL PRIMARY KEY,
    billing_id             BIGINT NOT NULL,

    words_overage_amount   NUMERIC(10, 3) DEFAULT 0.000,
    words_overage_fee      NUMERIC(10, 2) DEFAULT 0.00,

    storage_overage_amount NUMERIC(10, 3) DEFAULT 0.000,
    storage_overage_fee    NUMERIC(10, 2) DEFAULT 0.00,

    traffic_overage_amount NUMERIC(10, 3) DEFAULT 0.000,
    traffic_overage_fee    NUMERIC(10, 2) DEFAULT 0.00,

    cu_overage_amount      NUMERIC(10, 3) DEFAULT 0.000,
    cu_overage_fee         NUMERIC(10, 2) DEFAULT 0.00,

    total_overage_fee      NUMERIC(10, 2) DEFAULT 0.00,

    created_at             TIMESTAMPTZ    DEFAULT CURRENT_TIMESTAMP,
    updated_at             TIMESTAMPTZ    DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_billing FOREIGN KEY (billing_id) REFERENCES billing_records (id) ON DELETE CASCADE,

    CONSTRAINT uk_billing UNIQUE (billing_id)
);

-- 创建发票记录表
CREATE TABLE billing_invoices
(
    id                BIGSERIAL PRIMARY KEY,
    billing_id        BIGINT NOT NULL UNIQUE,

    invoice_required  BOOLEAN     DEFAULT FALSE,
    invoice_title     VARCHAR(200),
    invoice_tax_id    VARCHAR(50),
    invoice_url       VARCHAR(500),
    invoice_issued_at TIMESTAMPTZ,

    created_at        TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_billing FOREIGN KEY (billing_id) REFERENCES billing_records (id) ON DELETE CASCADE
);

-- 创建流量记录日志表
CREATE TABLE traffic_logs
(
    id           BIGSERIAL PRIMARY KEY,
    user_id      BIGINT         NOT NULL,
    team_id      BIGINT      DEFAULT NULL,

    traffic_type VARCHAR(50)    NOT NULL,

    data_size    BIGINT         NOT NULL,
    traffic_kb   NUMERIC(10, 3) NOT NULL,

    endpoint     VARCHAR(200),
    ip_address   VARCHAR(50),

    created_at   TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_traffic_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_traffic_team FOREIGN KEY (team_id) REFERENCES teams (id) ON DELETE SET NULL
);

CREATE TABLE user_project_purchases
(
    id             BIGSERIAL PRIMARY KEY,
    user_id        BIGINT      NOT NULL,
    project_id     BIGINT      NOT NULL,
    billing_id     BIGINT      NOT NULL,
    payment_id     BIGINT,
    purchase_price NUMERIC(10, 2) NOT NULL,
    status         VARCHAR(50) DEFAULT 'completed',
    created_at     TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    -- 外键约束
    CONSTRAINT fk_user_purchase_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_user_purchase_project FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE,
    CONSTRAINT fk_user_purchase_billing FOREIGN KEY (billing_id) REFERENCES billing_records (id) ON DELETE CASCADE,
    CONSTRAINT fk_user_purchase_payment FOREIGN KEY (payment_id) REFERENCES billing_payments (id) ON DELETE SET NULL,

    -- 确保一个用户购买一个项目的记录唯一性
    CONSTRAINT uk_user_project UNIQUE (user_id, project_id)
);

-- 创建索引以提高查询性能
CREATE INDEX idx_projects_user_id ON projects (user_id);
CREATE INDEX idx_projects_graph_id ON projects (graph_id);
CREATE INDEX idx_materials_project_id ON materials (project_id);
CREATE INDEX idx_tasks_pipeline_id ON tasks (pipeline_id);
CREATE INDEX idx_pipelines_project_id ON pipelines (project_id);
CREATE INDEX idx_views_user_id ON views (user_id);
CREATE INDEX idx_views_project_id ON views (project_id);
CREATE INDEX idx_plan_status ON user_subscriptions (user_plan, subscription_status);
CREATE INDEX idx_user_reset_date ON user_subscriptions (user_id, quota_reset_date);
CREATE INDEX idx_team_id ON user_subscriptions (team_id);
CREATE INDEX idx_owner ON teams (owner_id);
CREATE INDEX idx_team_code ON teams (team_code);
CREATE INDEX idx_status ON teams (status);
CREATE INDEX idx_user_status ON team_members (user_id, status);
CREATE INDEX idx_role ON team_members (role);
CREATE INDEX idx_invite_code ON team_members (invite_code);
CREATE INDEX idx_user_period ON billing_records (user_id, billing_period);
CREATE INDEX idx_billing_date ON billing_records (billing_date);
CREATE INDEX idx_billing_payments ON billing_payments (billing_id);
CREATE INDEX idx_status_time ON billing_payments (payment_status, created_at);
CREATE INDEX idx_transaction_id ON billing_payments (payment_transaction_id);
CREATE INDEX idx_billing_overages ON billing_overages (billing_id);
CREATE INDEX idx_invoice_issued ON billing_invoices (invoice_issued_at);
CREATE INDEX idx_invoice_title ON billing_invoices (invoice_title);
CREATE INDEX idx_user_time ON traffic_logs (user_id, created_at);
CREATE INDEX idx_traffic_type ON traffic_logs (traffic_type);
CREATE INDEX idx_team_time ON traffic_logs (team_id, created_at);
CREATE INDEX idx_team_invite_code ON teams (invite_code);

-- 创建更新时间触发器函数
CREATE
OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at
= CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$
language 'plpgsql';

-- 为 billing_overages 表创建触发器
CREATE
OR REPLACE FUNCTION update_billing_totals()
RETURNS TRIGGER AS $$
DECLARE
total_overage DECIMAL(10,2);
    billing_record
RECORD;
BEGIN
    DECLARE
affected_billing_id BIGINT;
BEGIN
        IF
TG_OP = 'DELETE' THEN
            affected_billing_id := OLD.billing_id;
ELSE
            affected_billing_id := NEW.billing_id;
END IF;

SELECT COALESCE(SUM(
                        COALESCE(words_overage_fee, 0) +
                        COALESCE(storage_overage_fee, 0) +
                        COALESCE(traffic_overage_fee, 0) +
                        COALESCE(cu_overage_fee, 0)
                ), 0)
INTO total_overage
FROM billing_overages
WHERE billing_id = affected_billing_id;

SELECT *
INTO billing_record
FROM billing_records
WHERE id = affected_billing_id;

IF
FOUND THEN
UPDATE billing_records
SET overage_fee  = total_overage,
    subtotal     = billing_record.base_subscription_fee + total_overage,
    total_amount = billing_record.base_subscription_fee + total_overage - COALESCE(billing_record.discount_amount, 0),
    updated_at   = CURRENT_TIMESTAMP
WHERE id = affected_billing_id;
END IF;

END;

    IF
TG_OP = 'DELETE' THEN
        RETURN OLD;
ELSE
        RETURN NEW;
END IF;
END;
$$
LANGUAGE plpgsql;

-- 为超额明细表添加计算字段的触发器
CREATE
OR REPLACE FUNCTION calculate_overage_totals()
RETURNS TRIGGER AS $$
BEGIN

    NEW.total_overage_fee
:=
        COALESCE(NEW.words_overage_fee, 0) +
        COALESCE(NEW.storage_overage_fee, 0) +
        COALESCE(NEW.traffic_overage_fee, 0) +
        COALESCE(NEW.cu_overage_fee, 0);

RETURN NEW;
END;
$$
LANGUAGE plpgsql;

-- 为需要updated_at字段的表创建触发器
CREATE TRIGGER update_graphs_updated_at
    BEFORE UPDATE
    ON graphs
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_projects_updated_at
    BEFORE UPDATE
    ON projects
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_materials_updated_at
    BEFORE UPDATE
    ON materials
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_tasks_updated_at
    BEFORE UPDATE
    ON tasks
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_pipelines_updated_at
    BEFORE UPDATE
    ON pipelines
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_views_updated_at
    BEFORE UPDATE
    ON views
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- 为超额明细表创建触发器
CREATE TRIGGER trigger_update_billing_totals_insert
    AFTER INSERT
    ON billing_overages
    FOR EACH ROW
    EXECUTE FUNCTION update_billing_totals();

CREATE TRIGGER trigger_update_billing_totals_update
    AFTER UPDATE
    ON billing_overages
    FOR EACH ROW
    EXECUTE FUNCTION update_billing_totals();

CREATE TRIGGER trigger_update_billing_totals_delete
    AFTER DELETE
    ON billing_overages
    FOR EACH ROW
    EXECUTE FUNCTION update_billing_totals();

-- 为插入和更新创建触发器
CREATE TRIGGER trigger_calculate_overage_totals_insert
    BEFORE INSERT
    ON billing_overages
    FOR EACH ROW
    EXECUTE FUNCTION calculate_overage_totals();

CREATE TRIGGER trigger_calculate_overage_totals_update
    BEFORE UPDATE
    ON billing_overages
    FOR EACH ROW
    EXECUTE FUNCTION calculate_overage_totals();

CREATE TRIGGER update_user_project_purchases_updated_at
    BEFORE UPDATE ON user_project_purchases
    FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

-- 添加注释
COMMENT
ON TABLE users IS '用户表';
COMMENT
ON TABLE projects IS '项目表';
COMMENT
ON TABLE graphs IS '图谱表';
COMMENT
ON TABLE materials IS '材料表';
COMMENT
ON TABLE tasks IS '任务表';
COMMENT
ON TABLE pipelines IS '流水线表';
COMMENT
ON TABLE views IS '视图收藏表';
COMMENT
ON TABLE likes IS '点赞表';
COMMENT
ON TABLE comments IS '评论表';
COMMENT
ON TABLE comment_likes IS '评论点赞表';
COMMENT
ON TABLE user_subscriptions IS '用户订阅配额表';
COMMENT
ON TABLE teams IS '团队表';
COMMENT
ON TABLE team_members IS '团队成员表';
COMMENT
ON TABLE billing_payments IS '账单支付记录表';
COMMENT
ON TABLE billing_overages IS '账单超额费用明细表';
COMMENT
ON TABLE billing_invoices IS '账单发票信息表';
COMMENT
ON TABLE traffic_logs IS '流量使用日志表';

COMMENT
ON COLUMN users.username IS '用户名';
COMMENT
ON COLUMN users.phone IS '手机号';
COMMENT
ON COLUMN users.email IS '邮箱';
COMMENT
ON COLUMN projects.project_name IS '项目名称';
COMMENT
ON COLUMN projects.project_progress IS '项目进度(0-100)';
COMMENT
ON COLUMN graphs.url IS 'Neo4j数据库URL';
COMMENT
ON COLUMN materials.url IS '材料URL';
COMMENT
ON COLUMN materials.text_url IS '文字化结果存储URL';
COMMENT
ON COLUMN materials.triple_url IS '三元组抽取结果存储URL';
COMMENT
ON COLUMN tasks.type IS '任务类型';
COMMENT
ON COLUMN pipelines.start_step IS '起始步骤';
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
COMMENT
ON COLUMN traffic_logs.id IS '日志ID';
COMMENT
ON COLUMN traffic_logs.user_id IS '用户ID';
COMMENT
ON COLUMN traffic_logs.team_id IS '团队ID(如果是团队使用)';
COMMENT
ON COLUMN traffic_logs.traffic_type IS '流量类型: graph_query/api_call/file_transfer等';
COMMENT
ON COLUMN traffic_logs.data_size IS '数据大小(字节)';
COMMENT
ON COLUMN traffic_logs.traffic_gb IS '流量(GB)';
COMMENT
ON COLUMN traffic_logs.project_id IS '关联项目ID';
COMMENT
ON COLUMN traffic_logs.graph_id IS '关联图谱ID';
COMMENT
ON COLUMN traffic_logs.endpoint IS '请求端点';
COMMENT
ON COLUMN traffic_logs.ip_address IS 'IP地址';
COMMENT
ON COLUMN traffic_logs.created_at IS '创建时间';