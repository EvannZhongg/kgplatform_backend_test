-- PostgreSQL数据库schema for kg平台
--如果存在则删除
DROP DATABASE IF EXISTS kg;

-- 创建数据库
CREATE DATABASE kg;
-- 连接到kg数据库
\c kg;

-- 创建枚举类型
CREATE TYPE user_plan_enum AS ENUM ('free', 'professional', 'team');
CREATE TYPE subscription_status_enum AS ENUM ('active', 'expired', 'cancelled', 'suspended');
CREATE TYPE team_status_enum AS ENUM ('active', 'suspended', 'deleted');
CREATE TYPE role_enum AS ENUM ('owner', 'admin', 'member');
CREATE TYPE status_enum AS ENUM ('active', 'pending', 'removed');
CREATE TYPE billing_type_enum AS ENUM ('subscription', 'overage', 'refund', 'month');
CREATE TYPE payment_status_enum AS ENUM ('pending', 'paid', 'failed', 'refunded', 'cancelled');
CREATE TYPE billing_status_enum AS ENUM ('unpaid', 'paid');

-- 创建用户表
create table users
(
    id         serial
        primary key,
    username   varchar(255),
    password   varchar(255),
    phone      varchar(255),
    email      varchar(255),
    created_at timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone default CURRENT_TIMESTAMP
);

comment on table users is '用户表';
comment on column users.username is '用户名';
comment on column users.phone is '手机号';
comment on column users.email is '邮箱';

alter table users
    owner to postgres;

-- 创建图谱表
create table graphs
(
    id         serial
        primary key,
    url        text,
    created_at timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone default CURRENT_TIMESTAMP
);

comment on table graphs is '图谱表';
comment on column graphs.url is 'graph生成的nodes/edges上传至云端后的url';

alter table graphs
    owner to postgres;

-- 创建项目表
create table projects
(
    id                 serial
        primary key,
    user_id            integer,
    project_name       varchar(255),
    project_progress   integer,
    graph_id           integer,
    created_at         timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at         timestamp with time zone default CURRENT_TIMESTAMP,
    schema_url         text,
    sample_text_url    text,
    sample_xlsx_url    text,
    snapshot_photo_url text,
    visibility         integer                  default 0 not null
);

comment on table projects is '项目表';
comment on column projects.project_name is '项目名称';
comment on column projects.project_progress is '项目进度(0-100)';
comment on column projects.schema_url is '主体结构url';
comment on column projects.sample_text_url is '示例原文的url';
comment on column projects.sample_xlsx_url is '示例抽取结果(三元组)的url';
comment on column projects.visibility is '可见性, 0-private, 1-public';
comment on column projects.snapshot_photo_url is '图谱照片url';

alter table projects
    owner to postgres;

create index idx_projects_user_id on projects (user_id);
create index idx_projects_graph_id on projects (graph_id);

-- 创建材料表
create table materials
(
    id         serial
        primary key,
    url        text,
    project_id integer,
    text_url   text,
    triple_url text,
    created_at timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone default CURRENT_TIMESTAMP
);

comment on table materials is '材料表';
comment on column materials.url is '材料URL';
comment on column materials.text_url is '文字化结果存储URL';
comment on column materials.triple_url is '三元组抽取结果存储URL';

alter table materials
    owner to postgres;

create index idx_materials_project_id on materials (project_id);

-- 创建流水线表
create table pipelines
(
    id         serial
        primary key,
    start_step varchar(255),
    project_id integer,
    created_at timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone default CURRENT_TIMESTAMP
);

comment on table pipelines is '流水线表';
comment on column pipelines.start_step is '起始步骤';

alter table pipelines
    owner to postgres;

create index idx_pipelines_project_id on pipelines (project_id);

-- 创建任务表
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

comment on table tasks is '任务表';
comment on column tasks.type is '任务类型';
comment on column tasks.status is '任务状态, pending-待处理, processing-处理中, completed-完成, failed-失败';
comment on column tasks.material_id_list is '任务需处理的材料';
comment on column tasks.error_message is '任务失败日志';

alter table tasks
    owner to postgres;

create index idx_tasks_pipeline_id on tasks (pipeline_id);

-- 创建视图收藏表
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

comment on table views is '视图收藏表';

alter table views
    owner to postgres;

create index idx_views_user_id on views (user_id);
create index idx_views_project_id on views (project_id);


-- 支持的模型（Models）
create table models
(
    id          serial primary key,
    provider    varchar(100) not null,                      -- 模型提供方，如 OpenAI/DeepSeek
    model_code  varchar(150) not null,                      -- 模型编码，如 gpt-4o、DeepSeek-V3-0324
    name        varchar(255) not null,                      -- 展示名称
    status      integer default 1 not null,                 -- 状态：1启用，0禁用
    description text,                                       -- 说明
    created_at  timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at  timestamp with time zone default CURRENT_TIMESTAMP
);

comment on table models is '支持的模型配置';
comment on column models.provider is '模型提供方';
comment on column models.model_code is '模型编码';
comment on column models.name is '展示名称';
comment on column models.status is '状态：1启用，0禁用';


-- 唯一约束：同一provider下model_code唯一
alter table models add constraint uq_models_provider_code unique (provider, model_code);

-- 索引
create index idx_models_status on models (status);


-- 支持的目标领域（Domains）
create table domains
(
    id           serial primary key,
    display_name varchar(255) not null,                     -- 展示名称，如 通用领域、建筑学
    status       integer default 1 not null,                -- 状态：1启用，0禁用
    description  text,                                      -- 说明
    created_at   timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at   timestamp with time zone default CURRENT_TIMESTAMP
);

comment on table domains is '支持的目标领域配置';
comment on column domains.display_name is '展示名称';
comment on column domains.status is '状态：1启用，0禁用';


-- 创建点赞表（点赞功能）
CREATE TABLE likes
(
    id         SERIAL PRIMARY KEY,
    user_id    INT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    project_id INT NOT NULL REFERENCES projects (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, project_id)
);
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
    team_name           VARCHAR(100)   NOT NULL,
    owner_id            BIGINT         NOT NULL,
    team_code           VARCHAR(50) UNIQUE,
    invite_code         VARCHAR(6) UNIQUE NOT NULL DEFAULT (
        substring(
                upper(
                        translate(md5(random()::text), 'abcdefghijklmnopqrstuvwxyz', '01234567890123456789012345')
                )
                from 1 for 6
        )
        ),

    total_words_quota   INT            NOT NULL DEFAULT 900000,
    total_storage_quota INT            NOT NULL DEFAULT 512000,
    total_cu_quota      INT            NOT NULL DEFAULT 2000,
    total_traffic_quota INT            NOT NULL DEFAULT 200,

    words_used          INT            NOT NULL DEFAULT 0,
    storage_used        INT            NOT NULL DEFAULT 0,
    cu_used             INT            NOT NULL DEFAULT 0,
    traffic_used        NUMERIC(10, 3) NOT NULL DEFAULT 0.000,

    member_count        INT            NOT NULL DEFAULT 3,

    status              team_status_enum        DEFAULT 'active',

    created_at          TIMESTAMPTZ             DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ             DEFAULT CURRENT_TIMESTAMP,

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
    user_id               BIGINT         NOT NULL,
    team_id               BIGINT,

    billing_period        VARCHAR(7)     NOT NULL,
    billing_date          DATE           NOT NULL,
    billing_type          billing_type_enum       DEFAULT 'subscription',

    base_subscription_fee NUMERIC(10, 2) NOT NULL DEFAULT 0.00,
    overage_fee           NUMERIC(10, 2) NOT NULL DEFAULT 0.00,
    subtotal              NUMERIC(10, 2) NOT NULL,
    discount_amount       NUMERIC(10, 2)          DEFAULT 0.00,
    total_amount          NUMERIC(10, 2) NOT NULL,

    status                billing_status_enum NOT NULL DEFAULT 'unpaid',

    remark                TEXT,
    created_at            TIMESTAMPTZ             DEFAULT CURRENT_TIMESTAMP,
    updated_at            TIMESTAMPTZ             DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_team FOREIGN KEY (team_id) REFERENCES teams (id) ON DELETE SET NULL
);
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




-- 统一创建触发器
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$ language 'plpgsql';

create trigger update_graphs_updated_at
    before update on graphs
    for each row execute procedure update_updated_at_column();

create trigger update_projects_updated_at
    before update on projects
    for each row execute procedure update_updated_at_column();

create trigger update_materials_updated_at
    before update on materials
    for each row execute procedure update_updated_at_column();

create trigger update_pipelines_updated_at
    before update on pipelines
    for each row execute procedure update_updated_at_column();

create trigger update_tasks_updated_at
    before update on tasks
    for each row execute procedure update_updated_at_column();

create trigger update_views_updated_at
    before update on views
    for each row execute procedure update_updated_at_column();

-- 新增触发器：models/support_domains
create trigger update_models_updated_at
    before update on models
    for each row execute procedure update_updated_at_column();

create trigger update_domains_updated_at
    before update on domains
    for each row execute procedure update_updated_at_column();

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