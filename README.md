# Project Structure

### Architecture
- #### Microservices

### Branching strategy
- #### Multi repositories

### Coding convention
- #### Clean architecture

### Project tree

```
project/
├── cmd/                           # Entry point for each Service
│   ├── order/
│   │   └── main.go
│   ├── payment/
│   │   └── main.go
├── internal/
│   ├── repository/                # General repository functions; Example: Exec() for database query mechanisms
│   ├── delivery/                  # General delivery functions; Example: Consume() for message consumption mechanisms
│   ├── constants/                 # Application constants
│   ├── errors/                    # Custom error types and handling
│   ├── utils/                     # Utility functions used across the service
│   └── models/                    # Shared models used across services
│       ├── order/                 # Example: Order models
├── libs/
│   ├── logger/                    # Logging library
│   ├── utils/                     # Utility functions shared across services
│   ├── config/                    # Configuration handling
│   ├── auth/                      # Authorization library (e.g., JWT, OAuth)
│   ├── middleware/                # Middleware (e.g., logging, rate limiting)
│   └── engine/                    # Wrappers around DB engines, HTTP frameworks, gRPC
│       ├── db/
│       ├── http/
│       └── grpc/
├── services/
│   ├── order/
│   │   ├── internal/
│   │   │   ├── delivery/          # Service-specific delivery functions; Example http, rpc, event message
│   │   │   ├── usecase/           # Business logic
│   │   │   ├── repository/        # Data access layer
│   │   │   └── domain/            # Domain logic and models
│   │   │       ├── order.go       # Domain model
│   │   │       └── mock/          # Mock data
│   │   └── api/
│   │       └── proto/
│   │           └── order.proto    # Protocol Buffers file
│   └── payment/
│       ├── internal/
│       │   ├── delivery/
│       │   ├── usecase/
│       │   ├── repository/
│       │   └── domain/
│       │       ├── payment.go
│       │       └── mock/
│       └── api/
│           └── proto/
│               └── payment.proto
├── go.mod                         # Single go.mod for the entire project
└── go.sum                         # Single go.sum for the entire project

```
