
COMMONENVVAR=GOOS=$(shell uname -s | tr A-Z a-z)
RELEASE_VERSION?=v$(shell date +%Y%m%d)-$(shell git describe --tags --match "v*")

.PHONY: all
all: cars miserables directed undirected usual-suspects

.PHONY: build
build: 
	go mod tidy
	mkdir -p ./examples/bin

# Build examples
.PHONY: cars
cars: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/bin/cars examples/cars/cars.go

.PHONY: miserables
miserables: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/bin/miserables examples/miserables/miserables.go

.PHONY: directed
directed: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/bin/hyper-directed examples/hyper-directed/hyperdirected.go

.PHONY: undirected
undirected: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/bin/hyper-undirected examples/hyper-undirected/hyperundirected.go

.PHONY: usual-suspects
usual-suspects: build
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o ./examples/bin/usual-suspects examples/usual-suspects/usual-suspects.go

.PHONY: test
test:
	./examples/bin/cars
	./examples/bin/miserables
	./examples/bin/hyper-directed
	./examples/bin/hyper-undirected	

.PHONY: clean
clean:
	rm -rf ./examples/bin/*