package models

type PaymentInitiatePayload struct {
	OrderID    int     `json:"order_id"`
	TotalPrice float64 `json:"total_price"`
}
