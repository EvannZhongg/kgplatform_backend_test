#!/bin/bash

# KG平台一键部署脚本

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 打印带颜色的消息
print_message() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

print_header() {
    echo -e "${BLUE}=== $1 ===${NC}"
}

# 检查Docker是否安装
check_docker() {
    print_header "检查Docker环境"
    
    if ! command -v docker &> /dev/null; then
        print_error "Docker未安装，请先安装Docker"
        exit 1
    fi
    
    if ! command -v docker-compose &> /dev/null; then
        print_error "Docker Compose未安装，请先安装Docker Compose"
        exit 1
    fi
    
    print_message "Docker环境检查通过"
}

# 检查配置文件
check_config() {
    print_header "检查配置文件"
    
    if [ ! -f "manifest/config/config.prod.yaml" ]; then
        print_error "生产环境配置文件不存在: manifest/config/config.prod.yaml"
        exit 1
    fi
    
    if [ ! -f "docker-compose.prod.yml" ]; then
        print_error "生产环境Docker Compose文件不存在: docker-compose.prod.yml"
        exit 1
    fi
    
    print_message "配置文件检查通过"
}

# 构建应用镜像
build_app() {
    print_header "构建应用镜像"
    
    print_message "开始构建应用镜像..."
    docker-compose -f docker-compose.prod.yml build app
    
    if [ $? -eq 0 ]; then
        print_message "应用镜像构建成功"
    else
        print_error "应用镜像构建失败"
        exit 1
    fi
}

# 启动服务
start_services() {
    print_header "启动服务"
    
    print_message "启动所有服务..."
    docker-compose -f docker-compose.prod.yml up -d
    
    # 等待服务启动
    print_message "等待服务启动..."
    sleep 10
    
    # 检查服务状态
    print_message "检查服务状态..."
    docker-compose -f docker-compose.prod.yml ps
    
    print_message "服务启动完成！"
}

# 启动管理工具
start_tools() {
    print_header "启动管理工具"
    
    print_message "启动Redis管理界面和PostgreSQL管理界面..."
    docker-compose -f docker-compose.prod.yml --profile tools up -d
    
    print_message "管理工具启动完成！"
    print_message "Redis管理界面: http://localhost:8081"
    print_message "PostgreSQL管理界面: http://localhost:8082"
    print_message "  - 邮箱: admin@kgplatform.com"
    print_message "  - 密码: admin123"
}

# 检查服务健康状态
check_health() {
    print_header "检查服务健康状态"
    
    # 检查应用健康状态
    print_message "检查应用服务..."
    if docker-compose -f docker-compose.prod.yml ps app | grep -q "healthy"; then
        print_message "✅ 应用服务运行正常"
    else
        print_warning "⚠️ 应用服务可能未完全启动，请稍等片刻"
    fi
    
    # 检查数据库
    print_message "检查PostgreSQL..."
    if docker-compose -f docker-compose.prod.yml ps postgres | grep -q "healthy"; then
        print_message "✅ PostgreSQL运行正常"
    else
        print_warning "⚠️ PostgreSQL可能未完全启动"
    fi
    
    # 检查Redis
    print_message "检查Redis..."
    if docker-compose -f docker-compose.prod.yml ps redis | grep -q "healthy"; then
        print_message "✅ Redis运行正常"
    else
        print_warning "⚠️ Redis可能未完全启动"
    fi
}

# 显示服务信息
show_info() {
    print_header "服务信息"
    
    echo -e "${GREEN}应用服务:${NC}"
    echo "  - 地址: http://localhost:8000"
    echo "  - API文档: http://localhost:8000/swagger"
    echo "  - OpenAPI: http://localhost:8000/api.json"
    echo ""
    
    echo -e "${GREEN}数据库服务:${NC}"
    echo "  - PostgreSQL: localhost:5432"
    echo "  - Redis: localhost:6379"
    echo ""
    
    echo -e "${GREEN}管理界面:${NC}"
    echo "  - Redis管理: http://localhost:8081"
    echo "  - PostgreSQL管理: http://localhost:8082"
    echo ""
    
    echo -e "${GREEN}测试SMS功能:${NC}"
    echo "curl -X POST http://localhost:8000/v1/sms/send \\"
    echo "  -H \"Content-Type: application/json\" \\"
    echo "  -d '{\"phone\":\"13800138000\"}'"
}

# 停止服务
stop_services() {
    print_header "停止服务"
    
    print_message "停止所有服务..."
    docker-compose -f docker-compose.prod.yml down
    
    print_message "服务已停止"
}

# 清理数据
clean_data() {
    print_warning "这将删除所有数据，是否继续？(y/N)"
    read -r response
    if [[ "$response" =~ ^([yY][eE][sS]|[yY])$ ]]; then
        print_header "清理数据"
        
        print_message "停止服务并清理数据..."
        docker-compose -f docker-compose.prod.yml down -v
        
        print_message "清理Docker镜像..."
        docker image prune -f
        
        print_message "数据清理完成"
    else
        print_message "操作已取消"
    fi
}

# 查看日志
view_logs() {
    print_header "查看服务日志"
    
    echo "选择要查看的服务日志:"
    echo "1. 应用服务"
    echo "2. PostgreSQL"
    echo "3. Redis"
    echo "4. 所有服务"
    
    read -p "请选择 (1-4): " choice
    
    case $choice in
        1)
            docker-compose -f docker-compose.prod.yml logs app
            ;;
        2)
            docker-compose -f docker-compose.prod.yml logs postgres
            ;;
        3)
            docker-compose -f docker-compose.prod.yml logs redis
            ;;
        4)
            docker-compose -f docker-compose.prod.yml logs
            ;;
        *)
            print_error "无效选择"
            ;;
    esac
}

# 主菜单
show_menu() {
    echo ""
    echo "=== KG平台部署管理工具 ==="
    echo "1. 一键部署 (构建+启动)"
    echo "2. 仅启动服务"
    echo "3. 启动管理工具"
    echo "4. 检查服务状态"
    echo "5. 查看服务信息"
    echo "6. 停止服务"
    echo "7. 查看日志"
    echo "8. 清理数据"
    echo "9. 退出"
    echo ""
}

# 一键部署
deploy_all() {
    print_header "开始一键部署"
    
    check_docker
    check_config
    build_app
    start_services
    check_health
    show_info
    
    print_message "🎉 部署完成！"
}

# 主函数
main() {
    if [ $# -eq 0 ]; then
        # 交互模式
        while true; do
            show_menu
            read -p "请选择操作 (1-9): " choice
            
            case $choice in
                1)
                    deploy_all
                    ;;
                2)
                    start_services
                    ;;
                3)
                    start_tools
                    ;;
                4)
                    check_health
                    ;;
                5)
                    show_info
                    ;;
                6)
                    stop_services
                    ;;
                7)
                    view_logs
                    ;;
                8)
                    clean_data
                    ;;
                9)
                    print_message "退出"
                    exit 0
                    ;;
                *)
                    print_error "无效选择，请重新输入"
                    ;;
            esac
        done
    else
        # 命令行模式
        case $1 in
            deploy)
                deploy_all
                ;;
            start)
                start_services
                ;;
            tools)
                start_tools
                ;;
            status)
                check_health
                ;;
            info)
                show_info
                ;;
            stop)
                stop_services
                ;;
            logs)
                view_logs
                ;;
            clean)
                clean_data
                ;;
            *)
                echo "用法: $0 [deploy|start|tools|status|info|stop|logs|clean]"
                echo "或者直接运行 $0 进入交互模式"
                exit 1
                ;;
        esac
    fi
}

# 运行主函数
main "$@"
