build:
	go build -o bin/mestre-da-colheita ./cmd/main.go

run:
	go run ./cmd/main.go

test:
	go test ./...

docker-build:
	docker build -t mestre-da-colheita .

docker-run:
	docker run -p 8080:8080 --env-file .env mestre-da-colheita
