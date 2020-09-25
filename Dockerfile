FROM golang:1.14 as build

COPY *.go /usr/src/sluglify/
COPY go.* /usr/src/sluglify/

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GOFLAGS="-trimpath"

RUN cd /usr/src/sluglify \
  && go mod download \
  && go mod verify \
  && go build -v -o sluglify -ldflags "-X main.buildTime=$(date +"%Y%m%d%H%M%S")"

FROM alpine:latest

COPY --from=build /usr/src/sluglify/sluglify /usr/local/bin/sluglify

CMD /usr/local/bin/sluglify
