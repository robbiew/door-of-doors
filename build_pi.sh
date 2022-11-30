
#!/bin/bash

# example script to building ARM32 pi (v7) executable and uploading using key

TARGET_USER=pi
TARGET_HOST=raspberrypi
TARGET_DIR=/mystic/doors/dod
ARM_VERSION=6
DIR=bin
BIN=dod-linux-armv$ARM_VERSION

 
echo "Building for Raspberry Pi..."
env CC=arm-linux-gnueabihf-gcc CXX=arm-linux-gnueabihf-g++ \
    CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=$ARM_VERSION \
    go build -v -o $DIR/$BIN

# env GOOS=linux GOARCH=arm GOARM=$ARM_VERSION go build -o $DIR/$BIN
 
echo "Uploading to Raspberry Pi..."
scp -i ~/.ssh/id_rsa $DIR/$BIN $TARGET_USER@$TARGET_HOST:$TARGET_DIR/$BIN