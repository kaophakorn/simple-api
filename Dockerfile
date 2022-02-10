FROM golang:1.16-alpine

RUN apk add --no-cache nginx git curl && apk add busybox-extras
RUN apk upgrade && apk upgrade
RUN apk add --no-cache nginx
RUN apk add --no-cache git
RUN apk add --no-cache curl
RUN apk add --no-cache busybox-extras
ENV http_proxy=
ENV https_proxy=
ENV no_proxy=
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
EXPOSE 5000