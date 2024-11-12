# syntax=docker/dockerfile:1

FROM golang:1.23 AS build

WORKDIR /mod
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

RUN mkdir /app

# Build the application.
# Leverage a cache mount to /go/pkg/mod/ to speed up subsequent builds.
# Leverage a bind mount to the current directory to avoid having to copy the
# source code into the container.
# Leverage a cache mount to GOCACHE to speed up subsequent builds.
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    GOARCH=amd64 CGO_ENABLED=0 go build -o /app .

FROM --platform=linux/amd64 scratch

WORKDIR /app

COPY --from=build /app/marwanio /app/marwanio

EXPOSE 3091

ENTRYPOINT ["/app/marwanio"]
