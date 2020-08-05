GO = go

SHELL = /bin/bash
$(info $(SHELL))

.PHONY: setup
setup: ## Install all the build and lint dependencies
	go get -u github.com/alecthomas/gometalinter
	go get -u golang.org/x/tools/cmd/cover
	# go get -u github.com/golang/dep/cmd/dep
	gometalinter --install --update
	@$(MAKE) dep

.PHONY: dep
dep: ## Run dep ensure and prune
	dep ensure
	dep prune

.PHONY: test
test: ## Run all the tests
	echo 'mode: atomic' > coverage.txt && $(GO) test -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=30s ./...

.PHONY: cover
cover: test ## Run all the tests and opens the coverage report
	go tool cover -html=coverage.txt

.PHONY: fmt
fmt: ## Run goimports on all go files
	gofmt -w .
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do goimports -w "$$file"; done
	find . -name '*.go' -not -path './vendor/*' -exec dirname {} \; | xargs go fmt -x


.PHONY: lint
lint: ## Run all the linters
	find . -type f \( -name "*.go" ! -name "*_test.go" \) -not -path './vendor/*' -exec dirname {} \; | xargs go vet || true
	gometalinter --vendor --disable-all \
    		--enable=deadcode \
    		--enable=ineffassign \
    		--enable=gosimple \
    		--enable=staticcheck \
    		--enable=gofmt \
    		--enable=goimports \
    		--enable=misspell \
    		--enable=errcheck \
    		--enable=vet \
    		--enable=vetshadow \
    		--deadline=10m \
    		./...

.PHONY: ci
ci: lint test ## Run all the tests and code checks

.PHONY: tagver
tagver: ## Tag the branch and push new git tag to origin
	./bin/tagversion.bash

.PHONY: build
build: ## Build a version
	$(GO) build -v ./...

.PHONY: clean
clean: ## Remove temporary files
	go clean

.PHONY: help
help: ## make help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

changelog:
	git-chglog -o CHANGELOG.md

release:
	git-chglog --next-tag ${TAG} ${TAG} -o CHANGELOG.md
	git add CHANGELOG.md
	git commit -m "Update changelog with ${TAG} changes"
	git tag ${TAG}
	git-chglog $(TAG) | tail -n +4 | sed '1s/^/$(TAG)\n/gm' > release-notes.txt
	git push origin master ${TAG}

.DEFAULT_GOAL := test
