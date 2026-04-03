package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a rota e quem vai cuidar dela (o Handler)
	http.HandleFunc("/v1/scan", ScanHandler)

	fmt.Println("🚚 Mercado Envios - Microserviço de Carregamento Ativo [Porta 8080]")
	
	// Inicia o servidor
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Erro ao subir servidor: %s\n", err)
	}
}