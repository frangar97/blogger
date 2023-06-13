include .env

apidocs:
	swag init --dir ./cmd/api

dev: apidocs
	@echo Running on ${PORT}
	@go run ./cmd/api

.PHONY: apidocs dev
