FROM alpine:3.23.3

RUN apk add --no-cache g++ make cmake
RUN adduser --home /sandbox --disabled-password sandbox

USER sandbox
WORKDIR /sandbox
