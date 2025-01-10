FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy
RUN ls -la /app
RUN go build -o myapp .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/myapp .

CMD ["./myapp"]
