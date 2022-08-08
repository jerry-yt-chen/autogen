TYPE=category
PROJ_NAME=sample-svc

.PHONY: build
build:
	mkdir bin || true
	go build -o ./bin/autogen ./cmd/autogen/main.go

.PHONY: run
run: build
	./bin/autogen new $(PROJ_NAME) --type=$(TYPE)
