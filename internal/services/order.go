package services

import (
	"context"
	"ecommerce-order/constants"
	"ecommerce-order/external"
	"ecommerce-order/helpers"
	"ecommerce-order/internal/interfaces"
	"ecommerce-order/internal/models"
	"encoding/json"

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
