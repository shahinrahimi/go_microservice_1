FROM docker.arvancloud.ir/golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -o /bin/authApp ./cmd/api

FROM alpine:latest

WORKDIR /root

COPY --from=builder /bin/authApp .

CMD ["/root/authApp"]
