FRONT_END_BINARY = frontApp
BROKER_BINARY = brokerApp

up:
	@echo "Starting Docker images..."
	docker-compose up
	@echo "Docker images stared!"

build_broker:
	@echo "Building broker binary..."
	cd ./broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ./cmd/api/$(BROKER_BINARY) ./cmd
	@echo "Broker binary built!"

up_build: build_broker
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting Docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"


down:
	@echo "Stopping Docker images..."
	docker-compose down
	@echo "Docker images stopped!"

build_front:
	@echo "Building front-end binary..."
	cd ./front-end && env CGO_ENABLED=0 go build -o ./cmd/web/$(FRONT_END_BINARY) ./cmd
	@echo "Front-end binary built!"

start: build_front
	@echo "Starting front-end"
	cd ./front-end && ./cmd/web/$(FRONT_END_BINARY)

stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./cmd/web/${FRONT_END_BINARY}"
	@echo "Stopped front end!"


