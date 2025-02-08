FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o /bin/frontApp ./cmd

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /bin/frontApp .

COPY --from=builder /app/cmd/web/templates ./cmd/web/templates

CMD ["./frontApp"]
