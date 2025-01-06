NAME = admin
MAIN = ./cmd
VERSION=$(shell TZ="Asia/Shanghai" date +"%y.%m%d.%H%M")
GOPATH ?= $(shell go env GOPATH)
BUILD_OUTPUT = ./tmp

dev:
	@if ! command -v air > /dev/null; then \
		echo "air not found, installing..."; \
		make dev_install; \
	fi
	air --build.cmd "go build -v -trimpath -ldflags='-X admin/internal/constant.VERSION=$(VERSION) -X admin/internal/constant.MODE=debug -s -w -buildid=' -o $(BUILD_OUTPUT)/$(NAME) $(MAIN)" --build.bin "$(BUILD_OUTPUT)/${NAME} run"

dev_install:
	curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(GOPATH)/bin

push:
	.github/git-push-interactive.sh

# 提交版本变更到Git
commit-version: 
	git add .
	git commit -m "bump version to v$(VERSION)"
	git push

# 创建并推送标签
push-tag: commit-version
	@echo "Creating and pushing tag v$(VERSION)"
	git tag v$(VERSION)
	git push origin v$(VERSION)

build:
	go build -v -trimpath -ldflags="-X 'admin/internal/constant.VERSION=$(VERSION)' -s -w -buildid=" -o ./tmp/admin ./cmd

docker_build:
	docker build --build-arg VERSION=$(VERSION) -t admin:$(VERSION) .