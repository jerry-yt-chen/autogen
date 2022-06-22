APP_NAME={{.ProjectName}}
VERSION=latest
DOCKER_BUILD_PROGRESS=auto
DOCKER_BUILD_SSH=$(HOME)/.ssh/id_rsa

.PHONY: run
run: tidy di build
	docker-compose -f deployments/docker-compose.yaml up -d

.PHONY: build
build:
	DOCKER_BUILDKIT=1 docker build --progress=$(DOCKER_BUILD_PROGRESS) --ssh main=$(DOCKER_BUILD_SSH) -t $(APP_NAME):$(VERSION) -f ./build/Dockerfile .

.PHONY: stop
stop:
	docker-compose -f deployments/docker-compose.yaml down

.PHONY: di
di:
	go install github.com/google/wire/cmd/wire@v0.5.0
	wire gen ./internal/injector

.PHONY: tidy
tidy:
	go mod tidy