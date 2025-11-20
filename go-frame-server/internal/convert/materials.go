package convert

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"kgplatform-backend/internal/logic/materials"
	"kgplatform-backend/internal/logic/upload"
	"kgplatform-backend/internal/model/entity"
	"kgplatform-backend/internal/utils"
)

func ConvertMaterialsToDetail(ctx context.Context, material *entity.Materials) (*materials.MaterialDetail, error) {
	uploadLogic := upload.NewUpload()
	detail := &materials.MaterialDetail{
		Id:        material.Id,
		Enable:    material.Enable,
		ProjectId: material.ProjectId,
		CreatedAt: material.CreatedAt,
		UpdatedAt: material.UpdatedAt,
	}
	if material.Url != "" {
		detail.Url = uploadLogic.GenerateFileUrl(ctx, material.Url)
	}
	if material.TextUrl != "" {
		detail.Text = uploadLogic.GenerateFileUrl(ctx, material.TextUrl)
	}
	if material.TripleUrl != "" {
		tripleUrl := uploadLogic.GenerateFileUrl(ctx, material.TripleUrl)
		tripleJsonString, err := utils.DownloadTextFromURL(ctx, tripleUrl)
		if err != nil {
			g.Log().Errorf(ctx, "下载材料抽取结果失败: %v", err)
			return nil, err
		}
		err = json.Unmarshal([]byte(tripleJsonString), &detail.Triples)
		if err != nil {
			g.Log().Errorf(ctx, "解析材料抽取结果失败: %v", err)
			return nil, err
		}
	}

	return detail, nil
}

func ConvertMaterialListToDetailList(ctx context.Context, materialList []*entity.Materials) ([]*materials.MaterialDetail, error) {
	var detailList []*materials.MaterialDetail
	for _, material := range materialList {
		detail, err := ConvertMaterialsToDetail(ctx, material)
		if err != nil {
			return nil, err
		}
		detailList = append(detailList, detail)
	}
	return detailList, nil
}
