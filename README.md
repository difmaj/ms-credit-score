
├── cmd/
│   ├── score-hub/
│   │   └── main.go
│   ├── ms-admin/
│   │   └── main.go
│   └── workers/
│       └── score-worker/
│           └── main.go
├── internal/
│   ├── domain/
│   │   ├── assets.go
│   │   ├── debts.go
│   │   └── score.go
│   ├── queue/
│   │   └── rabbitmq.go
│   └── repository/
│       └── repository.go
├── pkg/
│   ├── auth/
│   │   └── middleware.go
│   ├── logger/
│   │   └── logger.go
│   └── utils/
│       └── utils.go
├── migrations/
│   ├── up/
│   └── down/
├── seeds/
│   └── dev/
├── docker/
│   ├── score-hub.Dockerfile
│   ├── ms-admin.Dockerfile
│   └── score-worker.Dockerfile
├── Makefile
├── go.mod
├── go.sum
└── README.md
