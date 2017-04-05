VERSION :=$(shell git describe --tags )

all:
	go build -ldflags "-X main.Version=$(VERSION)"

deps:
	rm -rvf Godeps vendor
	godep save ./...

version:
	git tag $(VERSION)
	git push --tags
	git push

.PHONY: deps version
