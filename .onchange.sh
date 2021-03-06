#!/bin/bash -e

path=$1
dir=$(dirname "$path")
filename=$(basename "$path")
extension="${filename##*.}"
nameonly="${filename%.*}"

case $extension in
    go)
	goimports -w $path
        gofmt -w $path
	go vet
        ;;
esac

case $filename in
    token.go)
	go generate
	;;
esac

go install github.com/gregoryv/go-timesheet/...
go test -coverprofile /tmp/c.out ./...
uncover /tmp/c.out
