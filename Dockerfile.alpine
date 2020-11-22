# Lightweight Alpine-based pg_tileserv Docker Image
# Author: Just van den Broecke
FROM golang:1.15.5-alpine3.12

# Build ARGS
ARG VERSION="latest-alpine-3.12"

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -v -ldflags "-s -w -X main.programVersion=${VERSION}"

# Multi-stage build: only copy build result and resources
FROM alpine:3.12

LABEL original_developer="Crunchy Data" \
    contributor="Just van den Broecke <justb4@gmail.com>" \
    vendor="Crunchy Data" \
	url="https://crunchydata.com" \
	release="${VERSION}" \
	org.opencontainers.image.vendor="Crunchy Data" \
	os.version="3.12"

RUN apk --no-cache add ca-certificates && mkdir /app
WORKDIR /app/
COPY --from=0 /app/pg_tileserv /app/
COPY --from=0 /app/assets /app/assets

VOLUME ["/config"]

USER 1001
EXPOSE 7800

ENTRYPOINT ["/app/pg_tileserv"]
CMD []

# To build and run specific version
#
# export VERSION="latest-alpine-3.12"
# docker build --build-arg VERSION=${VERSION} -t pramsey/pg_tileserv:${VERSION} -f Dockerfile.alpine
#
# Best is to use another PostGIS Docker Container whoose host is reachable from the pg_tileserv Container.
# docker run -dt -e DATABASE_URL=postgres://user:pass@host/dbname -p 7800:7800 pramsey/pg_tileserv:${VERSION}
#
# See a full example using Docker Compose under examples/docker
#
