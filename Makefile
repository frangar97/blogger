include .env

apidocs:
	swag init --dir ./cmd/api

dev: apidocs
	@echo Running on ${PORT}
	@go run ./cmd/api

migrateup:
	migrate -path migrations -database ${POSTGRES_URL} -verbose up

migratedown:
	migrate -path migrations -database ${POSTGRES_URL} -verbose down

.PHONY: apidocs dev migrateup migratedown
