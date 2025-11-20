package upload

import (
	"context"
	v1 "kgplatform-backend/api/upload/v1"
	uploadLogic "kgplatform-backend/internal/logic/upload"
)

func (c *ControllerV1) UploadFile(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error) {
	result, err := c.upload.UploadFile(ctx, &uploadLogic.UploadFileInput{})
	if err != nil {
		return nil, err
	}

	return &v1.UploadFileRes{
		FileName: result.FileName,
		FilePath: result.FilePath,
		FileSize: result.FileSize,
		FileType: result.FileType,
		FileUrl:  result.FileUrl,
	}, nil
}

func (c *ControllerV1) SaveData(ctx context.Context, req *v1.SaveDataReq) (res *v1.SaveDataRes, err error) {
	result, err := c.upload.SaveData(ctx, &uploadLogic.SaveDataInput{
		Content:  req.Content,
		FileName: req.FileName,
		DataType: req.DataType,
	})
	if err != nil {
		return nil, err
	}

	return &v1.SaveDataRes{
		FileName: result.FileName,
		FilePath: result.FilePath,
		FileSize: result.FileSize,
		FileType: result.FileType,
		FileUrl:  result.FileUrl,
	}, nil
}
