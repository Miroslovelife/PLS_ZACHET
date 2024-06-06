package services

import (
	"DeliveryClub/internal/models"
	"DeliveryClub/internal/repositories"
	"context"
)

type StoreService struct {
	repository repositories.StoreRepositoryInterface
}

func NewStoreService(repo repositories.StoreRepositoryInterface) *StoreService {
	return &StoreService{repository: repo}
}

func (ss *StoreService) GetStores(ctx context.Context) ([]models.Store, error) {
	return ss.repository.Get(ctx)
}

func (ss *StoreService) RegisterOrUpdateStores(ctx context.Context, stores []models.Store) error {
	return ss.repository.AddRange(ctx, stores)
}
