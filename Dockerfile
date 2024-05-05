FROM golang:1.22.1
ENV GIN_MODE=release

WORKDIR /app
COPY go.mod .
COPY go.sum .

COPY / .

RUN go mod download
RUN go build -o ../main ./cmd/main.go


EXPOSE 8080
CMD ["../main"]
