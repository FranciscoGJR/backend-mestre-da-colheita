APP_CONTAINER=mestre-da-colheita

all: shell

shell:
	docker exec -it $(APP_CONTAINER) bash

logs:
	docker logs $(APP_CONTAINER)

docker-build:
	docker build -t mestre-da-colheita .

stop:
	docker stop $(APP_CONTAINER)

start:
	docker start $(APP_COTAINER)

rm:
	docker rm $(APP_CONTAINER)

docker-run:
	docker run -p 8080:8080 --network rede-go --env-file .env mestre-da-colheita

help:
	@echo "Comandos disponíveis:"
	@echo "  make shell  -  Entra no shell do container"
	@echo "  make logs   -  Exibe os logs do container"
	@echo "  make stop  -  Para o container"
	@echo "  make start -  Inicia o container"
	@echo "  make rm    -  Remove o container"
	@echo "  make help  -  Mostra esta ajuda"
