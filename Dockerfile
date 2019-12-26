# An example Dockerfile for building the engine manager binary into an ubuntu image (local dev version)
FROM golang:1.10.0-alpine3.7 AS builder
ARG GITHUB_ACCESS_TOKEN
ADD . /go/src/github.com/veritone/edge-stream-engine-manager
RUN apk update && \
    apk add -U build-base git curl libstdc++ ca-certificates && \
    cd /go/src/github.com/veritone/edge-stream-engine-manager && \
    git config --global url."https://${GITHUB_ACCESS_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/" && \
    go env && go list all | grep cover && \
    GOPATH=/go make docker

FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/veritone/edge-stream-engine-manager/service.linux /app/app
RUN chmod +x /app/app

ENTRYPOINT ["/app/app"]
