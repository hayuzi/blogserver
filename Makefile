.PHONY: build clean tool help

BIN_FILE=blogserver
MAIN_FILE=cmd/api/main.go
export GOPROXY=https://goproxy.cn,direct

all: build

build:
	@echo "start build"
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "${BIN_FILE}" ${MAIN_FILE}

tool:
	go tool vet . | grep -v vendor; true
	gofmt -w .

clean:
	rm -rf blogserver
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make clean: remove object files and cached files"