package constants

const (
	Success             = "Success"
	ErrFailedBadRequest = "Data tidak sesuai"
	ErrServerError      = "Terjadi kesalahan pada server"
)

const (
	OrderStatusSuccess = "SUCCESS"
	OrderStatusPending = "PENDING"
	OrderStatusFailed  = "FAILED"
	OrderStatusRefund  = "REFUND"
)

var MappingOrderStatus = map[string]bool{
	OrderStatusSuccess: true,
	OrderStatusPending: true,
	OrderStatusFailed:  true,
	OrderStatusRefund:  true,
}

var MappingFlowOrderStatus = map[string][]string{
	OrderStatusPending: {OrderStatusFailed, OrderStatusSuccess},
	OrderStatusSuccess: {OrderStatusRefund},
}
