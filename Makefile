all:
	go build -ldflags "-X main.Version=$(shell git describe --tags)"

test:
	go test -bench=. ./...

deps:
	rm -rvf Godeps vendor
	godep get
	godep save ./...

version:
	git tag $(VERSION)
	git push --tags
	git push

.PHONY: deps version
