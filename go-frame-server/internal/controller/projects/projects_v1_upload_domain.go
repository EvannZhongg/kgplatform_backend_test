package projects

import (
	"context"
	"encoding/json"
	"fmt"
	"kgplatform-backend/internal/dao"
	"kgplatform-backend/internal/logic/tasks"
	"kgplatform-backend/internal/logic/upload"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "kgplatform-backend/api/projects/v1"
)

//func getUserIdFromToken(ctx context.Context) (int64, error) {
//	r := ghttp.RequestFromCtx(ctx)
//	tokenString := r.Header.Get("Authorization")
//
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		return []byte(consts.JwtKey), nil
//	})
//	if err != nil || !token.Valid {
//		return 0, gerror.New("无效的认证token")
//	}
//
//	claims, ok := token.Claims.(jwt.MapClaims)
//	if !ok {
//		return 0, gerror.New("无效的token claims")
//	}
//
//	userIdFloat, ok := claims["Id"].(float64)
//	if !ok {
//		return 0, gerror.New("token中缺少用户ID")
//	}
//
//	return int64(userIdFloat), nil
//}

func (c *ControllerV1) UploadSchema(ctx context.Context, req *v1.UploadSchemaReq) (res *v1.UploadSchemaRes, err error) {
	uploadLogic := upload.NewUpload()
	var schemaWrapper tasks.SchemaWrapper
	schemaWrapper.Schemas = req.Triples
	bytes, err := json.Marshal(schemaWrapper)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter)
	}
	// 使用UUID生成唯一标识
	uuidStr := uuid.New().String()
	timestamp := time.Now().Format("20060102150405")

	filename := fmt.Sprintf("%s_%s_%s", req.ProjectId, timestamp, uuidStr[:8])
	saveDataOutput, err := uploadLogic.SaveData(ctx, &upload.SaveDataInput{
		Content:  string(bytes),
		FileName: filename,
		DataType: "json",
	})
	if err != nil {
		return nil, err
	}
	//userId := g.RequestFromCtx(ctx).GetCtxVar("userId").Int()

	userId := g.RequestFromCtx(ctx).GetCtxVar("userID").Int()
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	_, err = dao.Projects.Ctx(ctx).Where("id", req.ProjectId).Where("user_id", userId).Update(g.Map{
		"schema_url": saveDataOutput.FileName,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UploadSchemaRes{}, nil
}
