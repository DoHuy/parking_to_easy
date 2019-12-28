# An example Dockerfile for building the engine manager binary into an ubuntu image (local dev version)
FROM golang:1.10.0-alpine3.7 AS builder
ADD . /go/src/github.com/DoHuy/parking-to-easy
RUN apk update && \
    apk add -U build-base git curl libstdc++ ca-certificates && \
    cd /go/src/github.com/DoHuy/parking-to-easy && \
    go get -u github.com/gin-gonic/gin && \
    go install -v math/bits && \
    go env && go list all | grep cover && \
    GOPATH=/go make docker

FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/DoHuy/parking-to-easy/parking_service.linux /app/app
RUN chmod +x /app/app
EXPOSE 8085
ENTRYPOINT ["/app/app"]
