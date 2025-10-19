
PROJECT_NAME=social-network
COMPOSE_DEV=./docker-compose.yaml
DATABASE_URL="postgres://postgres:example@localhost:5435/app?sslmode=disable"

setup:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

up:
	# Start all services (db, app, adminer) in the background
	docker compose -p ${PROJECT_NAME} -f ${COMPOSE_DEV} up -d

down:
	# Stop all services and remove containers
	docker compose -p ${PROJECT_NAME} -f ${COMPOSE_DEV} down

migrate:
	# Run migrations against a healthy db (one-off container)
	docker compose -p ${PROJECT_NAME} -f ${COMPOSE_DEV} run --rm migrator

db-up:
	# Ensure database is up and healthy before migrations
	docker compose -p ${PROJECT_NAME} -f ${COMPOSE_DEV} up -d db

app-up:
	# Start application and optional tools after migrations
	docker compose -p ${PROJECT_NAME} -f ${COMPOSE_DEV} up -d app adminer

check:
	# Full check: setup tools, start db, run migrations, then start app
	make setup && make db-up && make migrate && make app-up