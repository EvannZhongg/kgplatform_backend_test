# KG平台 Docker 一键部署指南

## 🚀 快速开始

### 一键部署

```bash
# 使用Makefile (推荐)
make deploy

# 或使用部署脚本
./scripts/deploy.sh deploy

# Windows用户
scripts\deploy.bat
```

### 手动部署

```bash
# 1. 构建应用镜像
make deploy-build

# 2. 启动所有服务
make deploy-start

# 3. 检查服务状态
make deploy-status
```

## 📋 部署命令

| 操作 | Makefile命令 | 脚本命令 | 说明 |
|------|-------------|----------|------|
| 一键部署 | `make deploy` | `./scripts/deploy.sh deploy` | 构建+启动所有服务 |
| 构建镜像 | `make deploy-build` | `./scripts/deploy.sh build` | 仅构建应用镜像 |
| 启动服务 | `make deploy-start` | `./scripts/deploy.sh start` | 启动所有服务 |
| 停止服务 | `make deploy-stop` | `./scripts/deploy.sh stop` | 停止所有服务 |
| 查看日志 | `make deploy-logs` | `./scripts/deploy.sh logs` | 查看服务日志 |
| 清理数据 | `make deploy-clean` | `./scripts/deploy.sh clean` | 清理所有数据 |
| 启动工具 | `make deploy-tools` | `./scripts/deploy.sh tools` | 启动管理界面 |
| 检查状态 | `make deploy-status` | `./scripts/deploy.sh status` | 检查服务状态 |

## 🏗️ 服务架构

### 核心服务
- **app**: Go应用服务 (端口: 8000)
- **postgres**: PostgreSQL数据库 (端口: 5432)
- **redis**: Redis缓存 (端口: 6379)

### 管理工具 (可选)
- **redis-commander**: Redis管理界面 (端口: 8081)
- **pgadmin**: PostgreSQL管理界面 (端口: 8082)

## 🔧 配置文件

### 生产环境配置
- `docker-compose.prod.yml`: 生产环境Docker Compose配置
- `manifest/config/config.prod.yaml`: 生产环境应用配置
- `manifest/docker/Dockerfile`: 应用镜像构建文件

### 配置说明

#### 数据库配置
```yaml
database:
  default:
    link: "pgsql:postgres:12345678@tcp(postgres:5432)/kg?sslmode=disable"
```

#### Redis配置
```yaml
redis:
  host: "redis"
  port: 6379
  password: ""
  db: 0
```

#### SMS配置
```yaml
aliyun:
  sms:
    accessKeyId: "your_access_key_id"
    accessKeySecret: "your_access_key_secret"
    signName: "your_sign_name"
    templateCode: "SMS_123456789"
```

## 🌐 服务访问

### 应用服务
- **主服务**: http://localhost:8000
- **API文档**: http://localhost:8000/swagger
- **OpenAPI**: http://localhost:8000/api.json

### 数据库服务
- **PostgreSQL**: localhost:5432
- **Redis**: localhost:6379

### 管理界面
- **Redis管理**: http://localhost:8081
- **PostgreSQL管理**: http://localhost:8082
  - 邮箱: admin@kgplatform.com
  - 密码: admin123

## 🧪 功能测试

### 测试SMS功能

```bash
# 发送验证码
curl -X POST http://localhost:8000/v1/sms/send \
  -H "Content-Type: application/json" \
  -d '{"phone":"13800138000"}'

# 验证验证码
curl -X POST http://localhost:8000/v1/sms/verify \
  -H "Content-Type: application/json" \
  -d '{"phone":"13800138000","code":"123456"}'
```

### 测试用户功能

```bash
# 用户注册
curl -X POST http://localhost:8000/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"123456"}'

# 用户登录
curl -X POST http://localhost:8000/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"123456"}'
```

## 🔍 监控和维护

### 查看服务状态
```bash
make deploy-status
# 或
docker-compose -f docker-compose.prod.yml ps
```

### 查看服务日志
```bash
# 查看所有服务日志
make deploy-logs

# 查看特定服务日志
docker-compose -f docker-compose.prod.yml logs app
docker-compose -f docker-compose.prod.yml logs postgres
docker-compose -f docker-compose.prod.yml logs redis
```

### 重启服务
```bash
# 重启所有服务
docker-compose -f docker-compose.prod.yml restart

# 重启特定服务
docker-compose -f docker-compose.prod.yml restart app
```

## 🛠️ 故障排除

### 常见问题

#### 1. 服务启动失败
```bash
# 检查Docker状态
docker --version
docker-compose --version

# 查看详细日志
docker-compose -f docker-compose.prod.yml logs
```

#### 2. 数据库连接失败
```bash
# 检查PostgreSQL状态
docker-compose -f docker-compose.prod.yml ps postgres

# 测试数据库连接
docker exec -it kgplatform-postgres psql -U postgres -d kg -c "SELECT 1;"
```

#### 3. Redis连接失败
```bash
# 检查Redis状态
docker-compose -f docker-compose.prod.yml ps redis

# 测试Redis连接
docker exec -it kgplatform-redis redis-cli ping
```

#### 4. 应用服务异常
```bash
# 查看应用日志
docker-compose -f docker-compose.prod.yml logs app

# 检查应用健康状态
curl http://localhost:8000/api.json
```

### 性能优化

#### 1. 数据库优化
- 调整PostgreSQL配置参数
- 添加数据库索引
- 配置连接池

#### 2. Redis优化
- 调整内存限制
- 配置持久化策略
- 设置过期策略

#### 3. 应用优化
- 调整Go应用配置
- 配置日志级别
- 设置健康检查

## 🔒 安全配置

### 生产环境安全建议

1. **修改默认密码**
   - PostgreSQL密码
   - 管理界面密码

2. **网络安全**
   - 使用防火墙限制端口访问
   - 配置SSL/TLS证书

3. **数据安全**
   - 定期备份数据库
   - 加密敏感配置

4. **访问控制**
   - 限制管理界面访问
   - 配置API访问控制

## 📊 数据备份

### 备份数据库
```bash
# 备份PostgreSQL
docker exec kgplatform-postgres pg_dump -U postgres kg > backup_$(date +%Y%m%d_%H%M%S).sql

# 备份Redis
docker exec kgplatform-redis redis-cli BGSAVE
docker cp kgplatform-redis:/data/dump.rdb ./backup/
```

### 恢复数据
```bash
# 恢复PostgreSQL
docker exec -i kgplatform-postgres psql -U postgres -d kg < backup.sql

# 恢复Redis
docker cp backup/dump.rdb kgplatform-redis:/data/
docker-compose -f docker-compose.prod.yml restart redis
```

## 🔄 更新部署

### 更新应用
```bash
# 1. 停止服务
make deploy-stop

# 2. 拉取最新代码
git pull

# 3. 重新构建
make deploy-build

# 4. 启动服务
make deploy-start
```

### 滚动更新
```bash
# 更新应用服务
docker-compose -f docker-compose.prod.yml up -d --no-deps app
```

## 📞 技术支持

如果遇到问题，可以：

1. **查看日志**: `make deploy-logs`
2. **检查状态**: `make deploy-status`
3. **重启服务**: `make deploy-stop && make deploy-start`
4. **清理重建**: `make deploy-clean && make deploy`

## 📝 注意事项

1. **首次部署**: 需要下载镜像，可能需要几分钟
2. **数据持久化**: 数据存储在Docker卷中
3. **端口冲突**: 确保8000、5432、6379端口未被占用
4. **资源要求**: 建议至少2GB内存
5. **网络访问**: 确保Docker网络正常工作
