.PHONY: generator

PKGS := $(shell go list ./... | grep -v go-errgen)

generate: generator
	./generator

generator:
	go build ./internal/generator.go

check: test lint vet fmt-check

test:
	go test -v -cover $(PKGS)

lint:
	golint -set_exit_status $(PKGS)

vet:
	go vet $(PKGS)

fmt-check:
	gofmt -l -s **/*.go | grep [^*][.]go$$; \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then exit 1; fi; \
	goimports -l **/*.go | grep [^*][.]go$$; \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then exit 1; fi \

fmt:
	gofmt -w -s **/*.go
	goimports -w **/*.go
