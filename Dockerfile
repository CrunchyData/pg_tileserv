# copy build result to a centos base image to match other
# crunchy containers
FROM centos:7
RUN mkdir /app
ADD ./pg_tileserv /app/
ADD ./assets /app/assets

ARG VERSION

LABEL vendor="Crunchy Data" \
	url="https://crunchydata.com" \
	release="${VERSION}" \
	org.opencontainers.image.vendor="Crunchy Data" \
	os.version="7.7"

VOLUME ["/config"]

USER 1001
EXPOSE 7800

WORKDIR /app
ENTRYPOINT ["/app/pg_tileserv"]
CMD []

# To build
# make APPVERSION=1.0.2 clean build build-docker

# To build using binaries from golang docker image
# make APPVERSION=1.0.2 clean bin-docker build-docker

# To run
# docker run -dt -e DATABASE_URL=postgres://user:pass@host/dbname -p 7800:7800 pramsey/pg_tileserv:1.0.2
