.PHONY: test golint tagversion

all: test golint tagversion

test:
	make golint
	go test -cover -race -v ./...

# check syntax is valid for all files in project
golint:
	find . -name '*.go' -not -path './vendor/*' -exec dirname {} \; | xargs go vet
	find . -name '*.go' -not -path './vendor/*' -exec dirname {} \; | xargs go fmt -x

tagversion:
	./bin/tagversion.bash
