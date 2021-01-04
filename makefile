.PHONY: ysm execute

execute: ysm
	./ysm

ysm:
	go build -o ysm cmd/yuchikistackmachine/main.go
