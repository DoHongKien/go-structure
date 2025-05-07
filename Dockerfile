FROM golang:1.24-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/server .
COPY --from=builder /app/config ./config

# Thêm biến môi trường cho cấu hình database
ENV MYSQL_HOST=mysql_container
ENV MYSQL_PORT=3306
ENV MYSQL_USER=root
ENV MYSQL_PASSWORD=root1234
ENV MYSQL_DB=shopdevgo
ENV REDIS_HOST=redis_container_bloom
ENV REDIS_PORT=6379

EXPOSE 9999

CMD ["./server"]
