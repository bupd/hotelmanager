build:
	@go build -o bin/api

run: build
	@sudo systemctl start mongodb.service
	@./bin/api

test:
	go test -count=1 -v ./...

