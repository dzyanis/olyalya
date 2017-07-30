TOPDIR:=$(shell pwd)
OLLURL:=localhost
OLLPORT:=3000

server:
	go run olldb/olldb.go --http.addr=${OLLURL}:${OLLPORT}

cli:
	go run oll-cli/oll-cli.go --http.url=${OLLURL} --http.port=${OLLPORT}