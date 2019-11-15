.PHONY: dependency unit-test integration-test docker-up docker-down clear 

dependency:
	@go get -v ./...

itest: docker-up dependency
	@go test -v ./...

utest: dependency
	@go test -v -short ./...

docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down

clear: docker-down