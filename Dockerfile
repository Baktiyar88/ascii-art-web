# Этап сборки
FROM golang:1.24.2 AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -o main ./cmd
RUN chmod +x /app/main

# Этап выполнения
FROM alpine:latest

# Устанавливаем bash, если нужен shell-доступ
RUN apk add --no-cache bash libc6-compat ca-certificates

WORKDIR /app
COPY --from=builder /app/main .
COPY static /app/static
COPY templates /app/templates
COPY banners /app/banners

EXPOSE 8080
CMD ["./main"]
