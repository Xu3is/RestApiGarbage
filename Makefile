build:
	docker compose build app
run:
	docker compose up -d app
stop:
	docker compose down
migrate:
	migrate -path ./schema -database 'postgres://postgres:swinota@0.0.0.0:5432/postgres?sslmode=disable' up