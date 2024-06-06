package repositories_test

import (
	"DeliveryClub/internal/models"
	"DeliveryClub/internal/repositories"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreRepository_Get(t *testing.T) {
	stores := []models.Store{
		{ID: 1, Location: "Location 1"},
		{ID: 2, Location: "Location 2"},
	}
	repo := repositories.NewStoreRepository(stores)

	ctx := context.Background()
	result, err := repo.Get(ctx)

	assert.NoError(t, err)
	assert.Equal(t, stores, result)
}

func TestStoreRepository_AddRange(t *testing.T) {
	stores := []models.Store{
		{ID: 1, Location: "Location 1"},
	}
	repo := repositories.NewStoreRepository(stores)

	newStores := []models.Store{
		{ID: 2, Location: "Location 2"},
		{ID: 3, Location: "Location 3"},
	}

	ctx := context.Background()
	err := repo.AddRange(ctx, newStores)
	assert.NoError(t, err)

	expected := append(stores, newStores...)
	result, err := repo.Get(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestStoreRepository_UpdateRange(t *testing.T) {
	initialStores := []models.Store{
		{ID: 1, Location: "Location 1"},
		{ID: 2, Location: "Location 2"},
	}
	repo := repositories.NewStoreRepository(initialStores)

	updatedStores := []models.Store{
		{ID: 1, Location: "Updated Location 1"},
		{ID: 2, Location: "Updated Location 2"},
	}

	ctx := context.Background()
	err := repo.UpdateRange(ctx, updatedStores)
	assert.NoError(t, err)

	result, err := repo.Get(ctx)
	assert.NoError(t, err)
	assert.Equal(t, updatedStores, result)
}
