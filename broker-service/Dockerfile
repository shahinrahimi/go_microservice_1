FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o /bin/brokerApp ./cmd

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /bin/brokerApp .

CMD ["./brokerApp"]
