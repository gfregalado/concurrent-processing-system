build:
	go build -o ./bin/concurrent-processing-system

run: build
	./bin/concurrent-processing-system

test:
	go test ./...