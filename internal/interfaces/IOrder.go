package interfaces

import (
	"context"
	"ecommerce-order/external"
	"ecommerce-order/internal/models"

	"github.com/labstack/echo/v4"
)

type IOrderRepo interface {
	InsertNewOrder(ctx context.Context, order *models.Order) error
	UpdateStatusOrder(ctx context.Context, orderID int, status string) error
}

type IOrderService interface {
	CreateOrder(ctx context.Context, profile external.Profile, req *models.Order) (*models.Order, error)
}

type IOrderAPI interface {
	CreateOrder(e echo.Context) error
}
