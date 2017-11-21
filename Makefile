.PHONY: server cli test
TOPDIR:=$(shell pwd)
OLLURL:=localhost
OLLPORT:=3000

server:
	go run server/server.go --http.addr=${OLLURL}:${OLLPORT}

cli:
	go run cli/cli.go --http.addr=http://${OLLURL}:${OLLPORT}

test:
	go tool vet ./. && go test ./...
