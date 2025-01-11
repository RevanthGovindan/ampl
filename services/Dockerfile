FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o myapp ./src

FROM alpine:latest

WORKDIR /root/

COPY config.yaml .
COPY ./keys ./keys

COPY --from=builder /app/myapp .

CMD ["./myapp"]
