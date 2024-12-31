dev:
	go run cmd/cli/*.go

build:
	go build -o build/btree cmd/cli/*.go

run:
	./build/btree
