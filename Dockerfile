# Многоступенчатая сборка
FROM golang:1.22.3 AS builder
ENV GIN_MODE=release

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /main ./cmd/main.go

# Финальный образ с необходимыми библиотеками
FROM debian:bookworm-slim
COPY --from=builder /main /main

EXPOSE 8082
ENTRYPOINT ["/main"]
