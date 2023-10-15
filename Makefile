build:
	docker-compose up --build epta-app
run:
	docker-compose up epta-app
swag:
	swag init -g cmd/main.go