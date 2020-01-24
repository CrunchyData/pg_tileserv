
PROGRAM := pg_tileserv
CONTAINER := pramsey/$(PROGRAM)

.PHONY: all check clean test docker

GOFILES := $(wildcard *.go)

all: $(PROGRAM)

check:
	@go version

clean:
	$(info Cleaning project...)
	@rm -f $(PROGRAM)

$(PROGRAM): $(GOFILES)
	go build -v

docker: $(PROGRAM) Dockerfile.ci
	docker build -f Dockerfile.ci --build-arg VERSION=`./$(PROGRAM) --version | cut -f2 -d' '` -t $(CONTAINER):latest .
	docker image prune --force

test:
	go test -v


