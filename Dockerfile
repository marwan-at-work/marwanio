FROM golang:1.8.3 AS builder

RUN mkdir -p /go/src/github.com/marwan-at-work/marwanio && \
    go get -u github.com/gopherjs/gopherjs

COPY . /go/src/github.com/marwan-at-work/marwanio

WORKDIR /go/src/github.com/marwan-at-work/marwanio

RUN CGO_ENABLED=0 go build -a -ldflags '-s' && \
    cd /go/src/github.com/marwan-at-work/marwanio && \
    gopherjs build .

FROM busybox

RUN mkdir -p /go/src/github.com/marwan-at-work/marwanio/frontend && \
    mkdir -p /go/src/github.com/marwan-at-work/marwanio/blog/posts && \
    mkdir -p /go/src/github.com/marwan-at-work/marwanio/public

WORKDIR /go/src/github.com/marwan-at-work/marwanio

COPY --from=builder /go/src/github.com/marwan-at-work/marwanio/marwanio /go/src/github.com/marwan-at-work/marwanio

COPY --from=builder /go/src/github.com/marwan-at-work/marwanio/frontend/frontend.js /go/src/github.com/marwan-at-work/marwanio/frontend

COPY --from=builder /go/src/github.com/marwan-at-work/marwanio/frontend/frontend.js.map /go/src/github.com/marwan-at-work/marwanio/frontend

COPY --from=builder /go/src/github.com/marwan-at-work/marwanio/public /go/src/github.com/marwan-at-work/marwanio/public

COPY --from=builder /go/src/github.com/marwan-at-work/marwanio/blog/posts /go/src/github.com/marwan-at-work/marwanio/blog/posts

ENV GO_MODE=production

ENTRYPOINT ["./marwanio"]