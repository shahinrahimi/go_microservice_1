front: 
	cd ./front-end && go run ./cmd/web/main.go
broker:
	cd ./broker-service/cmd && go run .
dock:
	docker-compose up -d


