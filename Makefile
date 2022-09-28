drun:
	swag init -g cmd/main/app.go
	docker-compose up -d --build
run:
	swag init -g cmd/main/app.go
	docker-compose up 
brun:
	swag init -g cmd/main/app.go
	docker-compose up --build
migrate:
	migrate -database postgres://postgres:postgres@localhost:5433/postgres?sslmode=disable -path db/migrations up