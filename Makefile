TARGETS           ?= linux/amd64 darwin/amd64 linux/amd64 windows/amd64 linux/386 linux/ppc64le linux/s390x linux/arm linux/arm64
PROJECT_NAME	  := prometheus-cachethq
PKG				  := github.com/oxyno-zeta/$(PROJECT_NAME)

# go option
GO        ?= go
# Uncomment to enable vendor
GO_VENDOR := # -mod=vendor
TAGS      :=
TESTS     := .
TESTFLAGS :=
LDFLAGS   := -w -s
GOFLAGS   :=
BINDIR    := $(CURDIR)/bin
DISTDIR   := dist

# Required for globs to work correctly
SHELL=/usr/bin/env bash

#  Version

GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
DATE	   = $(shell date +%F_%T%Z)

BINARY_VERSION = ${GIT_SHA}
LDFLAGS += -X ${PKG}/pkg/${PROJECT_NAME}/version.Version=${BINARY_VERSION}
LDFLAGS += -X ${PKG}/pkg/${PROJECT_NAME}/version.GitCommit=${GIT_COMMIT}
LDFLAGS += -X ${PKG}/pkg/${PROJECT_NAME}/version.BuildDate=${DATE}

HAS_GORELEASER := $(shell command -v goreleaser;)
HAS_GIT := $(shell command -v git;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
HAS_CURL:=$(shell command -v curl;)
HAS_MOCKGEN:=$(shell command -v mockgen;)
# Uncomment to use gox instead of goreleaser
# HAS_GOX := $(shell command -v gox;)

.DEFAULT_GOAL := code/lint

#############
#   Build   #
#############

.PHONY: code/lint
code/lint: setup/dep/install
	golangci-lint run ./...

.PHONY: code/generate
code/generate:
	$(GO) $(GO_VENDOR) generate ./...

.PHONY: code/graphql
code/graphql: code/graphql/generate code/graphql/concat

.PHONY: code/build
code/build: code/clean setup/dep/install
	$(GO) build $(GO_VENDOR) -o $(BINDIR)/$(PROJECT_NAME) $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' $(PKG)/cmd/${PROJECT_NAME}

# Uncomment to use gox instead of goreleaser
# .PHONY: code/build-cross
# code/build-cross: code/clean setup/dep/install
# 	CGO_ENABLED=0 GOFLAGS="-trimpath $(GO_VENDOR)" gox -output="$(DISTDIR)/bin/{{.OS}}-{{.Arch}}/{{.Dir}}" -osarch='$(TARGETS)' $(if $(TAGS),-tags '$(TAGS)',) -ldflags '$(LDFLAGS)' ${PKG}/cmd/${PROJECT_NAME}

.PHONY: code/build-cross
code/build-cross: code/clean setup/dep/install
ifdef HAS_GORELEASER
	goreleaser --snapshot --skip-publish
endif
ifndef HAS_GORELEASER
	curl -sL https://git.io/goreleaser | bash -s -- --snapshot --skip-publish
endif

.PHONY: code/clean
code/clean:
	@rm -rf $(BINDIR) $(DISTDIR)

#############
#  Release  #
#############

# Uncomment to use gox instead of goreleaser
# .PHONY: release/all
# release/all: code/clean setup/dep/install code/build-cross
# 	cp Dockerfile $(DISTDIR)/bin/linux-amd64

.PHONY: release/all
release/all: code/clean setup/dep/install
ifdef HAS_GORELEASER
	goreleaser
endif
ifndef HAS_GORELEASER
	curl -sL https://git.io/goreleaser | bash
endif

#############
#   Tests   #
#############

.PHONY: test/all
test/all: setup/dep/install
	$(GO) test $(GO_VENDOR) --tags=unit,integration -v -coverpkg=./pkg/... -covermode=count -coverprofile=c.out.tmp ./pkg/...

.PHONY: test/unit
test/unit: setup/dep/install
	$(GO) test $(GO_VENDOR) --tags=unit -v -coverpkg=./pkg/... -covermode=count -coverprofile=c.out.tmp ./pkg/...

.PHONY: test/integration
test/integration: setup/dep/install
	$(GO) test $(GO_VENDOR) --tags=integration -v -coverpkg=./pkg/... -covermode=count -coverprofile=c.out.tmp ./pkg/...

.PHONY: test/coverage
test/coverage:
	cat c.out.tmp | grep -v "mock_" > c.out
	$(GO) tool cover -html=c.out -o coverage.html
	$(GO) tool cover -func c.out

#############
#   Setup   #
#############

.PHONY: down/services
down/services:
	@echo "Down services"
	docker rm -f postgres || true
	docker rm -f maildev || true
	docker rm -f cachet || true

.PHONY: down/metrics-services
down/metrics-services:
	@echo "Down metrics services"
	docker rm -f prometheus || true
	docker rm -f grafana || true
	docker rm -f jaeger || true

.PHONY: down/dev-services
down/dev-services:
	@echo "Down dev services"
	docker rm -f pgadmin || true

.PHONY: setup/dev-services
setup/dev-services: down/dev-services
	@echo "Setup dev services"
	docker run --rm --name pgadmin -p 8090:80 --link postgres:postgres -e 'PGADMIN_DEFAULT_EMAIL=user@domain.com' -e 'PGADMIN_DEFAULT_PASSWORD=SuperSecret' -d dpage/pgadmin4

.PHONY: setup/metrics-services
setup/metrics-services: down/metrics-services
	@echo "Setup metrics services"
	docker run --rm -d --name prometheus -v $(CURDIR)/.local-resources/prometheus/prometheus.yml:/prometheus/prometheus.yml --network=host prom/prometheus:v2.18.0 --web.listen-address=:9191
	docker run --rm -d --name grafana --network=host grafana/grafana:7.0.3
	docker run --rm --name jaeger -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest

.PHONY: setup/services
setup/services: down/services
	@echo "Setup services"
	docker run -d --rm --name postgres -p 5432:5432 -e POSTGRES_USER=cachet -e POSTGRES_PASSWORD=cachet -e PGDATA=/var/lib/postgresql/data/pgdata -v $(CURDIR)/.run/postgres:/var/lib/postgresql/data postgres:12
	sleep 1
	PGPASSWORD=cachet psql --username cachet -h localhost -f .local-resources/postgresql/cachet.sql
	docker run --rm --name maildev -p 1080:1080 -p 1025:1025 -d maildev/maildev:1.1.0 --incoming-user fake --incoming-pass fakepassword
	docker run -d --name cachet -v $(CURDIR)/.local-resources/cachethq:/var/www/html/bootstrap/cachet/ --link postgres:postgres -p 8000:8000 -e DB_DRIVER=pgsql \
		-e DB_HOST=postgres -e DB_DATABASE=cachet -e DB_USERNAME=cachet -e DB_PASSWORD=cachet \
		-e MAIL_ADDRESS=fake@fake.com -e MAIL_DRIVER=smtp -e MAIL_HOST=localhost -e MAIL_PORT=1025 -e MAIL_USERNAME=fake -e MAIL_PASSWORD=fakepassword \
		-e APP_KEY=base64:5azW+xGOYjEX9Bq9RKksKvJlvNLbUXrUT4e0TpduS1g= \
		cachethq/docker:2.3.14

.PHONY: setup/dep/install
setup/dep/install:
ifndef HAS_GOLANGCI_LINT
	@echo "=> Installing golangci-lint tool"
ifndef HAS_CURL
	$(error You must install curl)
endif
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.42.1
endif
ifndef HAS_GIT
	$(error You must install Git)
endif
ifndef HAS_MOCKGEN
	@echo "=> Installing mockgen tool"
	$(GO) get -u github.com/golang/mock/mockgen@v1.6.0
endif
# Uncomment to use gox instead of goreleaser
# ifndef HAS_GOX
# 	@echo "=> Installing gox"
# 	$(GO) get -u github.com/mitchellh/gox
# endif
	$(GO) mod download

.PHONY: setup/dep/tidy
setup/dep/tidy:
	$(GO) mod tidy

.PHONY: setup/dep/update
setup/dep/update:
	$(GO) get -u ./...

.PHONY: setup/dep/vendor
setup/dep/vendor:
	$(GO) mod vendor
