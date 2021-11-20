build:
	@go build  -o dist/tcpscan .
format:
	@go fmt .
all: format build
