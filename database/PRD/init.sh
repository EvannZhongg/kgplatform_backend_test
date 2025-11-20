#!/bin/bash
# PostgreSQL数据库初始化脚本 for kg平台

# 设置数据库连接参数
DB_NAME="kg"
DB_USER="postgres"
DB_PASSWORD="12345678"
DB_HOST="localhost"
DB_PORT="5432"

echo "开始初始化kg数据库..."

# 检查PostgreSQL是否运行
if ! pg_isready -h $DB_HOST -p $DB_PORT -U $DB_USER; then
    echo "错误: PostgreSQL服务未运行或无法连接"
    echo "请确保PostgreSQL服务已启动，并且用户 $DB_USER 有足够权限"
    exit 1
fi

# 设置密码环境变量
export PGPASSWORD=$DB_PASSWORD

# 创建数据库（如果不存在）
echo "创建数据库 $DB_NAME..."
createdb -h $DB_HOST -p $DB_PORT -U $DB_USER $DB_NAME 2>/dev/null || echo "数据库 $DB_NAME 已存在"

# 执行schema文件
echo "执行数据库schema..."
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f ./schema.sql

if [ $? -eq 0 ]; then
    echo "数据库初始化成功！"
    echo "数据库名称: $DB_NAME"
    echo "连接信息:"
    echo "  Host: $DB_HOST"
    echo "  Port: $DB_PORT"
    echo "  Database: $DB_NAME"
    echo "  User: $DB_USER"
else
    echo "数据库初始化失败！"
    exit 1
fi
