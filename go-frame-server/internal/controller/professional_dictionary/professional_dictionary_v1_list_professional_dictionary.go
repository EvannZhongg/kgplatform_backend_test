package professional_dictionary

import (
	"context"
	"kgplatform-backend/api/professional_dictionary/v1"
	logic "kgplatform-backend/internal/logic/professional_dictionary"
)

func (c *ControllerV1) ListProfessionalDictionary(ctx context.Context, req *v1.ListProfessionalDictionaryReq) (res *v1.ListProfessionalDictionaryRes, err error) {
	return logic.ListProfessionalDictionary(ctx, req)
}
