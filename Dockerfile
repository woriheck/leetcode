FROM golang:1.13.4-alpine3.10

RUN apk add curl git build-base && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh 

WORKDIR /go/src/leetcode
RUN dep ensure