
FROM golang:1.23-alpine AS builder

# Установка зависимостей
WORKDIR /app
COPY go.mod go.sum .
RUN ls -a
RUN go mod download

# Копируем исходный код и компилируем
COPY . ./
RUN go build -o main ./

# # Финальный контейнер для запуска приложения
# FROM alpine:latest
# WORKDIR /root/
# COPY --from=builder /app/main .
CMD ["./main"]
