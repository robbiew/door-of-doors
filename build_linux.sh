#!/bin/bash

# example script to build linux 386 & amd64 executable on amd64.
# after building, it copies them over the to BBS door directory.

DIR=bin
BIN=dod-linux
TARGET_DIR=/mystic/doors/dod
 
echo "Building for Linux 386..."
env GOOS=linux GOARCH=386 go build -o $DIR/$BIN-386
cp $DIR/$BIN-386 /bbs/doors/dod
 
echo "Building for Linux amd64..."
env GOOS=linux GOARCH=amd64 go build -o $DIR/$BIN-amd64
cp $DIR/$BIN-amd64 /bbs/doors/dod
