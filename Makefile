include .cnf.makefile

GOVERSION := 1.15
PROGRAM := pg_tileserv

RM = /bin/rm
CP = /bin/cp
MKDIR = /bin/mkdir

.PHONY: bin-docker build-docker buildx-docker build check clean docs install install-buildx release test uninstall

.DEFAULT_GOAL := help

GOFILES := $(wildcard *.go)

all: $(PROGRAM)

check:  ##          This checks the current version of Go installed locally
	@go version

clean:  ##          This will clean all local build artifacts
	$(info Cleaning project...)
	@rm -f $(PROGRAM)
	@rm -rf docs/*
	docker image prune --force

docs:   ##           Generate docs
	@rm -rf docs/* && cd hugo && hugo && cd ..

build: $(GOFILES)  ##          Build a local binary using APPVERSION parameter or CI as default
	go build -v -ldflags "-s -w -X main.programVersion=$(APPVERSION)"

bin-docker:  ##     Build a local binary based off of a golang base docker image
	sudo docker run --rm -v "$(PWD)":/usr/src/myapp:z -w /usr/src/myapp golang:$(GOVERSION) make APPVERSION=$(APPVERSION) build

build-docker: $(PROGRAM) Dockerfile  ##   Generate a CentOS 7 container with APPVERSION tag, using binary from current environment
	docker build -f $(DOCKERFILE) --build-arg VERSION=$(APPVERSION) -t $(REPO)/$(PROGRAM):$(APPVERSION) .

buildx-docker:	##  Compiles multi-arch images and pushes them to a custom repository
	docker buildx build -f $(DOCKERFILE) --platform=$(BUILDPLATFORM) --build-arg VERSION=$(APPVERSION) -t $(REPO)/$(PROGRAM):$(APPVERSION) --push .

release: clean docs build build-docker  ##        Generate the docs, a local build, and then uses the local build to generate a CentOS 7 container

test:  ##           Run the tests locally
	go test -v

install: $(PROGRAM) docs  ##        This will install the program locally
	$(MKDIR) -p $(DESTDIR)/usr/bin
	$(MKDIR) -p $(DESTDIR)/usr/share/$(PROGRAM)
	$(MKDIR) -p $(DESTDIR)/etc
	$(CP) $(PROGRAM) $(DESTDIR)/usr/bin/$(PROGRAM)
	$(CP) config/$(PROGRAM).toml.example $(DESTDIR)/etc/$(PROGRAM).toml
	$(CP) -r assets $(DESTDIR)/usr/share/$(PROGRAM)/assets
	$(CP) -r docs $(DESTDIR)/usr/share/$(PROGRAM)/docs

install-buildx:	## This will install the buildx plugin for docker depending on the system
	wget -nc --quiet https://github.com/docker/buildx/releases/download/v0.4.2/buildx-v0.4.2.$(uname_S)-$(ARCH) -P ~/.docker/cli-plugins -O docker-buildx
	chmod a+x ~/.docker/cli-plugins/docker-buildx

uninstall:  ##      This will uninstall the program from your local system
	$(RM) $(DESTDIR)/usr/bin/$(PROGRAM)
	$(RM) $(DESTDIR)/etc/$(PROGRAM).toml
	$(RM) -r $(DESTDIR)/usr/share/$(PROGRAM)

help:   ##           Prints this help message
	@echo ""
	@echo ""
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/:.*##/:/'
	@echo ""
	@echo ""
