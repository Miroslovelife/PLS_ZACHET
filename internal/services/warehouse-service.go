package services

import (
	"DeliveryClub/internal/models"
	"DeliveryClub/internal/repositories"
	"context"
)

type WarehouseService struct {
	repository *repositories.WarehouseRepository
}

func NewWarehouseService(repo *repositories.WarehouseRepository) *WarehouseService {
	return &WarehouseService{repository: repo}
}

func (ws *WarehouseService) GetWarehouses(ctx context.Context) ([]models.Warehouse, error) {
	return ws.repository.Get(ctx)
}

func (ws *WarehouseService) RegisterOrUpdateWarehouses(ctx context.Context, warehouses []models.Warehouse) error {
	return ws.repository.AddRange(ctx, warehouses)
}
