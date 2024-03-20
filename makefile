all: download build run

download:
	go mod download
build:
	go build -o builded_app ./cmd/user_app
run:
	cd deployments && docker compose up -d
	go run ./cmd/user_app/main.go