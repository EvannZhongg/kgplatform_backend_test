# 文件上传接口说明

## 接口概述

本项目实现了一个完整的文件上传功能，支持多种文件类型的上传，并提供了安全的文件存储和访问机制。

## API 接口

### 上传文件

**接口地址**: `POST /v1/upload/file`

**请求方式**: `multipart/form-data`

**请求参数**:
- `file`: 文件对象（必须）

**响应格式**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "fileName": "example_20241224153045_a1b2c3d4.jpg",
    "filePath": "./resource/public/uploads/example_20241224153045_a1b2c3d4.jpg",
    "fileSize": 1024000,
    "fileType": "image/jpeg",
    "fileUrl": "http://localhost:8000/uploads/example_20241224153045_a1b2c3d4.jpg"
  }
}
```

**响应字段说明**:
- `fileName`: 存储的文件名（包含时间戳和UUID）
- `filePath`: 服务器上的文件路径
- `fileSize`: 文件大小（字节）
- `fileType`: 文件MIME类型
- `fileUrl`: 文件的访问URL

## 配置说明

在 `manifest/config/config.yaml` 中的上传配置：

```yaml
# 文件上传配置
upload:
  path: "./resource/public/uploads"      # 文件上传路径
  maxSize: 10485760                      # 最大文件大小（10MB）
  allowedTypes:                          # 允许的文件类型
    - "image/jpeg"
    - "image/png"
    - "image/gif"
    - "application/pdf"
    - "text/plain"
    - "application/msword"
    - "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
```

### 配置参数说明
- `path`: 文件上传保存路径
- `maxSize`: 文件大小限制（字节），默认10MB
- `allowedTypes`: 允许上传的文件MIME类型列表

## 文件命名规则

上传的文件会自动重命名，格式为：
```
{原文件名}_{时间戳}_{UUID前8位}.{扩展名}
```

例如：`photo.jpg` → `photo_20241224153045_a1b2c3d4.jpg`

## 文件访问

上传成功后，文件可以通过以下方式访问：
- 直接通过返回的 `fileUrl` 访问
- 通过 `/uploads/{fileName}` 路径访问

## 使用示例

### 1. 使用 curl 命令
```bash
curl -X POST \
  http://localhost:8000/v1/upload/file \
  -H 'Content-Type: multipart/form-data' \
  -F 'file=@/path/to/your/file.jpg'
```

### 2. 使用 JavaScript (前端)
```javascript
const formData = new FormData();
formData.append('file', fileInput.files[0]);

fetch('http://localhost:8000/v1/upload/file', {
    method: 'POST',
    body: formData
})
.then(response => response.json())
.then(data => {
    if (data.code === 0) {
        console.log('上传成功:', data.data.fileUrl);
    } else {
        console.error('上传失败:', data.message);
    }
});
```

### 3. 使用测试页面
项目根目录下的 `test_upload.html` 提供了一个完整的测试界面，可以直接在浏览器中打开测试文件上传功能。

## 错误处理

常见错误及解决方案：

1. **文件大小超限**
   - 错误信息：`文件大小不能超过 X MB`
   - 解决方案：减小文件大小或调整配置中的 `maxSize`

2. **文件类型不支持**
   - 错误信息：`不支持的文件类型: xxx`
   - 解决方案：检查文件类型是否在 `allowedTypes` 配置中

3. **未选择文件**
   - 错误信息：`请选择要上传的文件`
   - 解决方案：确保请求中包含文件

4. **目录权限问题**
   - 错误信息：相关的文件系统错误
   - 解决方案：确保上传目录存在且有写入权限

## 安全考虑

1. **文件类型验证**: 通过MIME类型白名单限制可上传的文件类型
2. **文件大小限制**: 防止大文件占用过多服务器空间
3. **文件名重命名**: 避免文件名冲突和潜在的安全风险
4. **存储隔离**: 上传文件存储在专门的目录中

## 目录结构

```
resource/
└── public/
    └── uploads/          # 文件上传目录
        ├── file1.jpg
        ├── file2.pdf
        └── ...
```

## 扩展功能

如需扩展功能，可以考虑：

1. **图片处理**: 自动生成缩略图、压缩等
2. **云存储**: 支持上传到阿里云OSS、AWS S3等
3. **文件分类**: 按文件类型或日期自动分类存储
4. **权限控制**: 基于用户权限的文件访问控制
5. **文件管理**: 提供文件列表、删除等管理功能

## 故障排除

1. **确保服务器已启动**: `go run main.go`
2. **检查端口占用**: 默认端口8000
3. **验证配置文件**: 确保配置文件格式正确
4. **检查目录权限**: 确保上传目录可写
5. **查看日志**: 检查服务器日志获取详细错误信息
