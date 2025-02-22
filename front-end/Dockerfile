FROM docker.arvancloud.ir/golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/frontApp ./cmd/web

FROM docker.arvancloud.ir/alpine:latest

WORKDIR /root

COPY --from=builder /bin/frontApp . 

COPY --from=builder /app/cmd/web/templates ./cmd/web/templates

CMD ["/root/frontApp"]
