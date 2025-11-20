# GoFrame Template For SingleRepo

Quick Start: 
- https://goframe.org/quick

# 一键启动

```bash
docker-compose -f docker-compose.prod.yml up -d
```


#安装

## Postgresql


## Redis

### windows
下载地址：https://github.com/MicrosoftArchive/redis/releases


### mac
```bash
    brew install redis
    redis-server /usr/local/etc/redis.conf
```

### linux

```bash
    sudo apt-get install redis-server
    redis-server
```

# 手机号码验证功能配置示例

## 1. Redis配置

确保Redis服务正在运行，默认配置：
- 主机: localhost
- 端口: 6379
- 密码: (空)
- 数据库: 0

如果需要修改，请编辑 `manifest/config/config.yaml` 中的redis配置。

## 2. 阿里云SMS配置

### 2.1 获取AccessKey
1. 登录阿里云控制台
2. 进入"访问控制" -> "用户" -> "创建用户"
3. 创建用户并分配"SMS"相关权限
4. 创建AccessKey，获取AccessKey ID和AccessKey Secret

### 2.2 申请短信签名
1. 进入"短信服务" -> "国内消息" -> "签名管理"
2. 申请短信签名（如：您的应用名称）
3. 等待审核通过

### 2.3 申请短信模板
1. 进入"短信服务" -> "国内消息" -> "模板管理"
2. 申请短信模板，内容示例：
   ```
   您的验证码是${code}，5分钟内有效，请勿泄露给他人。
   ```
3. 等待审核通过

### 2.4 配置参数
编辑 `manifest/config/config.yaml`，替换以下配置：

```yaml
aliyun:
  sms:
    accessKeyId: "LTAI5tM37fRqPKMuWCSRCWNC"  # 替换为你的AccessKey ID
    accessKeySecret: "bZKXUlQotgul0t0t4cp4Jwuof7MYz0"  # 替换为你的AccessKey Secret
    signName: "南京尺牍科技文化"  # 替换为你的短信签名
    templateCode: "SMS_325940806"  # 替换为你的模板代码
```

## 3. 启动服务

1. 确保Redis服务运行
2. 配置好阿里云SMS参数
3. 启动服务：
   ```bash
   go mod tidy
   go run main.go
   ```

或者使用curl命令：
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

## 5. 注意事项

1. **费用说明**: 阿里云SMS按条计费，请合理使用
2. **频率限制**: 建议添加发送频率限制，防止恶意调用
3. **安全配置**: 生产环境请使用环境变量存储敏感配置
4. **监控告警**: 建议监控短信发送失败率
5. **日志记录**: 记录所有验证码操作便于排查问题
