package main

import "fmt"

// ValidateLoading executa a regra de negócio e dispara o evento de rastreio
func ValidateLoading(req ScanRequest) error {
	// 1. Busca o pacote no repositório
	shipment, err := GetShipmentByID(req.PackageID)
	if err != nil {
		return err
	}

	// 2. Valida a Rota (Regra Crítica)
	// Simulamos que o veículo atual está designado para a "ROTA-SUL"
	currentVehicleRoute := "ROTA-SUL" 
	if shipment.RouteID != currentVehicleRoute {
		return fmt.Errorf("ROUTE_MISMATCH: Pacote deveria ir para %s", shipment.RouteID)
	}

	// 3. Se tudo ok, dispara Rastreabilidade (Kafka) de forma assíncrona
	go PublishScanEvent(req.PackageID, req.VehicleID)

	return nil
}