.DEFAULT_GOAL := help
MAKEFILE_DIR := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: setup
setup: ## Set up 
	cp -n ./env/sample.env ./env/local.env || true

.PHONY: up
up: ## Run docker containers 
	docker-compose up

.PHONY: build
build: ## Build docker containers
	docker-compose build

.PHONY: ps 
ps: ## Show docker container processes 
	docker-compose ps 

.PHONY: down
down: ## Stop docker containers
	docker-compose down

.PHONY: bash 
bash: ## Run bash in app container
	docker-compose exec app bash 

.PHONY: ping 
ping: ## Send email to mock-sendgrid form app container
	docker-compose exec app curl \
		--dump-header -  \
		--request POST \
		--url http://mock-sendgrid:3030/v3/mail/send \
		--header 'Authorization: Bearer SG.xxxxx' \
		--header 'Content-Type: application/json' \
		--data '{"personalizations": [{ \
    				"to": [{"email": "to@example.com"}]}], \
    				"from": {"email": "from@example.com"}, \
    				"subject": "Test Subject", \
    				"content": [{"type": "text/plain", "value": "Test Content"}] \
			}'

.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'