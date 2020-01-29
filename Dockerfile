FROM library/golang:1.13 as builder

ARG SOURCE=/go/src/github.com/isaactl/HelloWorld

ADD ./ ${SOURCE}
WORKDIR ${SOURCE}

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /go/bin/HelloWorld github.com/isaactl/HelloWorld

FROM alpine:3.5

MAINTAINER Tan Liang

LABEL Description="Hello World"

RUN apk update && \
    apk upgrade && \
    apk add \
        bash \
        ca-certificates \
    && rm -rf /var/cache/apk/*

COPY --from=builder /go/bin/HelloWorld /usr/local/bin/HelloWorld

ENTRYPOINT ["/usr/local/bin/HelloWorld"]
