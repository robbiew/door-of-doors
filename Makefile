BINARY_NAME=dod
DIR=bin
TARGET_DIR=/mystic/doors/dod

# BUILDARCH is the host architecture
# ARCH is the target architecture
# we need to keep track of them separately
BUILDARCH ?= $(shell uname -m)


build:

# Host is 64 bit, build 64 and 32 versions
ifeq ($(BUILDARCH),x86_64)
	GOARCH=386 GOOS=linux go build -o ${DIR}/${BINARY_NAME}-linux-386
	GOARCH=amd64 GOOS=linux go build -o ${DIR}/${BINARY_NAME}-linux-amd64
endif

# Host is 32 bit
ifeq ($(BUILDARCH),i386)
	GOARCH=386 GOOS=linux go build -o ${DIR}/${BINARY_NAME}-linux-386
endif

# Host is armv7
ifeq ($(BUILDARCH),arm7l)
	GOARCH=arm GOOS=linux GOARM=7 go build -o ${DIR}/${BINARY_NAME}-linux-armv7
endif

clean:
	go clean
	rm ${DIR}/${BINARY_NAME}-linux-386
	rm ${DIR}/${BINARY_NAME}-linux-amd64