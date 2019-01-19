# Build
FROM golang:alpine AS build

RUN apk add --no-cache -U add git make build-base

WORKDIR /src/irccat
COPY . /src/irccat
RUN make build install

# Runtime
FROM alpine:latest

COPY --from=build /go/bin/irccat /irccat

ENTRYPOINT ["/irccat"]
CMD []
