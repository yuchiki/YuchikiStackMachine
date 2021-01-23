.PHONY: ysm execute test

execute: ysm
	./ysm

test:
	go test ./...

ysm:
	go build -o ysm cmd/yuchikistackmachine/main.go
