package repositories_test

import (
	"DeliveryClub/internal/models"
	"DeliveryClub/internal/repositories"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWarehouseRepository_Get(t *testing.T) {
	warehouses := []models.Warehouse{
		{ID: 1, Location: "Location 1", Email: "email1@example.com"},
		{ID: 2, Location: "Location 2", Email: "email2@example.com"},
	}
	repo := repositories.NewWarehouseRepository(warehouses)

	ctx := context.Background()
	result, err := repo.Get(ctx)

	assert.NoError(t, err)
	assert.Equal(t, warehouses, result)
}

func TestWarehouseRepository_AddRange(t *testing.T) {
	warehouses := []models.Warehouse{
		{ID: 1, Location: "Location 1", Email: "email1@example.com"},
	}
	repo := repositories.NewWarehouseRepository(warehouses)

	newWarehouses := []models.Warehouse{
		{ID: 2, Location: "Location 2", Email: "email2@example.com"},
		{ID: 3, Location: "Location 3", Email: "email3@example.com"},
	}

	ctx := context.Background()
	err := repo.AddRange(ctx, newWarehouses)
	assert.NoError(t, err)

	expected := append(warehouses, newWarehouses...)
	result, err := repo.Get(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestWarehouseRepository_UpdateRange(t *testing.T) {
	initialWarehouses := []models.Warehouse{
		{ID: 1, Location: "Location 1", Email: "email1@example.com"},
		{ID: 2, Location: "Location 2", Email: "email2@example.com"},
	}
	repo := repositories.NewWarehouseRepository(initialWarehouses)

	updatedWarehouses := []models.Warehouse{
		{ID: 1, Location: "Updated Location 1", Email: "updatedemail1@example.com"},
		{ID: 2, Location: "Updated Location 2", Email: "updatedemail2@example.com"},
	}

	ctx := context.Background()
	err := repo.UpdateRange(ctx, updatedWarehouses)
	assert.NoError(t, err)

	result, err := repo.Get(ctx)
	assert.NoError(t, err)
	assert.Equal(t, updatedWarehouses, result)
}
