up:
	docker-compose up -d

down:
	docker-compose down

.build:
	go build -o ./bin/talktalk main.go