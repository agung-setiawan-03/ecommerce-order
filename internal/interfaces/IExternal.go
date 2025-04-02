package interfaces

import (
	"context"
	"ecommerce-order/external"
)

type IExternal interface {
	GetProfile(ctx context.Context, token string) (external.Profile, error)
}
