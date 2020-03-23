.PHONY: test golint

all: test golint

test:
	make golint
	go test -cover -race -v ./...

# check syntax is valid for all files in project
golint:
	find . -name '*.go' -not -path './vendor/*' -exec dirname {} \; | xargs go vet
	find . -name '*.go' -not -path './vendor/*' -exec dirname {} \; | xargs go fmt -x
