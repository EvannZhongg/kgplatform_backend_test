@echo off
REM PostgreSQL数据库初始化脚本 for kg平台 (Windows版本)

REM 设置数据库连接参数
set DB_NAME=kg
set DB_USER=postgres
set DB_PASSWORD=12345678
set DB_HOST=localhost
set DB_PORT=5432

echo 开始初始化kg数据库...

REM 检查PostgreSQL是否运行
pg_isready -h %DB_HOST% -p %DB_PORT% -U %DB_USER%
if %errorlevel% neq 0 (
    echo 错误: PostgreSQL服务未运行或无法连接
    echo 请确保PostgreSQL服务已启动，并且用户 %DB_USER% 有足够权限
    pause
    exit /b 1
)

REM 创建数据库（如果不存在）
echo 创建数据库 %DB_NAME%...
createdb -h %DB_HOST% -p %DB_PORT% -U %DB_USER% %DB_NAME% 2>nul
if %errorlevel% neq 0 (
    echo 数据库 %DB_NAME% 已存在或创建失败
)

REM 执行schema文件
echo 执行数据库schema...
psql -h %DB_HOST% -p %DB_PORT% -U %DB_USER% -d %DB_NAME% -f .\schema.sql

if %errorlevel% equ 0 (
    echo 数据库初始化成功！
    echo 数据库名称: %DB_NAME%
    echo 连接信息:
    echo   Host: %DB_HOST%
    echo   Port: %DB_PORT%
    echo   Database: %DB_NAME%
    echo   User: %DB_USER%
) else (
    echo 数据库初始化失败！
    pause
    exit /b 1
)

pause
