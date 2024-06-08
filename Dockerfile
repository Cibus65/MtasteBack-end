FROM golang:1.22.3 AS builder
ENV GIN_MODE=release

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /main ./cmd/main.go

# Создание финального минимального образа
FROM gcr.io/distroless/base-debian10
COPY --from=builder /main /main

EXPOSE 8082
ENTRYPOINT ["/main"]
