# Estilos de organização de projetos Go


## 1. Flat Layout – projetos pequenos/CLI
> O Minimalista

meu-projeto/
│── go.mod
│── main.go
│── user.go
│── service.go
│── handler.go

-> Usado em ferramentas simples, utilitários ou CLIs.
-> Tudo fica na raiz, sem subpacotes.
-> Fácil de navegar, mas pouco escalável.

## 2. Standard Go Project Layout
> O Equilibrado

meu-projeto/
│── go.mod
│── main.go
│
├── api/
│   ├── handler.go
│   └── router.go
│
├── models/
│   └── user.go
│
└── services/
    └── user_service.go

-> É o mais comum e recomendado pela comunidade.
-> Usa subdiretórios para separar responsabilidades.
-> Escalável, modular, facilita testes e reutilização.

## 3. Domain/Feature First (DDD)
> O Escalável

meu-projeto/
│── go.mod
│── main.go
│
├── users/
│   ├── user.go
│   ├── user_service.go
│   └── user_handler.go
│
├── orders/
│   ├── order.go
│   ├── order_service.go
│   └── order_handler.go

-> Estrutura organizada por domínios ou features.
-> Cada domínio tem seu próprio pacote com modelos, serviços e handlers.
-> Ideal para projetos grandes, orientados a domínio, seguindo DDD.

## 4. Hexagonal / Clean Architecture
> O Robusto

meu-projeto/
│── go.mod
│── main.go
│
├── domain/
│   └── user.go
│
├── application/
│   └── user_service.go
│
├── infrastructure/
│   ├── http/
│   │   └── user_handler.go
│   └── db/
│       └── user_repository.go

-> Separação clara entre camadas: domínio, aplicação e infraestrutura.
-> Facilita testes e independência de frameworks.
-> Mantém regras de negócio independentes de detalhes técnicos (como banco ou HTTP).

# Mais Info
https://github.com/golang-standards/project-layout