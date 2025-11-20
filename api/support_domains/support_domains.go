// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package support_domains

import (
	"context"

	"kgplatform-backend/api/support_domains/v1"
)

type ISupportDomainsV1 interface {
	CreateDomain(ctx context.Context, req *v1.CreateDomainReq) (res *v1.CreateDomainRes, err error)
	ListDomains(ctx context.Context, req *v1.ListDomainsReq) (res *v1.ListDomainsRes, err error)
}
