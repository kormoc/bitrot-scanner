
all:
	go build

deps:
	rm -rvf Godeps vendor
	godep save ./...

version:
	git tag $(VERSION)
	git push --tags
	git push

.PHONY: deps version

