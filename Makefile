
COMMONENVVAR=GOOS=$(shell uname -s | tr A-Z a-z)
RELEASE_VERSION?=v$(shell date +%Y%m%d)-$(shell git describe --tags --match "v*")

.PHONY: all
all: cars miserables directed undirected usual-suspects tiny

.PHONY: build
build: 
	go mod tidy
	mkdir -p ./examples/v1/bin
	mkdir -p ./examples/v2/bin

# Build examples
.PHONY: tiny
tiny: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/v1/bin/tiny examples/v1/tiny/tiny.go

.PHONY: cars
cars: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/v2/bin/cars examples/v2/cars/cars.go

.PHONY: miserables
miserables: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/v2/bin/miserables examples/v2/miserables/miserables.go

.PHONY: directed
directed: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/v2/bin/hyper-directed examples/v2/hyper-directed/hyperdirected.go

.PHONY: undirected
undirected: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/v2/bin/hyper-undirected examples/v2/hyper-undirected/hyperundirected.go

.PHONY: usual-suspects
usual-suspects: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/v2/bin/usual-suspects examples/v2/usual-suspects/usual-suspects.go

.PHONY: test
test:
	./examples/v2/bin/cars
	./examples/v2/bin/miserables
	./examples/v2/bin/hyper-directed
	./examples/v2/bin/hyper-undirected	
	./examples/v2/bin/usual-suspects
	./examples/v1/bin/tiny

.PHONY: clean
clean:
	rm -rf ./examples/bin/*