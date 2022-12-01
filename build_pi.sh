
#!/bin/bash

# example script to build armv6 executable on a Raspberry Pi.
# after building, it copies it to the BBS door directory.

TARGET_DIR=/mystic/doors/dod
ARM_VERSION=6
DIR=bin
BIN=dod-linux-armv$ARM_VERSION

echo "Building for Raspberry Pi..."
env GOOS=linux GOARCH=arm GOARM=$ARM_VERSION \
    go build -v -o $DIR/$BIN
echo "Done!"

echo "Copying to Mystic Door diectory..."
cp $DIR/$BIN $TARGET_DIR
echo "Done!"