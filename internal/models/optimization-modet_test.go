package models

import (
	"testing"
)

func TestCalculateTransportCost(t *testing.T) {
	transport := Transport{ID: 1, Type: "Truck", Capacity: 1000, Speed: 60.0}
	distance := 100.0
	expectedCost := 100.0 / 60.0 * 1000

	cost := CalculateTransportCost(transport, distance)
	if cost != expectedCost {
		t.Errorf("expected %f but got %f", expectedCost, cost)
	}
}

func TestBuildAndSolveModel(t *testing.T) {
	warehouses := []Warehouse{
		{ID: 1, Location: "Warehouse A", Email: "a@example.com"},
		{ID: 2, Location: "Warehouse B", Email: "b@example.com"},
	}

	stores := []Store{
		{ID: 1, Location: "Store A"},
		{ID: 2, Location: "Store B"},
	}

	transports := []Transport{
		{ID: 1, Type: "Truck", Capacity: 1000, Speed: 60.0},
		{ID: 2, Type: "Van", Capacity: 500, Speed: 80.0},
	}

	goods := []Goods{
		{ID: 1, Name: "Product A", Quantity: 100},
		{ID: 2, Name: "Product B", Quantity: 200},
	}

	om := OptimizationModel{
		ModelID:            1,
		OptimizationMethod: "simple",
		Transport:          transports,
	}

	om.BuildModel(warehouses, stores, goods, transports)
	om.AnalyzeResults()

	if len(om.Routes) == 0 {
		t.Errorf("expected routes to be built but got none")
	}
}
