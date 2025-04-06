package models

import (
	"time"

	"github.com/go-playground/validator"
)

type Order struct {
	ID         int         `json:"id"`
	UserID     int         `json:"user_id"`
	TotalPrice float64     `json:"total_price" gorm:"column:total_price;type:decimal(10,2)" validate:"required"`
	Status     string      `json:"status" gorm:"column:status;type:varchar(10)"`
	CreatedAt  time.Time   `json:"-"`
	UpdatedAt  time.Time   `json:"-"`
	OrderItem  []OrderItem `json:"items"`
}

func (*Order) TableName() string {
	return "orders"
}

func (l Order) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type OrderItem struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"order_id"`
	ProductID int       `json:"product_id" validate:"required"`
	VariantID int       `json:"variant_id" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required"`
	Price     float64   `json:"price" gorm:"column:price;type:decimal(10,2)" validate:"required"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (*OrderItem) TableName() string {
	return "order_items"
}

func (l OrderItem) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type OrderStatusRequest struct {
	Status string `json:"status" validate:"required"`
}

func (l OrderStatusRequest) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
