all: test golint tagver

SHELL = /bin/bash
$(info $(SHELL))

.PHONY: test
test:
	go test -cover -race -v ./...

# check syntax is valid for all files in project
.PHONY: golint
golint:
	find . -type f \( -name "*.go" ! -name "*_test.go" \) -not -path './vendor/*' -exec dirname {} \; | xargs go vet || true
	find . -name '*.go' -not -path './vendor/*' -exec dirname {} \; | xargs go fmt -x

.PHONY: tagver
tagver:
	./bin/tagversion.bash
