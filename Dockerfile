FROM golang:alpine
MAINTAINER Eric Youngberg <eric@lmtlss.net>

ENV GOPATH /go
ENV ARCH=linux/amd64

WORKDIR /opt

RUN set -e; \
    \
    apk add --no-cache \
        git \
        curl \
        make \
    ; \
    echo 'make static ARCH=${ARCH}' >> entrypoint.sh \
    && chmod +x entrypoint.sh

WORKDIR /go/src/github.com/ericyoungberg/grogger

ENTRYPOINT ["sh", "/opt/entrypoint.sh"]
