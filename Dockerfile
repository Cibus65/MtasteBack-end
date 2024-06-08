FROM golang:1.22.3
ENV GIN_MODE=release

WORKDIR /app
COPY go.mod .
COPY go.sum .

COPY / .

RUN go mod download
RUN go build -o ../main ./cmd/main.go

FROM nginx:alpine

# Копирование собранного приложения в каталог html Nginx
COPY --from=build-stage /app/main /usr/share/nginx/html

# Открытие порта 80
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

