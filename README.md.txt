# 🚚 Envios - Loading Check-in API

Este microserviço foi desenvolvido para otimizar o processo de carregamento de veículos nas docas, garantindo a **rastreabilidade ponta a ponta** e a **integridade da rota**.

## Diferenciais Técnicos (Visão Project Leader)
- **Arquitetura Orientada a Eventos:** Utiliza Kafka para propagar atualizações de rastreio sem bloquear a operação física.
- **Alta Performance:** Desenvolvido em **Go** para garantir baixíssima latência nos bipes dos coletores.
- **Escalabilidade:** Estrutura modular (Handler, Service, Repository) preparada para crescimento e substituição de provedores de dados.

## 🛠️ Tecnologias
- Go (Golang)
- Kafka (Mensageria assíncrona)
- REST API

## 📋 Como Executar
1. Inicialize o projeto: `go mod init mercado-envios-loading`
2. Instale as dependências: `go get github.com/segmentio/kafka-go`
3. Execute a aplicação: `go run .`

## 🛡️ Regras de Negócio Implementadas
- Validação de `PackageID` vs `RouteID` do veículo.
- Disparo de evento de rastreio (`OUT_FOR_DELIVERY`) com timestamp preciso.