#!/bin/sh
go build -v -ldflags "-X main.buildTime=$(date +"%Y%m%d%H%M%S")" -o helm-sluglify && ./helm-sluglify $*