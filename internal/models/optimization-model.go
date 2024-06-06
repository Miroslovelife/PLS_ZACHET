package models

import (
	"fmt"
)

type OptimizationModel struct {
	ModelID            int
	OptimizationMethod string
	Transport          []Transport
	Routes             map[string]float64
}

func (om *OptimizationModel) BuildModel(warehouses []Warehouse, stores []Store, goods []Goods, transports []Transport) {
	om.Routes = make(map[string]float64)
	for _, warehouse := range warehouses {
		for _, store := range stores {
			distance := calculateDistance(warehouse.Location, store.Location)
			for _, transport := range transports {
				cost := CalculateTransportCost(transport, distance)
				routeKey := fmt.Sprintf("%s-%s-%s", warehouse.Location, store.Location, transport.Type)
				om.Routes[routeKey] = cost
			}
		}
	}
}

func (om *OptimizationModel) SolveModel() {
	minCost := float64(1<<63 - 1)
	bestRoute := ""
	for route, cost := range om.Routes {
		if cost < minCost {
			minCost = cost
			bestRoute = route
		}
	}
	fmt.Printf("Best Route: %s with cost: %.2f\n", bestRoute, minCost)
}

func (om *OptimizationModel) AnalyzeResults() {
	// Simple analysis by printing the best route and its cost
	om.SolveModel()
}

func CalculateTransportCost(transport Transport, distance float64) float64 {
	return distance / transport.Speed * float64(transport.Capacity)
}

func OptimizeTransport(warehouses []Warehouse, stores []Store, goods []Goods, transports []Transport) {
	om := OptimizationModel{
		ModelID:            1,
		OptimizationMethod: "simple",
		Transport:          transports,
	}
	om.BuildModel(warehouses, stores, goods, transports)
	om.AnalyzeResults()
}

func calculateDistance(location1, location2 string) float64 {
	// Simple mock of distance calculation, should be replaced by real geo distance calculation
	return float64(len(location1) + len(location2))
}
