package services_test

import (
	"DeliveryClub/internal/models"
	"DeliveryClub/internal/services"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStoreRepository struct {
	mock.Mock
}

func (m *MockStoreRepository) Get(ctx context.Context) ([]models.Store, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Store), args.Error(1)
}

func (m *MockStoreRepository) AddRange(ctx context.Context, items []models.Store) error {
	args := m.Called(ctx, items)
	return args.Error(0)
}

func (m *MockStoreRepository) UpdateRange(ctx context.Context, items []models.Store) error {
	args := m.Called(ctx, items)
	return args.Error(0)
}

func TestStoreService_GetStores(t *testing.T) {
	mockRepo := new(MockStoreRepository)
	service := services.NewStoreService(mockRepo)
	ctx := context.Background()

	expected := []models.Store{
		{ID: 1, Location: "Location 1"},
	}

	mockRepo.On("Get", ctx).Return(expected, nil)

	result, err := service.GetStores(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	mockRepo.AssertExpectations(t)
}

func TestStoreService_RegisterOrUpdateStores(t *testing.T) {
	mockRepo := new(MockStoreRepository)
	service := services.NewStoreService(mockRepo)
	ctx := context.Background()

	stores := []models.Store{
		{ID: 1, Location: "Location 1"},
		{ID: 2, Location: "Location 2"},
	}

	mockRepo.On("AddRange", ctx, stores).Return(nil)

	err := service.RegisterOrUpdateStores(ctx, stores)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
