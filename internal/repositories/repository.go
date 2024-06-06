package repositories

import (
	"DeliveryClub/internal/models"
	"context"
)

type WarehouseRepository struct {
	warehouses []models.Warehouse
}

func NewWarehouseRepository(warehouses []models.Warehouse) *WarehouseRepository {
	return &WarehouseRepository{
		warehouses: warehouses,
	}
}

func (wr *WarehouseRepository) Get(ctx context.Context) ([]models.Warehouse, error) {
	return wr.warehouses, nil
}

func (wr *WarehouseRepository) AddRange(ctx context.Context, items []models.Warehouse) error {
	wr.warehouses = append(wr.warehouses, items...)
	return nil
}

func (wr *WarehouseRepository) UpdateRange(ctx context.Context, items []models.Warehouse) error {
	for _, newItem := range items {
		for i, oldItem := range wr.warehouses {
			if oldItem.ID == newItem.ID {
				wr.warehouses[i] = newItem
				break
			}
		}
	}
	return nil
}
