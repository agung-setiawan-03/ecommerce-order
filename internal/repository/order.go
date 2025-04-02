package repository

import (
	"context"
	"ecommerce-order/internal/models"

	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func (r *OrderRepo) InsertNewOrder(ctx context.Context, order *models.Order) error {
	return r.DB.Transaction(
		func(tx *gorm.DB) error {
			err := tx.Create(order).Error
			if err != nil {
				return err
			}

			for i, orderItem := range order.OrderItems {
				orderItem.OrderID = order.ID
				err := tx.Create(&orderItem).Error
				if err != nil {
					return err
				}
				order.OrderItems[i].ID = orderItem.ID
				order.OrderItems[i].OrderID = order.ID
			}

			return nil
		},
 	 )
}


func (r *OrderRepo) UpdateStatusOrder(ctx context.Context, orderID int, status string) error {
	return r.DB.Exec("UPDATE orders SET status = ? WHERE id = ?", status, orderID).Error
}
