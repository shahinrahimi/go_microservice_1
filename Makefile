.PHONY: up down reset clear
up:
	docker-compose up -d --build
down:
	docker-compose down
reset: down up
clear:
	docker image prune -f
