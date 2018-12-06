FROM golang as godir

RUN mkdir -p /go/src/hello

ADD . /go/src/hello
WORKDIR /go/src/hello

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

FROM canthefason/go-watcher:latest

COPY --from=godir /go/src/hello /go/src/hello


