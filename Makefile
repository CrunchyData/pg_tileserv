##AVAILABLE BUILD OPTIONS -
##      APPVERSION - Variable to set the version label
##      GOVERSION - Defaults to 1.21.6 but can be overriden, uses alpine go container as base
##      PROGRAM - Name of binary, pg_tileserv
##      CONTAINER - prefix and name of the generated container
##      CONFIG - config file to be used
##      DATE - Date String used as alternate tag for generated containers
##      BASE_REGISTRY - This is the registry to pull the base image from
##      BASE_IMAGE - The base image to use for the final container
##      TARGETARCH - The architecture the resulting image is based on and the binary is compiled for
##      IMAGE_TAG - The tag to be applied to the container

APPVERSION ?= latest
GOVERSION ?= 1.21.6
PROGRAM ?= pg_tileserv
CONFIG ?= config/$(PROGRAM).toml
CONTAINER ?= harbor.internal.millcrest.dev/library/$(PROGRAM)
DATE ?= $(shell date +%Y%m%d)
BASE_REGISTRY ?= registry.access.redhat.com
BASE_IMAGE ?= ubi8-micro
SYSTEMARCH = $(shell uname -i)

# ifeq ($(SYSTEMARCH), x86_64)
# TARGETARCH ?= amd64
# PLATFORM=amd64
# else
# TARGETARCH ?= arm64
# PLATFORM=arm64
# endif
TARGETARCH ?= amd64
PLATFORM ?= amd64

IMAGE_TAG ?= $(APPVERSION)-$(TARGETARCH)
DATE_TAG ?= $(DATE)-$(TARGETARCH)

RM = /bin/rm
CP = /bin/cp
MKDIR = /bin/mkdir

.PHONY: bin-docker bin-for-docker build build-docker check clean common-build docker docs install multi-stage-docker release set-local set-multi-stage test uninstall

.DEFAULT_GOAL := help

GOFILES := $(wildcard *.go)

all: $(PROGRAM)

check:  ##              This checks the current version of Go installed locally
	@go version

clean:  ##              This will clean all local build artifacts
	$(info Cleaning project...)
	@rm -f $(PROGRAM)
	@rm -rf docs/*
	@docker image inspect $(CONTAINER):$(IMAGE_TAG) >/dev/null 2>&1 && docker rmi -f $(shell docker images --filter label=release=latest --filter=reference="*tileserv:*" -q) || echo -n ""

docs:   ##               Generate docs
	@rm -rf docs/* && cd hugo && hugo && cd ..

build: $(PROGRAM) ##              Build a local binary using APPVERSION parameter

$(PROGRAM): $(GOFILES)
	go build -v -ldflags "-s -w -X main.programVersion=$(APPVERSION)"

bin-docker:  ##         Build a local binary based off of a golang base docker image
	sudo docker run --rm -v "$(PWD)":/usr/src/myapp:z -w /usr/src/myapp golang:$(GOVERSION) make APPVERSION=$(APPVERSION) build

bin-for-docker: $(GOFILES)  ##     Build a local binary using APPVERSION parameter or CI as default (to be used in docker image)
# to be used in docker the built binary needs the CGO_ENABLED=0 option
	CGO_ENABLED=0 go build -v -ldflags "-s -w -X main.programVersion=$(APPVERSION)"

build-common: Dockerfile
	docker build -f Dockerfile \
		--target $(BUILDTYPE) \
		--build-arg VERSION=$(APPVERSION) \
		--build-arg GOLANG_VERSION=$(GOVERSION) \
		--build-arg TARGETARCH=$(TARGETARCH) \
		--build-arg PLATFORM=$(PLATFORM) \
		--build-arg BASE_REGISTRY=$(BASE_REGISTRY) \
		--build-arg BASE_IMAGE=$(BASE_IMAGE) \
		--label vendor="Crunchy Data" \
		--label url="https://crunchydata.com" \
		--label release="$(APPVERSION)" \
		--label org.opencontainers.image.vendor="Crunchy Data" \
		--label os.version="7.7" \
		-t $(CONTAINER):$(IMAGE_TAG) -t $(CONTAINER):$(DATE_TAG) .
	docker image prune --filter label=stage=tileservbuilder -f

push-docker:
	docker push ${CONTAINER}:${IMAGE_TAG}

set-local:
	$(eval BUILDTYPE = local)

set-multi-stage:
	$(eval BUILDTYPE = multi-stage)

# This is just an alias to keep the existing targets available
docker-build: docker

docker: bin-for-docker Dockerfile set-local build-common ##             Generate a BASE_IMAGE container with APPVERSION tag, using a locally built binary

multi-stage-docker: Dockerfile set-multi-stage build-common ## Generate a BASE_IMAGE container with APPVERSION tag, using a binary built in an alpine golang build container

release: clean docs docker  ##            Generate the docs, a local build, and then uses the local build to generate a BASE_IMAGE container

test:  ##               Run the tests locally
	go test -v

$(CONFIG): $(CONFIG).example
	sed 's/# AssetsPath/AssetsPath/' $< > $@

install: $(PROGRAM) docs $(CONFIG) ##            This will install the program locally
	$(MKDIR) -p $(DESTDIR)/usr/bin
	$(MKDIR) -p $(DESTDIR)/usr/share/$(PROGRAM)
	$(MKDIR) -p $(DESTDIR)/etc
	$(CP) $(PROGRAM) $(DESTDIR)/usr/bin/$(PROGRAM)
	$(CP) $(CONFIG) $(DESTDIR)/etc/
	$(CP) -r assets $(DESTDIR)/usr/share/$(PROGRAM)/assets
	$(CP) -r docs $(DESTDIR)/usr/share/$(PROGRAM)/docs

uninstall:  ##          This will uninstall the program from your local system
	$(RM) $(DESTDIR)/usr/bin/$(PROGRAM)
	$(RM) $(DESTDIR)/etc/$(PROGRAM).toml
	$(RM) -r $(DESTDIR)/usr/share/$(PROGRAM)

help:   ##               Prints this help message
	@echo ""
	@echo ""
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | fgrep -v : | sed -e 's/\\$$//' | sed -e 's/.*##//'
	@echo ""
	@echo "BUILD TARGETS:"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | fgrep : | sed -e 's/\\$$//' | sed -e 's/:.*##/:/'
	@echo ""
	@echo ""
