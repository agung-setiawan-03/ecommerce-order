package services

import (
	"context"
	"ecommerce-order/constants"
	"ecommerce-order/external"
	"ecommerce-order/helpers"
	"ecommerce-order/internal/interfaces"
	"ecommerce-order/internal/models"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type OrderService struct {
	OrderRepo interfaces.IOrderRepo
	External  interfaces.IExternal
}

func (s *OrderService) CreateOrder(ctx context.Context, profile external.Profile, req *models.Order) (*models.Order, error) {

	req.UserID = profile.Data.ID
	req.Status = constants.OrderStatusPending
	err := s.OrderRepo.InsertNewOrder(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert order")
	}

	// Produce new message
	kafkaPayload := models.PaymentInitiatePayload{
		OrderID:    req.ID,
		TotalPrice: req.TotalPrice,
	}
	jsonPayload, err := json.Marshal(kafkaPayload)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal kafka payload")
	}
	kafkaErr := s.External.ProduceKafkaMessage(ctx, jsonPayload)
	if kafkaErr != nil {
		err := s.OrderRepo.UpdateStatusOrder(ctx, req.ID, constants.OrderStatusFailed)
		if err != nil {
			helpers.Logger.Error("failed to update status to failed: ", err)
		}
		return nil, errors.Wrap(kafkaErr, "failed from kafka external")
	}

	return req, nil
}

func (s *OrderService) UpdateOrderStatus(ctx context.Context, orderID int, req models.OrderStatusRequest) error {
	if !constants.MappingOrderStatus[req.Status] {
		return fmt.Errorf("invalid status request: %s", req.Status)
	}

	order, err := s.OrderRepo.GetOrderDetail(ctx, orderID)
	if err != nil {
		return errors.Wrap(err, "failed to get order detail")
	}

	validStatusReq := false
	statusFlow := constants.MappingFlowOrderStatus[order.Status]
	for i := range statusFlow {
		if statusFlow[i] == req.Status {
			validStatusReq = true
		}
	} 

	if !validStatusReq {
		return fmt.Errorf("invalid status flow, current status: %s, new status: %s", order.Status, req.Status)
	}

	return s.OrderRepo.UpdateStatusOrder(ctx, orderID, req.Status)
}
