FROM docker.arvancloud.ir/golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -o /bin/mailerApp ./cmd/api

# build a tiny docker image
FROM docker.arvancloud.ir/alpine:latest

WORKDIR /root

COPY --from=builder /bin/mailerApp .
COPY --from=builder /app/templates ./templates
CMD ["/root/mailerApp"]