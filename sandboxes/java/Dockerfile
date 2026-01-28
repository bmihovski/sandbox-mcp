FROM alpine:3.23.3

RUN apk add --no-cache --repository=https://dl-cdn.alpinelinux.org/alpine/edge/testing openjdk25
RUN adduser --home /sandbox --disabled-password sandbox

USER sandbox
WORKDIR /sandbox
