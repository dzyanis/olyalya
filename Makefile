VERSION=0.0.1
DATE=`date +%Y-%m-%d`
TOPDIR:=$(shell pwd)
HOST:=localhost
PORT:=3000

.PHONY: test
test:
	go tool vet ./. && go test ./...

.PHONY: srv
srv: test
	go build -ldflags="-X main.buildName=${DATE} -X main.version=${VERSION}" server/oll-srv.go -o bin/olyalya

.PHONY: srv-install
srv: test
	go install -ldflags="-X main.buildName=${DATE} -X main.version=${VERSION}" server/oll-srv.go

.PHONY: run-srv
run-srv:
	go run apps/server/oll-srv.go -http.addr=${HOST}:${PORT}

.PHONY: dkr-srv-build
dkr-srv-build:
	docker build -f ./dockerfiles/Dockerfile.server -t dzyanis/olyalya:${VERSION} .

.PHONY: dkr-srv-run
dkr-srv-run:
	docker run --name olyalya -p ${PORT}:${PORT} -d -it dzyanis/olyalya:${VERSION} 1> .dkr-container-id.server

.PHONY: dkr-srv-stop
dkr-srv-stop:
	test -s .dkr-container-id.server && docker stop $(shell cat .dkr-container-id.server) && rm .dkr-container-id.server

.PHONY: cli
cli: test
	go build -ldflags="-X main.buildName=${DATE} -X main.version=${VERSION}" cli/oll-cli.go

.PHONY: run-cli
run-cli:
	go run apps/client/oll-cli.go -http.addr=http://${HOST}:${PORT}

.PHONY: dkr-cli-build
dkr-cli-build:
	docker build -f ./dockerfiles/Dockerfile.client -t dzyanis/oll-cli:${VERSION} .

.PHONY: dkr-cli-run
dkr-cli-run:
	docker run --name oll-cli -it dzyanis/oll-cli:${VERSION} > .dkr-container-id.client

.PHONY: dkr-cli-stop
dkr-cli-stop:
	test -s .dkr-container-id.client && docker stop $(shell cat .dkr-container-id.client) && rm .dkr-container-id.client
