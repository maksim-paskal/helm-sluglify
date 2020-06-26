FROM golang:1.12 as build

COPY *.go /usr/src/helm-sluglify/
COPY go.* /usr/src/helm-sluglify/

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN cd /usr/src/helm-sluglify \
  && go mod download \
  && go mod verify \
  && go build -v -o helm-sluglify -ldflags "-X main.buildTime=$(date +"%Y%m%d%H%M%S")"

FROM alpine:latest

COPY --from=build /usr/src/helm-sluglify/helm-sluglify /usr/local/bin/helm-sluglify

CMD helm-sluglify
