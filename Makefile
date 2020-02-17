.PHONY: all
all: test vet

.PHONY: test
test:
	@go test ./test -cover

.PHONY: vet
vet:
	@go vet -all ./rakuten
