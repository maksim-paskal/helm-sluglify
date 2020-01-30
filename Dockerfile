FROM golang:1.12 as build

COPY main.go /usr/src/helm-sluglify/main.go
COPY go.mod /usr/src/helm-sluglify/go.mod
COPY go.sum /usr/src/helm-sluglify/go.sum

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN cd /usr/src/helm-sluglify \
  && go mod download \
  && go mod verify \
  && go build -v -o helm-sluglify -ldflags "-X main.buildTime=$(date +"%Y%m%d%H%M%S")"

FROM alpine:latest

RUN apk add --no-cache --virtual .build-deps curl tar \
  && curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.14.9/bin/linux/amd64/kubectl \
  && chmod +x ./kubectl \
  && mv ./kubectl /usr/local/bin/kubectl \
  && curl -LO https://get.helm.sh/helm-v2.14.1-linux-amd64.tar.gz \
  && tar -zxvf helm-v2.14.1-linux-amd64.tar.gz \
  && rm -rf helm-v2.14.1-linux-amd64.tar.gz \
  && mv linux-amd64/helm /usr/local/bin/helm \
  && rm -rf linux-amd64 \
  && apk del .build-deps \
  && apk add --no-cache docker \
  && rm -rf /tmp/* /var/cache/apk/*

COPY --from=build /usr/src/helm-sluglify/helm-sluglify /usr/local/bin/helm-sluglify

CMD helm-sluglify
