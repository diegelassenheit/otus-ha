Social Network (Go) — Quick Start

Prerequisites
1) Docker and Docker Compose (v2)
2) Make

First run (one command)
1) make check
   - Installs migrate CLI (local), starts DB, runs migrations, then starts app and Adminer.

Common commands
1) make up      # start all services in background (db, app, adminer)
2) make down    # stop and remove containers
3) make migrate # run DB migrations (requires DB healthy)

Services & ports
- App API: http://localhost:8080
- Postgres: localhost:5435 (db:5432 inside compose)
- Adminer: http://localhost:8081 

Environment
- App connects to DB via DATABASE_URL in docker-compose.

API usage
- OpenAPI spec: api/openapi.json 
- Minimal Postman collection: postman_collection.json
