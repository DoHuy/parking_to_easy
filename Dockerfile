# An example Dockerfile for building the engine manager binary into an ubuntu image (local dev version)
FROM golang:1.13.5-alpine AS builder
#RUN mkdir -p /go/src/github.com/DoHuy/parking-to-easy
ADD . /go/src/github.com/DoHuy/parking-to-easy
#RUN cd /go/src/github.com/DoHuy/parking-to-easy
#RUN ls -la /go/src/github.com/DoHuy/parking-to-easy
RUN apk update && \
    apk add -U build-base git curl libstdc++ ca-certificates && \
    cd /go/src/github.com/DoHuy/parking-to-easy && \
#    go install -v github.com/go-playground/validator/v10 &&\
#    go get -u github.com/gin-gonic/gin && \
    go install -v math/bits && \
    go env && go list all | grep cover && \
    GOPATH=/go make docker

##RUN  go install -v github.com/go-playground/validator/v10
FROM alpine:latest

RUN mkdir /app
WORKDIR /app

#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/DoHuy/parking-to-easy/  ./
RUN ls -la
RUN chmod +x /app/parking_service.linux
EXPOSE 8085
ENTRYPOINT ["/app/parking_service.linux"]
