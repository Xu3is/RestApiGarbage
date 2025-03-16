dockerRun:
	docker run --name=restapigarbage -e POSTGRES_PASSWORD=swinota -p 5432:5432 -d --rm postgres
dockerStop:
	docker stop restapigarbage
migrate:
	migrate -path ./schema -database 'postgres://postgres:swinota@localhost:5432/postgres?sslmode=disable' up

