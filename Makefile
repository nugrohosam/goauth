.PHONY: all deps build

# This Makefile is a simple example that demonstrates usual steps to build a binary that can be run in the same
# architecture that was compiled in. The "ldflags" in the build assure that any needed dependency is included in the
# binary and no external dependencies are needed to run the service.

GOSAMPLEAPI_VERSION=$(shell git describe --always --long --dirty --tags)

all: deps build

deps:
	mv .seed.example seed.go
	go build seed.go
	mv seed.go .seed.example
	mv seed* ./cmd/

build:
	go build main.go
	mv main ./bin/
	@echo "You can now use ./bin/main"