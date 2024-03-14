.PHONY:
.SILENT:


build:
	go build -o ./.bin/bot cmd/app/main.go

run: build
	./.bin/app

build-image:
	docker build -t recurro:v0.1 .

start-container:
	docker run --name recurro-app -p 8065:8065 --env-file .env recurro:v0.1

start-container-front:
	docker run -p 8080:80 my-react-app

start-compose:
	docker-compose up

lint:
	golangci-lint -v run ./... --config=golangci.yml
