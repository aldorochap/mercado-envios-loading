package main

import "time"

// Shipment representa o pacote físico
type Shipment struct {
	ID      string `json:"id"`
	RouteID string `json:"route_id"`
	Status  string `json:"status"`
}

// ScanRequest é o dado que chega do coletor do motorista
type ScanRequest struct {
	PackageID string `json:"package_id"`
	VehicleID string `json:"vehicle_id"`
}

// ScanEvent é o rastro oficial para o Kafka (Rastreabilidade)
type ScanEvent struct {
	PackageID string    `json:"package_id"`
	VehicleID string    `json:"vehicle_id"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}