FROM golang

RUN go get -u golang.org/x/net/context
RUN go get -u google.golang.org/grpc
RUN go get -u google.golang.org/grpc/reflection

RUN go install github.com/t0ksss/calculator/server

ENTRYPOINT ["/go/bin/server"]

EXPOSE 50050
