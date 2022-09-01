.PHONY:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t telegram-bot:v0.4 .

start-container:
	docker run --env-file .env -p 80:80 telegram-bot:v0.4