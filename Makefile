export PROJECT_ROOT := $(shell pwd)

include .env
export

up: ## Start PostgreSQL
	@docker compose up -d todoapp-postgres

down: ## Stop and remove Compose containers, keeping database files
	@docker compose down

migrate: ## Run migrations (ACTION="up 1", ACTION="down 2", etc.)
	@docker compose run --rm todoapp-postgres-migrate \
	  -path /migrations \
	  -database="postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@todoapp-postgres:5432/$(POSTGRES_DB)?sslmode=disable&search_path=public" \
	  $(ACTION)

migrate-create: ## Create migration files (NAME=create_something)
	@if [ -n "$(NAME)" ]; then \
	  docker compose run --rm todoapp-postgres-migrate \
	    create \
	    -ext sql \
	    -dir /migrations \
	    -seq "$(NAME)"; \
	else \
	  echo 'Usage: make migrate-create NAME=create_something'; \
	  exit 1; \
	fi

clean: ## Stop PostgreSQL and delete only its database files
	@read -p 'Permanently delete database files in out/pgdata? [y/N] ' answer; \
	  if [ "$$answer" = y ]; then \
	    docker compose down && \
	    docker run --rm --user root \
	      -v "$(PROJECT_ROOT)/out:/workspace" \
	      --entrypoint /bin/rm postgres:18.1-bookworm \
	      -rf /workspace/pgdata && \
	    echo "Data has been removed"; \
	  else \
	    echo 'Cleanup cancelled.'; \
	  fi

port-forward:
	@docker compose up -d port-forwarder

port-close:
	@docker compose down port-forwarder



todoapp-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/api/main.go

logs-cleanup:
	@read -p 'Permanently delete log files in out/logs? [y/N] ' answer; \
	  if [ "$$answer" = y ]; then \
	    docker compose down && \
	    docker run --rm --user root \
	      -v "$(PROJECT_ROOT)/out:/workspace" \
	      --entrypoint /bin/rm postgres:18.1-bookworm \
	      -rf /workspace/logs && \
	    echo "logs has been removed"; \
	  else \
	    echo 'Cleanup cancelled.'; \
	  fi

todoapp-deploy:
	@docker compose up -d --build todoapp

todoapp-undeploy:
	@docker compose down todoapp
