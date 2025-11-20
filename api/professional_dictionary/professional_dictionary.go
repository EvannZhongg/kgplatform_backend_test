// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package professional_dictionary

import (
	"context"

	"kgplatform-backend/api/professional_dictionary/v1"
)

type IProfessionalDictionaryV1 interface {
	CreateProfessionalDictionary(ctx context.Context, req *v1.CreateProfessionalDictionaryReq) (res *v1.CreateProfessionalDictionaryRes, err error)
	ListProfessionalDictionary(ctx context.Context, req *v1.ListProfessionalDictionaryReq) (res *v1.ListProfessionalDictionaryRes, err error)
}
