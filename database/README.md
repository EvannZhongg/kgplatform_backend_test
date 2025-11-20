# KG平台数据库设计

## 概述
这是KG（知识图谱）平台的PostgreSQL数据库设计，包含用户管理、项目管理、图谱管理、材料处理、任务调度等核心功能。

## 数据库结构

### 表关系图
```
users (用户表)
  ├── projects (项目表) [user_id]
  │   ├── materials (材料表) [project_id]
  │   ├── pipelines (流水线表) [project_id]
  │   │   └── tasks (任务表) [pipeline_id]
  │   └── views (收藏表) [user_id, project_id]
  └── graphs (图谱表) [独立表，被projects引用]
```

### 表结构说明

#### 1. users (用户表)
- `id`: 主键，自增
- `username`: 用户名，唯一约束
- `phone`: 手机号
- `email`: 邮箱
- `created_at`: 创建时间

#### 2. projects (项目表)
- `id`: 主键，自增
- `user_id`: 外键，关联users表
- `project_name`: 项目名称
- `project_progress`: 项目进度 (0-100)
- `graph_id`: 外键，关联graphs表
- `created_at`: 创建时间
- `updated_at`: 更新时间

#### 3. graphs (图谱表)
- `id`: 主键，自增
- `url`: Neo4j数据库URL
- `created_at`: 创建时间
- `updated_at`: 更新时间

#### 4. materials (材料表)
- `id`: 主键，自增
- `url`: 材料URL
- `project_id`: 外键，关联projects表
- `text_url`: 文字化结果存储URL
- `triple_url`: 三元组抽取结果存储URL
- `created_at`: 创建时间
- `updated_at`: 更新时间

#### 5. pipelines (流水线表)
- `id`: 主键，自增
- `start_step`: 起始步骤
- `project_id`: 外键，关联projects表
- `created_at`: 创建时间
- `updated_at`: 更新时间

#### 6. tasks (任务表)
- `id`: 主键，自增
- `type`: 任务类型
- `pipeline_id`: 外键，关联pipelines表
- `created_at`: 创建时间
- `updated_at`: 更新时间

#### 7. views (收藏表)
- `id`: 主键，自增
- `user_id`: 外键，关联users表
- `project_id`: 外键，关联projects表
- `created_at`: 创建时间
- `updated_at`: 更新时间
- 唯一约束: (user_id, project_id)

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

## 特性

### 1. 自动更新时间戳
所有包含 `updated_at` 字段的表都配置了触发器，当记录更新时自动更新 `updated_at` 字段。

### 2. 外键约束
- 级联删除：删除用户时自动删除相关项目
- 设置NULL：删除图谱时，相关项目的graph_id设置为NULL

### 3. 性能优化
- 为所有外键字段创建了索引
- 为常用查询字段创建了索引

### 4. 数据完整性
- 项目进度字段有CHECK约束 (0-100)
- 收藏表有唯一约束防止重复收藏
- 用户名有唯一约束

## 示例查询

### 查询用户的所有项目
```sql
SELECT p.*, g.url as graph_url 
FROM projects p 
LEFT JOIN graphs g ON p.graph_id = g.id 
WHERE p.user_id = 1;
```

### 查询项目的所有材料
```sql
SELECT m.* 
FROM materials m 
WHERE m.project_id = 1;
```

### 查询用户收藏的项目
```sql
SELECT p.*, v.created_at as favorited_at
FROM projects p
JOIN views v ON p.id = v.project_id
WHERE v.user_id = 1;
```

## 维护

### 备份数据库
```bash
pg_dump -U postgres kg > kg_backup.sql
```

### 恢复数据库
```bash
psql -U postgres -d kg < kg_backup.sql
```
