# base go image
FROM docker.arvancloud.ir/golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/brokerApp ./cmd/api

# build a tiny docker image
FROM docker.arvancloud.ir/alpine:latest

WORKDIR /root

COPY --from=builder /bin/brokerApp .

CMD ["/root/brokerApp"]
