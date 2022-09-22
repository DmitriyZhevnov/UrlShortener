drun:
	docker-compose up -d --build
run:
	docker-compose up 
brun:
	docker-compose up --build
migrate:
	migrate -database postgres://postgres:postgres@localhost:5433/postgres?sslmode=disable -path db/migrations up