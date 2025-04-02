package models

type PaymentInitiatePayload struct {
	OrderID    int     `json:"order_id"`
	TotalPrice float64 `json:"total_price"`
}

type RefundPayload struct {
	OrderID int `json:"order_id"`
	AdminID int `json:"admin_id"`
}
