FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o mestre-da-colheita ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/mestre-da-colheita .
COPY .env .
EXPOSE 8080
CMD ["./mestre-da-colheita"]
