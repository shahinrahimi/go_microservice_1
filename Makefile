FRONT_BINARY = frontApp
BROKER_BINARY = brokerApp
AUTH_BINARY = authApp
BIN_DIR = "./bin"
FRONT_DIR = "./front-end"
BROKER_DIR = "./broker-service"
AUTH_DIR = "./auth-service"

up:
	@echo "Starting Docker images..."
	docker-compose up
	@echo "Docker images stared!"

build_front:
	@echo "Building front-end binary..."
	cd $(FRONT_DIR) && env CGO_ENABLED=0 go build -o ../bin/$(FRONT_BINARY) ./cmd
	@echo "Front-end binary built!"

build_broker:
	@echo "Building broker binary..."
	cd $(BROKER_DIR) && env GOOS=linux CGO_ENABLED=0 go build -o ../bin/$(BROKER_BINARY) ./cmd
	@echo "Broker binary built!"

build_auth:
	@echo "Building auth binary..."
	cd $(AUTH_DIR) && env GOOS=linux CGO_ENABLED=0 go build -o ../bin/$(AUTH_BINARY) ./cmd
	@echo "Auth binary built!"

up_build: build_front build_broker build_auth
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting Docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"


down:
	@echo "Stopping Docker images..."
	docker-compose down
	@echo "Docker images stopped!"



start: build_front
	@echo "Starting front-end"
	cd ./front-end && ./cmd/web/$(FRONT_END_BINARY)

stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./cmd/web/${FRONT_END_BINARY}"
	@echo "Stopped front end!"




