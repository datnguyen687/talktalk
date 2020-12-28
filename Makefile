up:
	docker-compose up -d

down:
	docker-compose down

.build:
	go build -o ./bin/authentication/server ./cmd/authentication/
	cp ./config/authentication_server.json ./bin/authentication/config.json