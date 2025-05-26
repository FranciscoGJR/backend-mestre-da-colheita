include .env

APP_CONTAINER=$(APP_CONTAINER_NAME)
DB_CONTAINER=$(DB_HOST)
NETWORK=rede-go

all: shell

shell:
	docker exec -it $(APP_CONTAINER) sh

logs:
	docker logs $(APP_CONTAINER)
docker-build:

	docker build -t mestre-da-colheita .

stop:
	-docker stop $(APP_CONTAINER)
	-docker stop $(DB_CONTAINER)

start:
	docker start $(APP_CONTAINER)
	docker start $(DB_CONTAINER)

rm:
	-docker rm $(APP_CONTAINER)
	-docker rm $(DB_CONTAINER)

network-create:
	docker network create $(NETWORK) || true

db-run:
	docker run -d \
		--name $(DB_CONTAINER) \
		--network $(NETWORK) \
		-e POSTGRES_USER=$(DB_USER) \
		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		-e POSTGRES_DB=$(DB_NAME) \
		-p 5432:5432 \
		postgres:alpine

docker-run: network-create
	docker run -d \
		--name $(APP_CONTAINER) \
		--network $(NETWORK) \
		--env-file .env \
		-p 8080:8080 \
		$(APP_CONTAINER)

up: db-run docker-run

help:
	@echo "Comandos disponíveis:"
	@echo "  make shell         - Entra no shell do container"
	@echo "  make logs          - Exibe os logs do container"
	@echo "  make stop          - Para os containers"
	@echo "  make start         - Inicia os containers"
	@echo "  make rm            - Remove os containers"
	@echo "  make docker-build  - Builda a imagem do app"
	@echo "  make db-run        - Sobe o banco Postgres"
	@echo "  make docker-run    - Sobe o backend Go"
	@echo "  make up            - Sobe tudo (banco + backend)"
	@echo "  make help          - Mostra esta ajuda"
