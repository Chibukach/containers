FROM golang:alpine

MAINTAINER chibu

ADD ./app /go/src/app

WORKDIR /go/src/app

ENV PORT=8000


