# Этап сборки
FROM golang:1.24.2 AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd

# Этап выполнения
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY static /app/static
COPY templates /app/templates
EXPOSE 8080
CMD ["./main"]