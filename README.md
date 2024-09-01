# Project Structure

### Architecture
- #### Microservices

### Branching strategy
- #### multi repositories

### Coding convention
- #### Clean architecture

### Project tree

```
project/
├── cmd/
│   ├── order_service/
│   │   └── main.go              # Entry point for the Order Service
├── internal/
│   ├── repository/              # General repository functions; Example: Exec() for database query mechanisms
│   ├── delivery/                # General delivery functions; Example: Consume() for message consumption mechanisms
│   ├── constants/               # Application constants
│   ├── errors/                  # Custom error types and handling
│   ├── utils/                   # Utility functions used across the service
│   └── models/                  # Shared models used across services
│       ├── order/               # Example: Order models
├── libs/
│   ├── logger/                  # Logging library
│   ├── utils/                   # Utility functions shared across services
│   ├── config/                  # Configuration handling
│   ├── auth/                    # Authorization library (e.g., JWT, OAuth)
│   ├── middleware/              # Middleware (e.g., logging, rate limiting)
│   └── engine/                  # Wrappers around DB engines, HTTP frameworks, gRPC
│       ├── db/
│       ├── http/
│       └── grpc/
├── services/
│   ├── order/
│   │   ├── internal/
│   │   │   ├── delivery/        # Service-specific delivery functions; Example http, rpc, event message
│   │   │   ├── usecase/         # Business logic for order management
│   │   │   ├── repository/      # Data access layer for order management
│   │   │   └── domain/          # Domain logic and models for order management
│   │   │       ├── order.go     # Domain model for orders
│   │   │       └── mock/        # Mock data or interfaces for testing
│   │   └── api/
│   │       └── proto/
│   │           └── order.proto  # Protocol Buffers file for order service
├── go.mod                       # Single go.mod for the entire project
└── go.sum                       # Single go.sum for the entire project

```