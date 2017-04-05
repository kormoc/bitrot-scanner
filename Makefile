
all:
	@echo "make deps 			- Refresh current godeps"
	@echo "make version VERSION=v0.0.0 	- Create a new release"

deps:
	rm -rvf Godeps vendor
	godep save ./...

version:
	git tag $(VERSION)
	git push --tags
	git push

.PHONY: deps version

