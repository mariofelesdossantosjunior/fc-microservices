# fc-microservices

🚀 **fc-microservices** é um projeto em **Golang** que demonstra a implementação de uma **arquitetura de microserviços**, com foco em organização de código, boas práticas e escalabilidade.

O objetivo do repositório é servir como base de estudo e também como boilerplate para projetos reais baseados em microserviços.

Todos os conceitos aplicados é com base no curso do Arquitetura baseada em microserviços da FullCycle

---

## 📌 Visão Geral

A arquitetura de microserviços permite dividir uma aplicação grande em serviços menores e independentes, possibilitando:

- Deploy independente de cada serviço
- Escalabilidade individual
- Isolamento de falhas
- Evolução contínua do sistema

Cada serviço possui uma responsabilidade bem definida e se comunica com outros serviços via API (HTTP/REST ou gRPC).

---

## 🧱 Estrutura do Projeto

```bash
fc-microservices/
├── cmd/                 # Entry points (main.go) de cada microserviço
│   ├── service-a/
│   │   └── main.go
│   └── service-b/
│       └── main.go
├── internal/            # Código interno de cada serviço
│   ├── service-a/
│   └── service-b/
├── pkg/                 # Pacotes compartilhados
├── docker-compose.yml   # Orquestração dos serviços
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

> A estrutura segue as convenções recomendadas para projetos Go, separando claramente responsabilidades e evitando acoplamento excessivo.

---

## ⚙️ Tecnologias Utilizadas

- **Go (Golang)**
- **HTTP / REST**
- **Docker**
- **Docker Compose**
- **Arquitetura de Microserviços**

---

## 📦 Pré-requisitos

Antes de executar o projeto, certifique-se de ter instalado:

- Go 1.18 ou superior
- Docker
- Docker Compose

---

## 🚀 Como Executar

### 🔹 Executando localmente (sem Docker)

```bash
# Clone o repositório
git clone https://github.com/mariofelesdossantosjunior/fc-microservices.git

# Acesse a pasta do projeto
cd fc-microservices

# Execute um serviço
go run ./cmd/service-a
```

---

### 🐳 Executando com Docker Compose

```bash
docker compose up --build
```

Esse comando irá subir todos os microserviços definidos no arquivo `docker-compose.yml`.

---

## 🔗 Endpoints (Exemplo)

> Ajuste conforme os serviços implementados no projeto

```http
GET  /health
POST /api/resource
GET  /api/resource/{id}
```

---

## 🧪 Testes

Para executar todos os testes do projeto:

```bash
go test ./...
```

---

## 📖 Boas Práticas Aplicadas

- Separação clara de responsabilidades
- Código desacoplado entre serviços
- Facilidade de manutenção e evolução
- Pronto para CI/CD

---

## 🤝 Contribuição

Contribuições são bem-vindas!

1. Faça um fork do projeto
2. Crie uma branch para sua feature
3. Commit suas alterações
4. Abra um Pull Request

---

## 📄 Licença

Este projeto está sob a licença **MIT**. Sinta-se livre para usar, modificar e distribuir.

---

## ✨ Autor

**Mario Feles dos Santos Junior**
GitHub: [https://github.com/mariofelesdossantosjunior](https://github.com/mariofelesdossantosjunior)

---
