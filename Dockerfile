FROM alpine:3.5

MAINTAINER Tan Liang

LABEL Description="Hello World"

RUN apk update && \
    apk upgrade && \
    apk add \
        bash \
        ca-certificates \
    && rm -rf /var/cache/apk/*

COPY HelloWorld /usr/local/bin/HelloWorld

ENTRYPOINT ["/usr/local/bin/HelloWorld"]
