FROM golang:1.7.1-wheezy

RUN mkdir -p /go/src/github.com/skycoin/skyhash
WORKDIR /go/src/github.com/skycoin/skyhash

ADD . /go/src/github.com/skycoin/skyhash

RUN cd /go/src/github.com/skycoin/skyhash && \
    go get github.com/skycoin/skyhash/...  && \
    go build -o skyhashd

ENTRYPOINT ./skyhashd --launch-browser=false
