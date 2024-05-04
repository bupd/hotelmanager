build:
	@go build -o bin/api

seed:
	@go run scripts/seed.go

run: build
	@sudo systemctl start mongodb.service
	@./bin/api

test:
	go test -count=1 -v ./...

