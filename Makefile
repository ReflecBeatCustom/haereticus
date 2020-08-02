###########################################
#    THIS MAKEFILE IS USED FOR GO 1.13    #
#    MAYBE NOT WORKING ON OTHER GO VER    #
###########################################
BIN_DIR=_output/cmd/bin
REPO_PATH="github.com/ReflecBeatCustom/haereticus"
REL_OSARCH="linux/amd64"
GitSHA=`git rev-parse HEAD`
Date=`date "+%Y-%m-%d %H:%M:%S"`
RELEASE_VERSION=$(shell git describe --tags --always --dirty)
IMG_BUILDER=docker
LD_FLAGS=" \
    -X '${REPO_PATH}/pkg/version.GitSHA=${GitSHA}' \
    -X '${REPO_PATH}/pkg/version.Built=${Date}'   \
    -X '${REPO_PATH}/pkg/version.Version=${RELEASE_VERSION}'"
build: all
all: init haereticus
haereticus:
	go build -ldflags ${LD_FLAGS} -o ${BIN_DIR}/haereticus ./cmd/
init:
	mkdir -p ${BIN_DIR}
clean:
	rm -fr ${BIN_DIR}
images:
	@echo "version: ${RELEASE_VERSION}"
	${IMG_BUILDER} build -t csighub.tencentyun.com/elihe/haereticus:${RELEASE_VERSION} .
.PHONY: clean