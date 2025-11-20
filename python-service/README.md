# 知识图谱三元组抽取服务 (Python Service)

这是一个基于Flask的RESTful API服务，专门用于从文本文件中抽取知识图谱三元组。该服务支持多种AI模型提供商，提供异步任务处理和实时进度推送功能。

## 📋 目录

- [项目架构](#项目架构)
- [核心功能](#核心功能)
- [技术栈](#技术栈)
- [安装部署](#安装部署)
- [API接口文档](#api接口文档)
- [配置说明](#配置说明)
- [使用示例](#使用示例)
- [项目结构](#项目结构)

## 🏗️ 项目架构

### 整体架构图

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Flask API     │    │  Triplet        │    │   AI Models     │
│   Server        │◄──►│  Extraction     │◄──►│   (DeepSeek/    │
│   (api_server)  │    │  Service        │    │    Qwen/etc)    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         ▼                       ▼                       ▼
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Task Queue    │    │   File          │    │   Prompt        │
│   Management    │    │   Processing    │    │   Generation   │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

### 核心组件

1. **API服务器** (`api_server.py`)
   - Flask应用主入口
   - 路由管理和错误处理
   - SSE实时推送
   - 文件下载和URL处理

2. **三元组抽取服务** (`triplet_extraction_service.py`)
   - 任务生命周期管理
   - 多线程并发处理
   - 进度跟踪和状态更新

3. **文档处理模块** (`extract_triplets_from_docx.py`)
   - 文本文件读取和分块
   - AI模型调用
   - 三元组解析和输出

4. **Prompt生成器** (`gen_triplet_prompt.py`)
   - 基于Schema的Prompt构建
   - 样例文本处理
   - 动态模板生成

## 🚀 核心功能

- **多格式文件支持**: 支持.txt文件处理，支持URL下载（自动下载到上传目录）
- **多种文件输入方式**: 支持绝对路径、上传目录文件名、URL三种方式
- **Material ID关联**: 支持在文件列表中关联material_id，便于追踪文件来源
- **多AI模型支持**: DeepSeek、Qwen、通用OpenAI兼容接口
- **异步任务处理**: 多线程并发，支持任务队列管理
- **实时进度推送**: SSE (Server-Sent Events) 实时状态更新
- **智能Prompt生成**: 基于知识图谱Schema自动生成抽取提示，支持自定义领域、词典、优先级等
- **三元组溯源**: 输出结果包含chunk索引和源文本，便于追溯三元组来源
- **任务管理**: 创建、查询、取消、清理任务
- **错误处理**: 完善的异常处理和错误恢复机制

## 🛠️ 技术栈

### 后端框架
- **Flask 2.3.3**: Web框架
- **Flask-CORS 4.0.0**: 跨域支持
- **Werkzeug 2.3.7**: WSGI工具库

### 数据处理
- **pandas 1.3.0+**: 数据处理
- **openpyxl 3.0.0+**: Excel文件处理
- **python-docx 0.8.11+**: Word文档处理

### HTTP客户端
- **requests 2.31.0**: HTTP请求库

### 其他依赖
- **python-dotenv 1.0.0**: 环境变量管理
- **pathlib**: 路径处理
- **concurrent.futures**: 并发处理
- **threading**: 多线程支持

## 📦 安装部署

### 环境要求

- Python 3.8+
- 至少2GB可用内存
- 网络连接（用于AI模型调用）

### 安装步骤

1. **克隆项目**
```bash
cd /home/ubuntu/ChiDu/Proj/KG/kgplatform-backend/python-service
```

2. **安装系统依赖**
```bash
# Ubuntu/Debian系统需要先安装python3-venv
sudo apt update
sudo apt install -y python3.12-venv python3-full
```

3. **创建虚拟环境**
```bash
# 创建虚拟环境
python3 -m venv venv

# 激活虚拟环境
source venv/bin/activate
```

4. **安装Python依赖**
```bash
# 确保虚拟环境已激活，然后安装依赖
pip install -r requirements.txt
```

5. **配置环境变量**
```bash
cp env.example .env
# 编辑.env文件，设置必要的环境变量
```

6. **启动服务**
```bash
# 确保虚拟环境已激活
source venv/bin/activate

# 开发模式
python api_server.py --debug

# 生产模式
python api_server.py --host 0.0.0.0 --port 8001 --max-workers 5
```

### 虚拟环境使用说明

**激活虚拟环境:**
```bash
source venv/bin/activate
```

**退出虚拟环境:**
```bash
deactivate
```

**每次使用前都需要激活虚拟环境:**
```bash
# 进入项目目录
cd /home/ubuntu/ChiDu/Proj/KG/kgplatform-backend/python-service

# 激活虚拟环境
source venv/bin/activate

# 运行服务
python api_server.py
```

### Docker部署（可选）

```dockerfile
FROM python:3.9-slim

WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt

COPY . .
EXPOSE 8001

CMD ["python", "api_server.py", "--host", "0.0.0.0", "--port", "8001"]
```

## 📚 API接口文档

### 基础信息

- **Base URL**: `http://localhost:8001`（默认端口，可通过`--port`参数修改）
- **Content-Type**: `application/json`
- **字符编码**: UTF-8

### 接口列表

#### 1. 健康检查

```http
GET /health
```

**响应示例:**
```json
{
  "status": "healthy",
  "timestamp": "2024-01-01T12:00:00.000Z",
  "service": {
    "service_status": "running",
    "total_tasks": 5,
    "processing_tasks": 1,
    "completed_tasks": 3,
    "failed_tasks": 1,
    "max_workers": 3,
    "active_workers": 1
  }
}
```

#### 2. 创建三元组抽取任务

```http
POST /api/v1/tasks
```

**请求参数:**
```json
{
  "files": [
    // 格式1：字符串数组（旧格式）
    "file1.txt", 
    "file2.txt",
    
    // 格式2：对象数组（新格式，支持material_id关联）
    {"material_id": 1, "url": "file1.txt"},
    {"material_id": 2, "url": "file2.txt"}
  ],
  // 注意：files中的所有项必须为同一格式（全为字符串或全为对象）
  // 且必须全为绝对路径、全为上传目录文件名、或全为URL（不可混合）
  
  "prompt_text": "请从以下文档中抽取知识图谱三元组...",
  "provider": "deepseek",               // deepseek, qwen, forward
  "model": "deepseek-chat",            // 可选，使用默认模型
  "api_key": "sk-xxx...",              // AI服务API密钥
  "base_url": "https://api.deepseek.com" // 可选，自定义API地址
}
```

**文件路径说明:**
- **绝对路径**: 如 `E:\path\to\file.txt`（Windows）或 `/home/user/file.txt`（Linux）
- **上传目录文件名**: 如 `file.txt`（相对于配置的上传目录）
- **URL**: 如 `https://example.com/file.txt`（自动下载到上传目录，仅支持.txt格式）

**响应示例:**
```json
{
  "task_id": "550e8400-e29b-41d4-a716-446655440000",
  "status": "created",
  "message": "任务创建成功"
}
```

#### 3. 获取任务状态

```http
GET /api/v1/tasks/{task_id}
```

**响应示例:**
```json
{
  "task_id": "550e8400-e29b-41d4-a716-446655440000",
  "status": "completed",
  "total_files": 2,
  "processed_files": 2,
  "failed_files": 0,
  "total_triples": 45,
  "processing_time": "2m30s",
  "results": [
    {
      "file_name": "file1.txt",
      "material_id": 1,  // 如果创建任务时提供了material_id
      "status": "success",
      "triples_count": 23,
      "output_files": {
        "txt": "/path/to/output/txt/file1_extractions.txt",
        "jsonl": "/path/to/output/jsonl/file1_extractions.jsonl",
        "chunks": "/path/to/output/txt/file1_chunks.txt"
      }
    }
  ],
  "errors": [],
  "created_at": "2024-01-01T12:00:00.000Z",
  "updated_at": "2024-01-01T12:02:30.000Z"
}
```

#### 4. 取消任务

```http
DELETE /api/v1/tasks/{task_id}
```

**响应示例:**
```json
{
  "task_id": "550e8400-e29b-41d4-a716-446655440000",
  "status": "cancelled",
  "message": "任务取消成功"
}
```

#### 5. 获取任务列表

```http
GET /api/v1/tasks?status=completed&limit=10&offset=0
```

**查询参数:**
- `status`: 任务状态过滤 (可选)
- `limit`: 返回数量限制 (默认10，最大100)
- `offset`: 偏移量 (默认0)

**响应示例:**
```json
{
  "tasks": [
    {
      "task_id": "550e8400-e29b-41d4-a716-446655440000",
      "status": "completed",
      "total_files": 2,
      "processed_files": 2,
      "failed_files": 0,
      "total_triples": 45,
      "processing_time": "2m30s",
      "created_at": "2024-01-01T12:00:00.000Z",
      "updated_at": "2024-01-01T12:02:30.000Z"
    }
  ],
  "total": 1,
  "limit": 10,
  "offset": 0
}
```

#### 6. 任务进度实时推送 (SSE)

```http
GET /api/v1/tasks/{task_id}/stream
```

**响应格式:** Server-Sent Events
```
data: {"task_id": "xxx", "status": "processing", "processed_files": 1, "total_files": 2}

data: {"task_id": "xxx", "status": "completed", "total_triples": 45}
```

#### 7. 生成三元组抽取Prompt

```http
POST /api/v1/genprompt
```

**请求参数:**
```json
{
  "schema_url": "https://example.com/schema.json",           // 必需：知识图谱Schema URL
  "sample_text_url": "https://example.com/sample.txt",       // 可选：样例文本URL（支持.txt, .docx格式）
  "sample_xlsx_url": "https://example.com/sample.xlsx",     // 可选：样例三元组Excel URL（.xlsx格式）
  "target_domain": "建筑学",                                  // 可选：目标领域描述
  "dictionary_url": "https://example.com/dictionary.txt",   // 可选：专业词典URL（.txt格式）
  "priority_extractions": ["城市", "建筑师", "设计作品"],    // 可选：抽取意向优先级列表
  "extraction_requirements": "重点关注建筑风格和设计理念",  // 可选：抽取要求描述
  "base_instruction": "自定义的基础指导语..."                // 可选：自定义基础指导语（不提供则使用默认）
}
```

**响应示例:**
```json
{
  "prompt": "你是东南大学建筑学院的一个行政人员...",
  "schema_url": "https://example.com/schema.json",
  "sample_text_url": "https://example.com/sample.txt",
  "sample_xlsx_url": "https://example.com/sample.xlsx",
  "target_domain": "建筑学",
  "dictionary_url": "https://example.com/dictionary.txt",
  "priority_extractions": ["城市", "建筑师", "设计作品"],
  "extraction_requirements": "重点关注建筑风格和设计理念",
  "message": "Prompt生成成功"
}
```

#### 8. 服务状态查询

```http
GET /api/v1/service/status
```

**响应示例:**
```json
{
  "service_status": "running",
  "total_tasks": 5,
  "processing_tasks": 1,
  "completed_tasks": 3,
  "failed_tasks": 1,
  "max_workers": 3,
  "active_workers": 1
}
```

#### 9. 清理旧任务

```http
POST /api/v1/service/cleanup
```

**请求参数:**
```json
{
  "days": 7  // 清理多少天前的任务，默认7天
}
```

**响应示例:**
```json
{
  "message": "清理完成",
  "removed_tasks": 3
}
```

## ⚙️ 配置说明

### 环境变量

| 变量名 | 默认值 | 说明 |
|--------|--------|------|
| `HOST` | `0.0.0.0` | 服务监听地址 |
| `PORT` | `8001` | 服务监听端口（可通过命令行参数覆盖） |
| `DEBUG` | `false` | 调试模式 |
| `MAX_WORKERS` | `3` | 最大工作线程数 |
| `OUTPUT_BASE_DIR` | `output` | 输出文件基础目录 |
| `UPLOAD_DIR` | `E:\GoLandProject\exam\item_reslove\gin-task-server\uploads` | 上传文件目录（用于文件名和URL下载） |
| `TASK_CLEANUP_DAYS` | `7` | 任务清理天数 |
| `DEFAULT_CHUNK_SIZE` | `800` | 默认文本分块大小 |
| `MAX_FILES_PER_TASK` | `50` | 单任务最大文件数 |
| `MAX_TASKS_PER_LIST` | `100` | 任务列表最大返回数 |
| `LOG_LEVEL` | `INFO` | 日志级别 |
| `LOG_FILE` | `api_server.log` | 日志文件路径 |

### AI模型配置

#### DeepSeek
```bash
export DEEPSEEK_API_KEY="sk-xxx..."
```

#### Qwen (DashScope)
```bash
export DASHSCOPE_API_KEY="sk-xxx..."
```

#### 通用代理
```bash
export FORWARD_BASE_URL="http://localhost:3000/v1"
export FORWARD_API_KEY="sk-xxx..."
export FORWARD_DEFAULT_MODEL="gpt-4o-mini"
```

## 💡 使用示例

### Python客户端示例

```python
import requests
import json
import time

# 服务地址
BASE_URL = "http://localhost:8001"

# 1. 创建任务
def create_task():
    data = {
        "files": ["sample.txt"],
        "prompt_text": "请从以下文档中抽取知识图谱三元组...",
        "provider": "deepseek",
        "api_key": "sk-xxx..."
    }
    
    response = requests.post(f"{BASE_URL}/api/v1/tasks", json=data)
    return response.json()["task_id"]

# 2. 监控任务进度
def monitor_task(task_id):
    while True:
        response = requests.get(f"{BASE_URL}/api/v1/tasks/{task_id}")
        status = response.json()
        
        print(f"状态: {status['status']}")
        print(f"进度: {status['processed_files']}/{status['total_files']}")
        
        if status['status'] in ['completed', 'failed', 'cancelled']:
            break
            
        time.sleep(2)

# 3. 获取结果
def get_results(task_id):
    response = requests.get(f"{BASE_URL}/api/v1/tasks/{task_id}")
    return response.json()

# 使用示例
if __name__ == "__main__":
    task_id = create_task()
    monitor_task(task_id)
    results = get_results(task_id)
    print(json.dumps(results, indent=2, ensure_ascii=False))
```

### JavaScript客户端示例

```javascript
// 创建任务
async function createTask() {
    const response = await fetch('http://localhost:8001/api/v1/tasks', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            files: ['sample.txt'],
            prompt_text: '请从以下文档中抽取知识图谱三元组...',
            provider: 'deepseek',
            api_key: 'sk-xxx...'
        })
    });
    
    const result = await response.json();
    return result.task_id;
}

// SSE实时监控
function monitorTask(taskId) {
    const eventSource = new EventSource(`http://localhost:8001/api/v1/tasks/${taskId}/stream`);
    
    eventSource.onmessage = function(event) {
        const data = JSON.parse(event.data);
        console.log('任务状态更新:', data);
        
        if (['completed', 'failed', 'cancelled'].includes(data.status)) {
            eventSource.close();
        }
    };
    
    eventSource.onerror = function(event) {
        console.error('SSE连接错误:', event);
        eventSource.close();
    };
}

// 使用示例
createTask().then(taskId => {
    console.log('任务ID:', taskId);
    monitorTask(taskId);
});
```

### cURL示例

```bash
# 1. 健康检查
curl -X GET http://localhost:8001/health

# 2. 创建任务
curl -X POST http://localhost:8001/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "files": ["sample.txt"],
    "prompt_text": "请从以下文档中抽取知识图谱三元组...",
    "provider": "deepseek",
    "api_key": "sk-xxx..."
  }'

# 3. 查询任务状态
curl -X GET http://localhost:8001/api/v1/tasks/{task_id}

# 4. 生成Prompt
curl -X POST http://localhost:8001/api/v1/genprompt \
  -H "Content-Type: application/json" \
  -d '{
    "schema_url": "https://example.com/schema.json"
  }'
```

## 📁 项目结构

```
python-service/
├── api_server.py                 # Flask API服务器主文件
├── triplet_extraction_service.py # 三元组抽取服务核心
├── extract_triplets_from_docx.py # 文档处理和AI调用
├── gen_triplet_prompt.py         # Prompt生成器
├── config.py                     # 配置管理
├── knowledge_graph_schema.json   # 知识图谱Schema定义
├── triple_extraction_prompt.txt  # 三元组抽取提示模板
├── requirements.txt              # Python依赖包
├── env.example                   # 环境变量示例
├── api_server.log               # 服务日志文件
├── partydata/                    # 测试数据目录
│   └── *.txt                    # 测试文本文件
├── output/                       # 输出文件目录
│   └── {task_id}/               # 按任务ID组织的输出
│       ├── txt/                 # 文本格式输出
│       └── jsonl/               # JSONL格式输出
└── test_*.py                    # 测试文件
    ├── test_url_download.py     # URL下载测试
    ├── test_genprompt_formats.py # Prompt生成测试
    └── test_filename_handling.py # 文件名处理测试
```

### 核心文件说明

- **`api_server.py`**: Flask应用入口，包含所有API路由和中间件
- **`triplet_extraction_service.py`**: 任务管理核心，处理异步任务和状态跟踪
- **`extract_triplets_from_docx.py`**: 文档处理逻辑，包含AI模型调用和三元组解析
- **`gen_triplet_prompt.py`**: 基于Schema动态生成抽取提示
- **`config.py`**: 配置类定义，支持开发/生产/测试环境
- **`knowledge_graph_schema.json`**: 知识图谱结构定义，用于Prompt生成

### 输出文件格式

#### 文本格式 (.txt)
```
实践活动: 南师附小弹性离校活动
三元组列表:
- head: {label: "南师附小弹性离校活动", type: "实践活动"}
  relationship: {label: "实践类型", type: "实践类型"}
  tail: {label: "支教", type: "实践类型"}
```

#### JSONL格式 (.jsonl)
```json
[
  {
    "head": {"label": "南师附小弹性离校活动", "type": "实践活动"},
    "relationship": {"label": "实践类型", "type": "实践类型"},
    "tail": {"label": "支教", "type": "实践类型"},
    "_chunk_index": 0,
    "_source_text": "原文中对应的chunk文本内容..."
  },
  {
    "head": {"label": "志愿者团队", "type": "团队"},
    "relationship": {"label": "实践", "type": "实践"},
    "tail": {"label": "南师附小弹性离校活动", "type": "实践活动"},
    "_chunk_index": 1,
    "_source_text": "原文中对应的chunk文本内容..."
  }
]
```

**注意**: JSONL文件中的每个三元组包含以下额外字段：
- `_chunk_index`: 该三元组来自哪个文本分块（从0开始）
- `_source_text`: 该三元组对应的源文本内容（便于溯源）

## 🔧 开发指南

### 添加新的AI模型提供商

1. 在 `extract_triplets_from_docx.py` 中的 `PROVIDERS` 字典添加新配置
2. 实现相应的API调用逻辑
3. 更新文档和测试用例

### 扩展文件格式支持

1. 在 `extract_triplets_from_docx.py` 中添加新的文件读取函数
2. 更新 `triplet_extraction_service.py` 中的文件验证逻辑
3. 修改API接口的参数验证

### 自定义Prompt模板

1. 修改 `triple_extraction_prompt.txt` 模板文件
2. 或在 `gen_triplet_prompt.py` 中自定义生成逻辑
3. 更新Schema结构以支持新的实体和关系类型

## 🐛 故障排除

### 常见问题

1. **任务创建失败**
   - 检查文件路径是否正确
   - 确认API密钥有效
   - 验证网络连接

2. **AI模型调用失败**
   - 检查API密钥和Base URL
   - 确认模型名称正确
   - 查看服务日志获取详细错误信息

3. **文件处理错误**
   - 确认文件格式为.txt（仅支持.txt格式）
   - 检查文件编码（支持UTF-8、GBK、GB2312等多种编码）
   - 验证文件大小不超过限制（URL下载限制100MB）
   - 确认文件路径格式正确（绝对路径、上传目录文件名或URL，不可混合）

4. **SSE连接断开**
   - 检查网络稳定性
   - 确认客户端支持SSE
   - 查看服务器日志

### 日志分析

服务日志保存在 `api_server.log` 文件中，包含：
- 请求和响应信息
- 任务处理进度
- 错误和异常信息
- 性能指标

### 性能优化

1. **调整工作线程数**: 根据CPU核心数调整 `MAX_WORKERS`
2. **优化文本分块**: 调整 `DEFAULT_CHUNK_SIZE` 平衡处理速度和内存使用
3. **定期清理**: 使用 `/api/v1/service/cleanup` 接口清理旧任务
4. **监控资源**: 定期检查内存和磁盘使用情况

## 📄 许可证

本项目采用 MIT 许可证，详见 LICENSE 文件。

## 🤝 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 📞 联系方式

如有问题或建议，请通过以下方式联系：

- 项目Issues: [GitHub Issues](https://github.com/your-repo/issues)
- 邮箱: your-email@example.com

---

**最后更新**: 2025年1月
