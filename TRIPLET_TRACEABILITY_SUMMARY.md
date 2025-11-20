# 三元组溯源功能实现总结

## ✅ 已完成的工作

### 1. 数据结构扩展
**文件**: `go-frame-server/internal/neo4j/neo4j.go`

为 `SimpleTriple` 结构添加了溯源信息：
```go
type SimpleTriple struct {
    Head         TripleEntity `json:"head"`
    Relationship TripleEntity `json:"relationship"`
    Tail         TripleEntity `json:"tail"`
    SourceInfo *TripleSourceInfo `json:"sourceInfo,omitempty"`  // 新增
}

type TripleSourceInfo struct {
    MaterialId   int    `json:"materialId"`
    MaterialName string `json:"materialName"`
    ChunkIndex   int    `json:"chunkIndex"`
    SourceText   string `json:"sourceText"`
    ChunkStart   int    `json:"chunkStart"`
    ChunkEnd     int    `json:"chunkEnd"`
}
```

### 2. Python服务修改
**文件**: `python-service/extract_triplets_from_docx.py`

修改 `save_outputs()` 函数，为每个三元组添加溯源信息：
- 添加 `_chunk_index` 字段：记录三元组来自哪个文本块
- 添加 `_source_text` 字段：保存原始文本片段

### 3. 后端聚合逻辑修改
**文件**: `go-frame-server/external/py_service/client.go`

在聚合各素材的三元组时，提取并保存溯源信息：
- 解析Python服务返回的带溯源信息的三元组
- 转换为 `SimpleTriple` 结构并附加 `SourceInfo`
- 保存到项目三元组文件中

新增辅助函数 `getStringValue()` 用于安全地从map中提取字符串值。

### 4. 逻辑层实现
**文件**: `go-frame-server/internal/logic/projects/projects.go`

实现了 `GetTripleSourceInfo()` 方法：
- 接收项目ID和三元组信息
- 如果三元组本身包含溯源信息，直接返回
- 否则遍历所有素材，在素材三元组中查找匹配的三元组
- 返回素材ID、素材名称、来源文本、Chunk索引等信息

包含辅助方法：
- `matchTriple()`: 判断两个三元组是否匹配
- `compareEntity()`: 比较三元组的实体部分

### 5. API接口定义
**文件**: `go-frame-server/api/projects/v1/projects.go`

定义了新的API接口：
```go
type GetTripleSourceInfoReq struct {
    g.Meta    `path:"projects/{projectId}/triplets/source" method:"post"`
    ProjectId int                `json:"projectId"`
    Triple    neo4j.SimpleTriple `json:"triple"`
}

type GetTripleSourceInfoRes struct {
    MaterialId   int    `json:"materialId"`
    MaterialName string `json:"materialName"`
    SourceText   string `json:"sourceText"`
    ChunkIndex   int    `json:"chunkIndex"`
}
```

### 6. Controller实现
**文件**: `go-frame-server/internal/controller/projects/projects_v1_get_triple_source_info.go`

创建了Controller方法，调用逻辑层接口。

**文件**: `go-frame-server/api/projects/projects.go`

在接口定义中添加了 `GetTripleSourceInfo` 方法。

## 📋 数据流程

```
1. Python服务提取三元组
   ↓
2. 保存时添加 _chunk_index 和 _source_text
   ↓
3. Go后端聚合三元组时提取溯源信息
   ↓
4. 保存到项目三元组文件（包含SourceInfo字段）
   ↓
5. 前端调用API获取溯源信息
   ↓
6. 显示三元组的来源文件和使用户看到来源文本
```

## 🎯 API使用示例

### 请求示例
```bash
POST /api/projects/{projectId}/triplets/source
Content-Type: application/json

{
    "projectId": 123,
    "triple": {
        "head": {
            "label": "南师附小弹性离校活动",
            "type": "实践活动"
        },
        "relationship": {
            "label": "实践类型",
            "type": "实践类型"
        },
        "tail": {
            "label": "支教",
            "type": "实践类型"
        }
    }
}
```

### 响应示例
```json
{
    "materialId": 456,
    "materialName": "南师附小弹性离校活动.txt",
    "sourceText": "2019年10月，南京师范大学附属小学...",
    "chunkIndex": 2
}
```

## 🚀 前端集成建议

前端可以这样使用：

```javascript
// 点击三元组时调用
async function showTripletSource(triplet) {
    const response = await fetch(
        `/api/projects/${projectId}/triplets/source`,
        {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                projectId: projectId,
                triple: triplet
            })
        }
    );
    
    const data = await response.json();
    
    // 显示来源信息
    showModal({
        title: '三元组来源',
        content: `
            <h3>来源文件: ${data.materialName}</h3>
            <p>文本片段:</p>
            <pre>${data.sourceText}</pre>
        `
    });
}
```

## ⚠️ 注意事项

1. **向后兼容**: 旧三元组可能不含 `SourceInfo`，已在实现中兼容
2. **性能考虑**: 素材较多时，查找可能较慢，建议添加缓存
3. **存储空间**: `_source_text` 增加存储，但提供更完整的追溯能力

## 🔜 后续优化建议

1. **缓存优化**: 对素材三元组数据添加缓存
2. **批量查询**: 支持一次查询多个三元组的来源信息
3. **全文搜索**: 支持在来源文本中进行关键词搜索
4. **高亮显示**: 在前端高亮显示抽取到的实体和关系

## 📝 测试建议

1. 单元测试：测试三元组匹配逻辑
2. 集成测试：测试端到端的溯源查询流程
3. 性能测试：测试大量素材情况下的查询性能

