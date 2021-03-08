#!/usr/bin/env bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

cd "$DIR"

rm -rf bin/*
rm -rf pkg/*
sleep 5

echo "Creating executables per platform"
GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/cnsbackup cmd/cnsbackup/main.go
GOOS=windows GOARCH=amd64 go build -o bin/windows/amd64/cnsbackup.exe cmd/cnsbackup/main.go
GOOS=darwin GOARCH=amd64 go build -o bin/darwin/amd64/cnsbackup cmd/cnsbackup/main.go