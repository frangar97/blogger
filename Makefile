include .env

apidocs:
	swag init --dir ./cmd/api

dev:
	@echo Running on ${PORT}

.PHONY: apidocs dev
