# parameters
GODIRS_NOVENDOR = $(shell go list ./... | grep -v vendor/)
GOFILES_NOVENDOR = $(shell find . -type f -name '*.go' -not -path "./vendor/*")


# targets
checkstyle:
	golangci-lint run --timeout=10m -v ./...

fmt:
	gofmt -l -w -s ${GOFILES_NOVENDOR}
	goimports -l -w ${GOFILES_NOVENDOR}

run:
	docker-compose up -d

stop:
	docker-compose stop
