# Golang Microservice Stack

Um projeto de microsserviços em Go para integrar Kafka, MongoDB, Prometheus e Grafana.

## 🚀 Visão Geral

Este projeto fornece um exemplo real de como construir uma arquitetura moderna de microsserviços com Go, envolvendo:

- Produtor e consumidor Kafka usando [Sarama](https://github.com/Shopify/sarama)
- Persistência em MongoDB
- Monitoramento com Prometheus e Grafana
- Docker Compose para orquestração de serviços

É ideal para mostrar habilidades técnicas e foco em observabilidade, sem ser apenas um repositório genérico.

## 🧱 Estrutura do Projeto

├── producer/ # REST API para produzir mensagens no Kafka
│ └── producer.go
├── worker/ # Worker consumindo do Kafka e salvando no MongoDB
│ └── worker.go
├── docker-compose.yml # Kafka, Zookeeper, MongoDB, Prometheus e Grafana
├── go.mod
└── README.md

## 🛠️ Tecnologias Utilizadas

- **Go** – back-end principal (producer + consumer)
- **Apache Kafka** – sistema de mensagens via Sarama :contentReference[oaicite:1]{index=1}
- **MongoDB** – banco de dados NoSQL
- **Prometheus** – coleta métricas
- **Grafana** – visualização e dashboards
- **Docker Compose** – orquestração local

## ⚙️ Pré-requisitos

- Docker & Docker Compose
- Go versão ≥ 1.18

## 🔧 Instalação e Execução

1. Clone o repositório  
   ```bash
   git clone https://github.com/verofreitt/golang_microservice_stack.git
   cd golang_microservice_stack

2. Inicie todos os serviços
    ```bash
    docker-compose up -d

3. No terminal 1: rode o producer
    ```bash
    go run producer/producer.go

4. No terminal 2: rode o worker
    ```bash
    go run worker/worker.go

## 📬 Testando o sistema
1. Para publicar uma mensagem via API REST (proxy para Kafka):
    ```bash
    curl -X POST localhost:3000/api/v1/messages \
    -H "Content-Type: application/json" \
    -d '{ "text": "mensagem de exemplo" }'

O worker irá consumir do Kafka, salvar no MongoDB e você verá os logs correspondentes.

## 📊 Monitoramento
Prometheus estará acessível em http://localhost:9090

Grafana disponível em http://localhost:3000 (usuário padrão: admin/admin)

dashboards pré-configurados para métricas do Kafka, consumer lag, latência, uso de memória, etc.

🧩 Possíveis extensões

- Autenticação via JWT

- Dockerfile para builder image

- Deploy em Kubernetes

- Testes unitários e de integração

📝 Licença
Licenciado sob a MIT License
