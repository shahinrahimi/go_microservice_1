.PHONY: up down reset
up:
	docker-compose up -d --build
down:
	docker-compose down
reset: down up

