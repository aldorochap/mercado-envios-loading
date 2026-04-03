package main

import "errors"

// GetShipmentByID simula a busca no banco de dados do Mercado Envios
func GetShipmentByID(id string) (Shipment, error) {
	// SIMULAÇÃO: No mundo real, aqui teria um SELECT * FROM shipments
	// Vamos fingir que o pacote "ABC-123" pertence à rota "ROTA-SUL"
	if id == "ABC-123" {
		return Shipment{ID: "ABC-123", RouteID: "ROTA-SUL", Status: "WAREHOUSE"}, nil
	}
	
	return Shipment{}, errors.New("SHIPMENT_NOT_FOUND: Pacote não registrado")
}