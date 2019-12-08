all: build run

build:
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -o ./bin/embryo-darwin-amd64-v1.0.0
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -o ./bin/embryo-linux-amd64-v1.0.0
    # CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o embryo-windows-amd64-v2.0.1

run:
	./bin/embryo-darwin-amd64-v1.0.0 -d=false