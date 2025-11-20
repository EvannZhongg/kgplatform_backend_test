package professional_dictionary

import (
	"context"
	"kgplatform-backend/api/professional_dictionary/v1"
	logic "kgplatform-backend/internal/logic/professional_dictionary"
)

func (c *ControllerV1) CreateProfessionalDictionary(ctx context.Context, req *v1.CreateProfessionalDictionaryReq) (res *v1.CreateProfessionalDictionaryRes, err error) {
	return logic.CreateProfessionalDictionary(ctx, req)
}
