include init.mk

GOPRIVATE:=gitlab.com/mcsolutions/find-psy
PROJECT:=find-psy
REGISTRY:=registry
REPOSITORY:=gitlab.com/mcsolutions/find-psy/back/article-v1
DOCKER_IMAGE_NAME:=$(REGISTRY).$(REPOSITORY)
CONFIG_VOLUME:=/data/conf
PORT_GRPC:=9000
PORT_HTTP:=8000
EXT_PORT_GRPC:=9000
EXT_PORT_HTTP:=8000
MUSEUM:=https://mcsolutions.gitlab.io/charts/find-psy/
CHART:=sfpg
RELEASE_NAME:=back-article-v1
STAGING:=find-psy-staging

GOPATH:=$(shell go env GOPATH)
VERSION:=$(shell git describe --tags --always --abbrev=0)
GIT_COMMIT:=$(shell git rev-list -1 HEAD)
CONF_PROTO_FILES:=$(shell find internal/conf -name *.proto)

print:
	@echo "GOPATH" $(GOPATH)
	@echo "VERSION" $(VERSION)
	@echo "GIT_COMMIT" $(GIT_COMMIT)
	@echo "CONF_PROTO_FILES" $(CONF_PROTO_FILES)

version:
	@echo $(VERSION)

v: version

init-ent:
	@ent init --target ./internal/outside/data/ent/schema Article
	@go get entgo.io/ent
	@go mod vendor
	@ent generate ./internal/outside/data/ent/schema
	@echo "package ent" > ./internal/outside/data/ent/generate.go
	@echo "\n//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema" >> ./internal/outside/data/ent/generate.go

.PHONY: config
config:
	@protoc \
		-I=. \
		-I=../../proto \
 		--go_out=paths=source_relative:. \
		$(CONF_PROTO_FILES)
#		--proto_path=$(GOPATH)/src/gitlab.com/mcsolutions/find-psy/proto/third_party \

.PHONY: tidy
tidy:
	@go mod tidy -compat=1.17

.PHONY: vendor
vendor:
	@go mod vendor

.PHONY: dep
dep: tidy vendor

.PHONY: wire
wire:
	@wire gen ./cmd/article-v1

.PHONY: ent
ent:
	@ent generate ./internal/outside/data/ent/schema

.PHONY: build
build:
	@mkdir -p bin/ && rm -f bin/* && \
	GO111MODULE=on CGO_ENABLE=0 GOOS=linux GOARCH=amd64 \
	go build -mod=vendor \
	-ldflags "-s -w -X main.Version=$(VERSION)" \
	-o ./bin/ \
	./... && \
	find bin -type f -exec mv {} bin/bin \;

.PHONY: compose
compose:
	@docker-compose up -d

.PHONY: compose-stop
compose-stop:
	@docker-compose stop

.PHONY: run
run:
	GO111MODULE=on go run -ldflags "-s -w -X main.Version=$(VERSION)" ./cmd/article-v1/

.PHONY: test
test:
	@GO111MODULE=on go test -v -race ./...

.PHONY: test-cover
test-cover:
	@GO111MODULE=on go test -v -race -cover ./internal/service
	@GO111MODULE=on go test ./internal/service -coverprofile=coverage.out
	@go tool cover -func=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@firefox coverage.html

init-push:
	@git push --set-upstream origin master

push:
	@git status
	@git push origin master
	@git push --tags

build-image:
	@docker build -t $(DOCKER_IMAGE_NAME):$(VERSION) .

run-image:
	@echo $(GOPATH)/src/$(REPOSITORY)/configs:$(CONFIG_VOLUME)
	@docker run -d \
	-p $(EXT_PORT_HTTP):$(PORT_HTTP) \
	-p $(EXT_PORT_GRPC):$(PORT_GRPC) \
	-v $(GOPATH)/src/$(REPOSITORY)/configs:$(CONFIG_VOLUME) \
	--name=article-v1 \
	$(DOCKER_IMAGE_NAME):$(VERSION)

push-image:
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) $(DOCKER_IMAGE_NAME):latest
	docker push $(DOCKER_IMAGE_NAME):$(VERSION)
	docker push $(DOCKER_IMAGE_NAME):latest

rmi:
	docker rmi $(docker images | grep $(DOCKER_IMAGE_NAME))

images:
	docker images $(DOCKER_IMAGE_NAME)

.PHONY: create-service
create-service:
	mkdir -p internal/service
	find api -name "*.proto" -exec kratos proto server {} \;

.PHONY: tags
tags:
	@git tag --sort=-creatordate -n | head -n 10
