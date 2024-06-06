package services

import (
	"DeliveryClub/internal/models"
	"context"
)

type WarehouseRepositoryInterface interface {
	Get(ctx context.Context) ([]models.Warehouse, error)
	AddRange(ctx context.Context, items []models.Warehouse) error
}

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
