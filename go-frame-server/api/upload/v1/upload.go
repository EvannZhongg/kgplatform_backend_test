package v1

import "github.com/gogf/gf/v2/frame/g"

type UploadFileReq struct {
	g.Meta `path:"upload/file" method:"post" sm:"上传文件" tags:"文件上传"`
}

type UploadFileRes struct {
	FileName string `json:"fileName" dc:"文件名"`
	FilePath string `json:"filePath" dc:"文件路径"`
	FileSize int64  `json:"fileSize" dc:"文件大小"`
	FileType string `json:"fileType" dc:"文件类型"`
	FileUrl  string `json:"fileUrl" dc:"文件访问URL"`
}

type SaveDataReq struct {
	g.Meta   `path:"upload/data" method:"post" sm:"保存文本或JSON数据到MinIO" tags:"文件上传"`
	Content  string `json:"content" v:"required|length:1,10000000#内容不能为空|内容长度不能超过10MB" dc:"要保存的内容"`
	FileName string `json:"fileName" v:"required|length:1,255#文件名不能为空|文件名长度不能超过255个字符" dc:"文件名"`
	DataType string `json:"dataType" v:"required|in:text,json#数据类型不能为空|数据类型必须是text或json" dc:"数据类型：text或json"`
}

type SaveDataRes struct {
	FileName string `json:"fileName" dc:"文件名"`
	FilePath string `json:"filePath" dc:"文件路径"`
	FileSize int64  `json:"fileSize" dc:"文件大小"`
	FileType string `json:"fileType" dc:"文件类型"`
	FileUrl  string `json:"fileUrl" dc:"文件访问URL"`
}
