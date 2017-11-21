.PHONY: server cli test
DATE=`date +%Y-%m-%d`
TOPDIR:=$(shell pwd)
OLLURL:=localhost
OLLPORT:=3000

srv:
	go build -ldflags="-X main.buildName=$(DATE)" server/oll-srv.go

cli:
	go build -ldflags="-X main.buildName=$(DATE)" cli/oll-cli.go

test:
	go tool vet ./. && go test ./...

run-srv:
	go run server/oll-srv.go -http.addr=${OLLURL}:${OLLPORT}

run-cli:
	go run cli/oll-cli.go -http.addr=http://${OLLURL}:${OLLPORT}