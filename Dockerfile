FROM golang:1.18.1

WORKDIR /app
COPY / /app

RUN go build -o main ./cmd/main.go


CMD ["main"]
