package main

import (
	"encoding/json"
	"net/http"
)

func ScanHandler(w http.ResponseWriter, r *http.Request) {
	var req ScanRequest
	
	// Transforma o JSON que chegou em uma Struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON Inválido", http.StatusBadRequest)
		return
	}

	// Chama o serviço para validar
	err := ValidateLoading(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Retorno de Sucesso para o motorista
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sucesso: Pacote autorizado para carregamento!"})
}