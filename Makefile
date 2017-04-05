all:
	go build -ldflags "-X main.Version=$(shell git describe --tags)"

deps:
	rm -rvf Godeps vendor
	godep save ./...

version:
	git tag $(VERSION)
	git push --tags
	git push

.PHONY: deps version
