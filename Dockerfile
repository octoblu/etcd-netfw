FROM golang:1.4.2
MAINTAINER Octoblu Inc. <docker@octoblu.com>

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

RUN go get github.com/mattn/gom

RUN mkdir -p /usr/local/go/src/github.com/octoblu/etcd-netfw
WORKDIR /usr/local/go/src/github.com/octoblu/etcd-netfw

COPY Gomfile /usr/local/go/src/github.com/octoblu/etcd-netfw/
RUN gom install

COPY . /usr/local/go/src/github.com/octoblu/etcd-netfw
RUN gom test
RUN gom build

# Expose default listening amb port
EXPOSE 1337

ENTRYPOINT ["/usr/local/go/src/github.com/octoblu/etcd-netfw/etcd-netfw"]
