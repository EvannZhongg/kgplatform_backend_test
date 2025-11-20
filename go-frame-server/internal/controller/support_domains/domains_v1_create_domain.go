// 方法：ControllerV1.CreateDomain
package support_domains

import (
	"context"
	"kgplatform-backend/api/support_domains/v1"
	domainsLogic "kgplatform-backend/internal/logic/support_domains"
)

func (c *ControllerV1) CreateDomain(ctx context.Context, req *v1.CreateDomainReq) (res *v1.CreateDomainRes, err error) {
	return domainsLogic.CreateDomain(ctx, req)
}
