#!/bin/bash -ex

go install ./cmd/...

go generate ./...
go get -t -v -u ./...
go test ./...
