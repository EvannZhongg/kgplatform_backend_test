// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package chat

import (
	"context"

	"kgplatform-backend/api/chat/v1"
)

type IChatV1 interface {
	GeneratePrompt(ctx context.Context, req *v1.GeneratePromptReq) (res *v1.GeneratePromptRes, err error)
}
