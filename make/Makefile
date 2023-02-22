include common/go.mk

.PHONY: all
all: ## build pipeline
all: fmt test

.PHONY: test
test: ## go test with race detector and code covarage
	$(call print-target)
	EXIT=0 ;\
	go test -race -covermode=atomic -coverprofile=coverage.out ./... || EXIT=$$? ;\
	go tool cover -html=coverage.out -o coverage.html ;\
	exit $$EXIT
