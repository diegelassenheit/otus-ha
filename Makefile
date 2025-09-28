
PROJECT_NAME=social-network
COMPOSE_DEV=./docker-compose.yaml
DATABASE_URL="postgres://postgres:example@localhost:5435/app?sslmode=disable"

setup:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

up:
	docker compose -p ${PROJECT_NAME} -f $(COMPOSE_DEV) up --build -d

down:
	docker compose -p ${PROJECT_NAME} -f $(COMPOSE_DEV) down

run:
	docker compose -p ${PROJECT_NAME} -f ${COMPOSE_DEV} up app -d

test:
	docker compose -p ${PROJECT_NAME} -f ${COMPOSE_DEV} exec app sh -c "go test ./internal/..."

migrate:
	docker compose -p ${PROJECT_NAME} -f ${COMPOSE_DEV} run --rm migrator
