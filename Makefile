all:
	go build -ldflags "-X main.Version=$(shell git describe --tags)"

test:
	go test -bench=. ./...

version:
	git tag $(VERSION)
	git push --tags
	git push

.PHONY: version
