FROM alpine:3.23.3

RUN apk add --no-cache gcc musl-dev make
RUN adduser --home /sandbox --disabled-password sandbox

USER sandbox
WORKDIR /sandbox
