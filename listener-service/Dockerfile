FROM docker.arvancloud.ir/golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -o /bin/listenerApp ./

# build a tiny docker image
FROM docker.arvancloud.ir/alpine:latest

WORKDIR /root

COPY --from=builder /bin/listenerApp .

CMD ["/root/listenerApp"]