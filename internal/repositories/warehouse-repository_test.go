package repositories

import (
	"DeliveryClub/internal/models"
	"context"
	"testing"
)

func TestWarehouseRepository(t *testing.T) {
	Warehouse1 := models.Warehouse{ID: 1, Location: "Moscow", Email: "example@email.com"}
	Warehouse2 := models.Warehouse{ID: 2, Location: "Smolensk", Email: "example1@email.com"}

	warehouses := []models.Warehouse{Warehouse1, Warehouse2}

	ctx := context.Background()
	repo := NewWarehouseRepository(warehouses)

	// Test Get method
	result, err := repo.Get(ctx)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if len(result) != 2 {
		t.Errorf("expected 2 warehouses but got %d", len(result))
	}

	// Test AddRange method
	newWarehouses := []models.Warehouse{
		{ID: 3, Location: "Warehouse C", Email: "c@example.com"},
	}
	err = repo.AddRange(ctx, newWarehouses)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}

	result, err = repo.Get(ctx)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if len(result) != 3 {
		t.Errorf("expected 3 warehouses but got %d", len(result))
	}

	// Test UpdateRange method
	updatedWarehouses := []models.Warehouse{
		{ID: 2, Location: "Smolensk Updated", Email: "updated@example.com"},
	}
	err = repo.UpdateRange(ctx, updatedWarehouses)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}

	result, err = repo.Get(ctx)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	if result[1].Location != "Smolensk Updated" {
		t.Errorf("expected 'Smolensk Updated' but got %s", result[1].Location)
	}
	if result[1].Email != "updated@example.com" {
		t.Errorf("expected 'updated@example.com' but got %s", result[1].Email)
	}
}
