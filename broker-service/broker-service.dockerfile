FROM golang:1.23-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

# not using c library just standard GO library
RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/

RUN chmod +x /app/brokerApp

# Build a small image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerApp /app 

CMD ["/app/brokerApp"]