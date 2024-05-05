BIN_FILE=cracklepop
BUILD_DIR=build

all: test build

build:
	go build -o ${BUILD_DIR}/${BIN_FILE}
test:
	go test -v ./...
clean:
	rm -rf ${BUILD_DIR}