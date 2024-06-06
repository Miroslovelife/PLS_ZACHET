package repositories

import (
	"DeliveryClub/internal/models"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreRepository_Get(t *testing.T) {
	stores := []models.Store{
		{ID: 1, Location: "Store 1"},
		{ID: 2, Location: "Store 2"},
	}

	repo := NewStoreRepository(stores)
	ctx := context.Background()

	result, err := repo.Get(ctx)

	assert.NoError(t, err)
	assert.Equal(t, stores, result)
}

func TestStoreRepository_AddRange(t *testing.T) {
	initialStores := []models.Store{
		{ID: 1, Location: "Store 1"},
	}

	newStores := []models.Store{
		{ID: 2, Location: "Store 2"},
		{ID: 3, Location: "Store 3"},
	}

	expectedStores := append(initialStores, newStores...)
	repo := NewStoreRepository(initialStores)
	ctx := context.Background()

	err := repo.AddRange(ctx, newStores)

	assert.NoError(t, err)

	result, _ := repo.Get(ctx)
	assert.Equal(t, expectedStores, result)
}

func TestStoreRepository_UpdateRange(t *testing.T) {
	initialStores := []models.Store{
		{ID: 1, Location: "Store 1"},
		{ID: 2, Location: "Store 2"},
	}

	updatedStores := []models.Store{
		{ID: 1, Location: "Updated Store 1"},
		{ID: 2, Location: "Updated Store 2"},
	}

	repo := NewStoreRepository(initialStores)
	ctx := context.Background()

	err := repo.UpdateRange(ctx, updatedStores)

	assert.NoError(t, err)

	result, _ := repo.Get(ctx)
	assert.Equal(t, updatedStores, result)
}

func TestStoreRepository_UpdateRange_PartialUpdate(t *testing.T) {
	initialStores := []models.Store{
		{ID: 1, Location: "Store 1"},
		{ID: 2, Location: "Store 2"},
	}

	updatedStores := []models.Store{
		{ID: 1, Location: "Updated Store 1"},
	}

	expectedStores := []models.Store{
		{ID: 1, Location: "Updated Store 1"},
		{ID: 2, Location: "Store 2"},
	}

	repo := NewStoreRepository(initialStores)
	ctx := context.Background()

	err := repo.UpdateRange(ctx, updatedStores)

	assert.NoError(t, err)

	result, _ := repo.Get(ctx)
	assert.Equal(t, expectedStores, result)
}
