package services

import (
	"DeliveryClub/internal/models"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockStoreRepository is a mock implementation of the StoreRepositoryInterface
type MockStoreRepository struct {
	mock.Mock
}

func (m *MockStoreRepository) Get(ctx context.Context) ([]models.Store, error) {
	args := m.Called(ctx)
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
	service := NewStoreService(mockRepo)
	ctx := context.Background()

	expectedStores := []models.Store{
		{ID: 1, Location: "Store 1"},
		{ID: 2, Location: "Store 2"},
	}

	mockRepo.On("Get", ctx).Return(expectedStores, nil)

	stores, err := service.GetStores(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedStores, stores)

	mockRepo.AssertExpectations(t)
}

func TestStoreService_GetStores_Error(t *testing.T) {
	mockRepo := new(MockStoreRepository)
	service := NewStoreService(mockRepo)
	ctx := context.Background()

	expectedError := errors.New("database error")

	mockRepo.On("Get", ctx).Return(nil, expectedError)

	stores, err := service.GetStores(ctx)

	assert.Error(t, err)
	assert.Nil(t, stores)
	assert.Equal(t, expectedError, err)

	mockRepo.AssertExpectations(t)
}

func TestStoreService_RegisterOrUpdateStores(t *testing.T) {
	mockRepo := new(MockStoreRepository)
	service := NewStoreService(mockRepo)
	ctx := context.Background()

	stores := []models.Store{
		{ID: 1, Location: "Store 1"},
		{ID: 2, Location: "Store 2"},
	}

	mockRepo.On("AddRange", ctx, stores).Return(nil)

	err := service.RegisterOrUpdateStores(ctx, stores)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestStoreService_RegisterOrUpdateStores_Error(t *testing.T) {
	mockRepo := new(MockStoreRepository)
	service := NewStoreService(mockRepo)
	ctx := context.Background()

	stores := []models.Store{
		{ID: 1, Location: "Store 1"},
		{ID: 2, Location: "Store 2"},
	}

	expectedError := errors.New("database error")

	mockRepo.On("AddRange", ctx, stores).Return(expectedError)

	err := service.RegisterOrUpdateStores(ctx, stores)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)

	mockRepo.AssertExpectations(t)
}
