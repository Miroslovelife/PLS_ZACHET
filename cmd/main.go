package main

import (
	"DeliveryClub/internal/models"
	"DeliveryClub/internal/repositories"
	services "DeliveryClub/internal/services"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	// Warehouse
	warehouseRepo := repositories.NewWarehouseRepository()
	warehouseService := services.NewWarehouseService(warehouseRepo)

	warehouses, err := warehouseService.GetWarehouses(ctx)
	if err != nil {
		fmt.Println("Error getting warehouses:", err)
		return
	}
	fmt.Println("Warehouses:", warehouses)

	// Store
	storeRepo := repositories.NewStoreRepository()
	storeService := services.NewStoreService(storeRepo)

	stores, err := storeService.GetStores(ctx)
	if err != nil {
		fmt.Println("Error getting stores:", err)
		return
	}
	fmt.Println("Stores:", stores)

	// Mock transports and goods for the demonstration
	transports := []models.Transport{
		{ID: 1, Type: "Truck", Capacity: 1000, Speed: 60.0},
		{ID: 2, Type: "Van", Capacity: 500, Speed: 80.0},
	}

	goods := []models.Goods{
		{ID: 1, Name: "Product A", Quantity: 100},
		{ID: 2, Name: "Product B", Quantity: 200},
	}

	// Optimization
	models.OptimizeTransport(warehouses, stores, goods, transports)
}
