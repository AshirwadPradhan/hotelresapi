build:
	@go build -o bin/api

run: build
	@./bin/api

test:
	go test -v ./...

rundb:
	@docker run --name mongodb -p 27017:27017 -d docker.io/mongo:latest 

rmdb:
	@docker stop mongodb
	@docker rm mongodb