ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"

include ./hack/hack-cli.mk
include ./hack/hack.mk

# Docker Redis 管理命令
.PHONY: redis-start redis-stop redis-restart redis-status redis-logs redis-clean

redis-start:
	@echo "启动Redis服务..."
	docker-compose up -d redis
	@echo "Redis服务已启动，端口: 6379"

redis-start-all:
	@echo "启动Redis服务和管理界面..."
	docker-compose up -d
	@echo "Redis服务已启动，端口: 6379"
	@echo "管理界面: http://localhost:8081"

redis-stop:
	@echo "停止Redis服务..."
	docker-compose down
	@echo "Redis服务已停止"

redis-restart:
	@echo "重启Redis服务..."
	docker-compose restart redis
	@echo "Redis服务已重启"

redis-status:
	@echo "Redis服务状态:"
	docker-compose ps

redis-logs:
	@echo "Redis服务日志:"
	docker-compose logs redis

redis-clean:
	@echo "清理Redis数据..."
	docker-compose down -v
	@echo "Redis数据已清理"

# 开发环境快速启动
dev-start: redis-start
	@echo "开发环境已启动"
	@echo "Redis: localhost:6379"
	@echo "应用: go run main.go"

dev-stop: redis-stop
	@echo "开发环境已停止"

# Docker 生产环境部署命令
.PHONY: deploy deploy-build deploy-start deploy-stop deploy-logs deploy-clean deploy-tools

deploy:
	@echo "开始一键部署..."
	@chmod +x scripts/deploy.sh
	@./scripts/deploy.sh deploy

deploy-build:
	@echo "构建应用镜像..."
	docker-compose -f docker-compose.prod.yml build app

deploy-start:
	@echo "启动生产环境服务..."
	docker-compose -f docker-compose.prod.yml up -d

deploy-stop:
	@echo "停止生产环境服务..."
	docker-compose -f docker-compose.prod.yml down

deploy-logs:
	@echo "查看服务日志..."
	docker-compose -f docker-compose.prod.yml logs

deploy-clean:
	@echo "清理生产环境数据..."
	docker-compose -f docker-compose.prod.yml down -v
	docker image prune -f

deploy-tools:
	@echo "启动管理工具..."
	docker-compose -f docker-compose.prod.yml --profile tools up -d
	@echo "Redis管理界面: http://localhost:8081"
	@echo "PostgreSQL管理界面: http://localhost:8082"

# 部署状态检查
deploy-status:
	@echo "检查服务状态..."
	docker-compose -f docker-compose.prod.yml ps

# 部署信息显示
deploy-info:
	@echo "=== 服务信息 ==="
	@echo "应用服务: http://localhost:8000"
	@echo "API文档: http://localhost:8000/swagger"
	@echo "PostgreSQL: localhost:5432"
	@echo "Redis: localhost:6379"
	@echo "Redis管理: http://localhost:8081"
	@echo "PostgreSQL管理: http://localhost:8082"

# Docker网络问题解决
.PHONY: fix-network fix-docker-network

fix-network:
	@echo "解决Docker网络问题..."
	@chmod +x scripts/fix-docker-network.sh
	@./scripts/fix-docker-network.sh

fix-docker-network: fix-network