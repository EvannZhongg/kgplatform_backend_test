package upload

import (
	"kgplatform-backend/api/upload"
	uploadLogic "kgplatform-backend/internal/logic/upload"
)

type ControllerV1 struct {
	upload *uploadLogic.Upload
}

func NewV1() upload.IUploadV1 {
	return &ControllerV1{
		upload: uploadLogic.NewUpload(),
	}
}
