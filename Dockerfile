FROM golang:1.16.6-alpine3.14

ENV GOOS linux
ENV GOARCH amd64
ENV PATH $PATH:/go/bin/linux_amd64
ENV CGO_ENABLED 0

RUN \
  echo http://dl-cdn.alpinelinux.org/alpine/edge/community/ >> /etc/apk/repositories && \
  apk update --no-cache && \
  apk add --virtual .build \
  curl \
  bash

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . /app
