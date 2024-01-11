ARG GOLANG_VERSION
ARG TARGETARCH
ARG VERSION
ARG BASE_REGISTRY
ARG BASE_IMAGE
ARG PLATFORM
FROM --platform=${PLATFORM} golang:${GOLANG_VERSION}-alpine AS builder
LABEL stage=tileservbuilder

ARG TARGETARCH
ARG VERSION

WORKDIR /app
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -v -ldflags "-s -w -X main.programVersion=${VERSION}"

FROM --platform=${TARGETARCH} ${BASE_REGISTRY}/${BASE_IMAGE} AS inherited

COPY --from=builder /app/pg_tileserv /app/
COPY --from=builder /app/assets /app/assets

VOLUME ["/config"]

USER 1001
EXPOSE 7800

WORKDIR /app
ENTRYPOINT ["/app/pg_tileserv"]
CMD []

FROM --platform=${PLATFORM} ${BASE_REGISTRY}/${BASE_IMAGE} AS local

RUN mkdir /app
ADD ./pg_tileserv /app/
ADD ./assets /app/assets

VOLUME ["/config"]

USER 1001
EXPOSE 7800

WORKDIR /app
ENTRYPOINT ["/app/pg_tileserv"]
CMD []

# To build
# make APPVERSION=1.0.2 clean local-docker

# To build using binaries from golang docker image
# make APPVERSION=1.0.2 clean docker

# To run
# docker run -dt -e DATABASE_URL=postgres://user:pass@host/dbname -p 7800:7800 pramsey/pg_tileserv:1.0.2
