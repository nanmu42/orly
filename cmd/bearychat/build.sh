#!/usr/bin/env bash

VERSION=`git describe --tags --dirty`
BUILD=`date +%FT%T%z`

go build -ldflags "-s -w -X main.Version=$VERSION -X main.BuildDate=$BUILD" -o bearychat