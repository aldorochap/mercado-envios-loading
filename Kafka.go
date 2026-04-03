package main

import (
	"context"
	"encoding/json"
	"time" // Adicionamos o time aqui
	"github.com/segmentio/kafka-go"
)

// PublishScanEvent envia a confirmação de carregamento para o Kafka
func PublishScanEvent(packageID, vehicleID string) error {
	writer := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "shipment.tracking.events",
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	event := ScanEvent{
		PackageID: packageID,
		VehicleID: vehicleID,
		Status:    "OUT_FOR_DELIVERY",
		Timestamp: time.Now(),
	}

	payload, _ := json.Marshal(event)

	return writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(packageID),
		Value: payload,
	})
}