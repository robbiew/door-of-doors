BINARY_NAME=dod
DIR=bin

# BUILDARCH is the host architecture
BUILDARCH ?= $(shell uname -m)

build:
# Host is 64 bit linux, build 64 and 32 versions
ifeq ($(BUILDARCH),x86_64)
	GOARCH=386 GOOS=linux go build -o ${DIR}/${BINARY_NAME}-linux-386
	GOARCH=amd64 GOOS=linux go build -o ${DIR}/${BINARY_NAME}-linux-amd64
endif

# Host is 32 bit linux
ifeq ($(BUILDARCH),i386)
	GOARCH=386 GOOS=linux go build -o ${DIR}/${BINARY_NAME}-linux-386
endif

# Host is 32 bit linux armv7
ifeq ($(BUILDARCH),armv7l)
	GOARCH=arm GOOS=linux GOARM=7 go build -o ${DIR}/${BINARY_NAME}-linux-armv7
	GOARCH=arm GOOS=linux GOARM=6 go build -o ${DIR}/${BINARY_NAME}-linux-armv6
endif

# Host is 64 bit linux armv8
ifeq ($(BUILDARCH),armv8)
	GOARCH=arm GOOS=linux GOARM=8 go build -o ${DIR}/${BINARY_NAME}-linux-armv8
	GOARCH=arm GOOS=linux GOARM=7 go build -o ${DIR}/${BINARY_NAME}-linux-armv7
	GOARCH=arm GOOS=linux GOARM=6 go build -o ${DIR}/${BINARY_NAME}-linux-armv6
endif

clean:
	go clean
	rm ${DIR}/${BINARY_NAME}-linux-386
	rm ${DIR}/${BINARY_NAME}-linux-amd64

# change or remove. I used this to automate testing.
copy:
ifeq ($(BUILDARCH),i386)
	cp ${DIR}/${BINARY_NAME}-linux-386 /bbs/doors/dod
endif
ifeq ($(BUILDARCH),x86_64)
	cp ${DIR}/${BINARY_NAME}-linux-386 /bbs/doors/dod
	cp ${DIR}/${BINARY_NAME}-linux-amd64 /bbs/doors/dod
endif
ifeq ($(BUILDARCH),armv7l)
	cp ${DIR}/${BINARY_NAME}-linux-armv7 /mystic/doors/dod
endif

build_and_copy: build copy