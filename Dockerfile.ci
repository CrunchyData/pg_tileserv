# build the binary using a golang build image
FROM golang:1.13 as buildimage
WORKDIR /go/src/app
COPY . .
RUN go build -v

# copy build result to a centos base image to match other
# crunchy containers
FROM centos:7
RUN mkdir /app
COPY --from=buildimage /go/src/app/pg_tileserv /app/
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
# docker build -f Dockerfile.build --build-arg VERSION=0.1 -t pramsey/pg_tileserv:latest .

# To run
# docker run -dt -e DATABASE_URL=postgres://user:pass@host/dbname -p 7800:7800 pramsey/pg_tileserv
