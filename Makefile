# Makefile для создания миграций

# Переменные которые будут использоваться в наших командах (Таргетах)
DB_DSN := "postgres://user:pass@localhost:5432/db?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN) -verbose

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# Принудительное указание версии
migrate-force:
	migrate -path ./migrations -database $(DB_DSN) force 20241102162031

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

lint:
	golangci-lint run --out-format=colored-line-number

run:
	go run cmd/app/main.go