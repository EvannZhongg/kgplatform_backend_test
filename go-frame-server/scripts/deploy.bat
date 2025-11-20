@echo off
chcp 65001 >nul
setlocal enabledelayedexpansion

REM KG平台一键部署脚本 (Windows版本)

:menu
cls
echo.
echo === KG平台部署管理工具 ===
echo 1. 一键部署 (构建+启动)
echo 2. 仅启动服务
echo 3. 启动管理工具
echo 4. 检查服务状态
echo 5. 查看服务信息
echo 6. 停止服务
echo 7. 查看日志
echo 8. 清理数据
echo 9. 退出
echo.

set /p choice="请选择操作 (1-9): "

if "%choice%"=="1" goto deploy_all
if "%choice%"=="2" goto start_services
if "%choice%"=="3" goto start_tools
if "%choice%"=="4" goto check_health
if "%choice%"=="5" goto show_info
if "%choice%"=="6" goto stop_services
if "%choice%"=="7" goto view_logs
if "%choice%"=="8" goto clean_data
if "%choice%"=="9" goto exit
echo [ERROR] 无效选择，请重新输入
pause
goto menu

:deploy_all
echo.
echo === 开始一键部署 ===

echo [INFO] 检查Docker环境...
docker --version >nul 2>&1
if errorlevel 1 (
    echo [ERROR] Docker未安装，请先安装Docker Desktop
    pause
    goto menu
)

docker-compose --version >nul 2>&1
if errorlevel 1 (
    echo [ERROR] Docker Compose未安装，请先安装Docker Compose
    pause
    goto menu
)
echo [INFO] Docker环境检查通过

echo [INFO] 检查配置文件...
if not exist "manifest\config\config.prod.yaml" (
    echo [ERROR] 生产环境配置文件不存在: manifest\config\config.prod.yaml
    pause
    goto menu
)

if not exist "docker-compose.prod.yml" (
    echo [ERROR] 生产环境Docker Compose文件不存在: docker-compose.prod.yml
    pause
    goto menu
)
echo [INFO] 配置文件检查通过

echo [INFO] 构建应用镜像...
docker-compose -f docker-compose.prod.yml build app
if errorlevel 1 (
    echo [ERROR] 应用镜像构建失败
    pause
    goto menu
)
echo [INFO] 应用镜像构建成功

echo [INFO] 启动所有服务...
docker-compose -f docker-compose.prod.yml up -d
timeout /t 10 /nobreak >nul

echo [INFO] 检查服务状态...
docker-compose -f docker-compose.prod.yml ps

echo [INFO] 部署完成！
call :show_info
pause
goto menu

:start_services
echo.
echo === 启动服务 ===
echo [INFO] 启动所有服务...
docker-compose -f docker-compose.prod.yml up -d
echo [INFO] 服务启动完成！
pause
goto menu

:start_tools
echo.
echo === 启动管理工具 ===
echo [INFO] 启动Redis管理界面和PostgreSQL管理界面...
docker-compose -f docker-compose.prod.yml --profile tools up -d
echo [INFO] 管理工具启动完成！
echo [INFO] Redis管理界面: http://localhost:8081
echo [INFO] PostgreSQL管理界面: http://localhost:8082
echo [INFO]   - 邮箱: admin@kgplatform.com
echo [INFO]   - 密码: admin123
pause
goto menu

:check_health
echo.
echo === 检查服务健康状态 ===
echo [INFO] 检查服务状态...
docker-compose -f docker-compose.prod.yml ps
pause
goto menu

:show_info
echo.
echo === 服务信息 ===
echo 应用服务:
echo   - 地址: http://localhost:8000
echo   - API文档: http://localhost:8000/swagger
echo   - OpenAPI: http://localhost:8000/api.json
echo.
echo 数据库服务:
echo   - PostgreSQL: localhost:5432
echo   - Redis: localhost:6379
echo.
echo 管理界面:
echo   - Redis管理: http://localhost:8081
echo   - PostgreSQL管理: http://localhost:8082
echo.
echo 测试SMS功能:
echo curl -X POST http://localhost:8000/v1/sms/send ^
echo   -H "Content-Type: application/json" ^
echo   -d "{\"phone\":\"13800138000\"}"
goto :eof

:stop_services
echo.
echo === 停止服务 ===
echo [INFO] 停止所有服务...
docker-compose -f docker-compose.prod.yml down
echo [INFO] 服务已停止
pause
goto menu

:view_logs
echo.
echo === 查看服务日志 ===
echo 选择要查看的服务日志:
echo 1. 应用服务
echo 2. PostgreSQL
echo 3. Redis
echo 4. 所有服务
echo.
set /p log_choice="请选择 (1-4): "

if "%log_choice%"=="1" (
    docker-compose -f docker-compose.prod.yml logs app
) else if "%log_choice%"=="2" (
    docker-compose -f docker-compose.prod.yml logs postgres
) else if "%log_choice%"=="3" (
    docker-compose -f docker-compose.prod.yml logs redis
) else if "%log_choice%"=="4" (
    docker-compose -f docker-compose.prod.yml logs
) else (
    echo [ERROR] 无效选择
)
pause
goto menu

:clean_data
echo.
echo === 清理数据 ===
echo [WARNING] 这将删除所有数据，是否继续？(y/N)
set /p response=
if /i "%response%"=="y" (
    echo [INFO] 停止服务并清理数据...
    docker-compose -f docker-compose.prod.yml down -v
    echo [INFO] 清理Docker镜像...
    docker image prune -f
    echo [INFO] 数据清理完成
) else (
    echo [INFO] 操作已取消
)
pause
goto menu

:exit
echo [INFO] 退出
exit /b 0
