FROM       arken/gom-base
MAINTAINER Octoblu Inc. <docker@octoblu.com>

COPY . /usr/local/go/src/github.com/octoblu/etcd-netfw
WORKDIR /usr/local/go/src/github.com/octoblu/etcd-netfw
RUN go get
RUN gom install
RUN gom test
RUN gom build

# Expose default listening amb port
EXPOSE 1337

ENTRYPOINT ["/usr/local/go/src/github.com/octoblu/etcd-netfw/etcd-netfw"]
