.PHONY: deps test lint lint-check-deps ci-check run-migrations dockerize push dockerize-and-push run-cdb-migrations migrate-check-deps check-cdb-env

SHELL=/bin/bash -o pipefail
IMAGE = linksrus-monolith
SHA = $(shell git rev-parse --short HEAD)

ifeq ($(origin PRIVATE_REGISTRY),undefined)
PRIVATE_REGISTRY := $(shell minikube ip 2>/dev/null):5000
endif

ifneq ($(PRIVATE_REGISTRY),)
	PREFIX:=${PRIVATE_REGISTRY}/
endif

deps:
	@go mod tidy

test:
	@echo "[go test] running tests and collecting coverage metrics"
	@go test -v -tags all_tests -race -coverprofile=coverage.txt -covermode=atomic ./...

lint: lint-check-deps
	@echo "[golangci-lint] linting sources"
	@golangci-lint run \
		-E misspell \
		-E golint \
		-E gofmt \
		-E unconvert \
		--exclude-use-default=false \
		./...

lint-check-deps:
	@if [ -z `which golangci-lint` ]; then \
		echo "[go get] installing golangci-lint";\
		GO111MODULE=on go get -u github.com/golangci/golangci-lint/cmd/golangci-lint;\
	fi

ci-check: deps lint run-cdb-migrations test

run-db-migrations: run-cdb-migrations


run-cdb-migrations: migrate-check-deps check-cdb-env
	migrate -source file://linkgraph/store/cdb/migrations -database '$(subst postgresql,cockroach,${CDB_DSN})' up

migrate-check-deps:
	@if [ -z `which migrate` ]; then \
		echo "[go get] installing golang-migrate cmd with cockroachdb support";\
		if [ "${GO111MODULE}" = "off" ]; then \
			echo "[go get] installing github.com/golang-migrate/migrate/cmd/migrate"; \
			go get -tags 'cockroachdb postgres' -u github.com/golang-migrate/migrate/cmd/migrate;\
		else \
			echo "[go get] installing github.com/golang-migrate/migrate/v4/cmd/migrate"; \
			go get -tags 'cockroachdb postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrate;\
		fi \
	fi

define dsn_missing_error

CDB_DSN envvar is undefined. To run the migrations this envvar
must point to a cockroach db instance. For example, if you are
running a local cockroachdb (with --insecure) and have created
a database called 'linkgraph' you can define the envvar by
running:

export CDB_DSN='postgresql://root@localhost:26257/linkgraph?sslmode=disable'

endef
export dsn_missing_error

check-cdb-env:
ifndef CDB_DSN
	$(error ${dsn_missing_error})
endif

generate-mocks:
	@go generate ./...

dockerize-and-push: dockerize push

dockerize:
	@echo "[docker build] building ${IMAGE} (tags: ${PREFIX}${IMAGE}:latest, ${PREFIX}${IMAGE}:${SHA})"
	@docker build --file ./Dockerfile \
		--tag ${PREFIX}${IMAGE}:latest \
		--tag ${PREFIX}${IMAGE}:${SHA} \
		. 2>&1 | sed -e "s/^/ | /g"

push:
	@echo "[docker push] pushing ${PREFIX}${IMAGE}:latest"
	@docker push ${PREFIX}${IMAGE}:latest 2>&1 | sed -e "s/^/ | /g"
	@echo "[docker push] pushing ${PREFIX}${IMAGE}:${SHA}"
	@docker push ${PREFIX}${IMAGE}:${SHA} 2>&1 | sed -e "s/^/ | /g"

