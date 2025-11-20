// 方法：ControllerV1.ListDomains
package support_domains

import (
	"context"

	"kgplatform-backend/api/support_domains/v1"
	domainsLogic "kgplatform-backend/internal/logic/support_domains"
)

func (c *ControllerV1) ListDomains(ctx context.Context, req *v1.ListDomainsReq) (res *v1.ListDomainsRes, err error) {
	return domainsLogic.ListDomains(ctx, req)
}
