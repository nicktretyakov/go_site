ifeq ($(BINARY_PATH),)
BINARY_PATH := dist
endif

empty:

clean:
	rm -rf ./${BINARY_PATH}/*


build-website:cmd/website/main/*.go
	GO111MODULE=on go build -o ${BINARY_PATH}/website cmd/website/main/*.go

build: build-website


build-website-st:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o ${BINARY_PATH}/website_static cmd/website/main/*.go

build-st: build-website-st


.PHONY: fmt
fmt:
	gofmt -w -s .
