.PHONY: build
build:
	mkdir bin || true
	go build -o ./bin