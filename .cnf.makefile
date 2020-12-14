uname_S := $(shell sh -c 'uname -s 2>/dev/null || echo ""')
uname_M := $(shell sh -c 'uname -m 2>/dev/null || echo ""')

ifeq ($(uname_M),x86_64)
    ARCH=amd64
endif

ifeq ($(uname_S),Linux)
    ARCH=$(shell sh -c 'dpkg --print-architecture 2>/dev/null || echo not')
endif

DOCKERFILE := Dockerfile
APPVERSION := latest
REPO := pramsey
TARGET_GOAL := $(firstword $(subst " ", ,$(wordlist 1,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))))
BUILDPLATFORM := linux/amd64,linux/arm64

ifeq ($(DOCKERFILE),Dockerfile.alpine)
    APPVERSION := $(APPVERSION)-alpine-3.12
endif
