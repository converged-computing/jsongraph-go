
COMMONENVVAR=GOOS=$(shell uname -s | tr A-Z a-z)
RELEASE_VERSION?=v$(shell date +%Y%m%d)-$(shell git describe --tags --match "v*")

.PHONY: all
all: build

# Build examples
.PHONY: build
build: 
	go mod tidy
	mkdir -p ./examples/bin
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/bin/cars examples/cars.go

.PHONY: clean
clean:
	rm -rf ./bin/server