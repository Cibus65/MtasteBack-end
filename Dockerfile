FROM golang:1.22.1

WORKDIR /app
COPY go.mod .
COPY go.sum .

COPY / .

RUN go mod download
RUN go build -o ../main ./cmd/main.go


CMD ["../main"]
