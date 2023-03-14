PROJECT_PKG = github.com/talgat065/shortmate

VERSION ?=$(shell git describe --tags --exact-match 2>/dev/null || git symbolic-ref -q --short HEAD)
COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE ?= $(shell date +%FT%T%z)
BUILD_INFO_PATH = internal/health

LDFLAGS += -X ${PROJECT_PKG}/${BUILD_INFO_PATH}.Version=${VERSION}
LDFLAGS += -X ${PROJECT_PKG}/${BUILD_INFO_PATH}.CommitHash=${COMMIT_HASH}
LDFLAGS += -X ${PROJECT_PKG}/${BUILD_INFO_PATH}.BuildDate=${BUILD_DATE}

BUILD_DIR = bin/web

.DEFAULT_GOAL=all


.PHONY: build
.PHONY: test

all:
	make build
	make serve

test:
	go test -v -race ./...

build:
	go build ${GOARGS} -ldflags "${LDFLAGS}" -o ${BUILD_DIR} ./cmd/web

build/dev:
	sudo docker build -t notionassistant .

serve:
	${BUILD_DIR}

serve/dev:
	sudo docker stop notionassistant \
	sudo docker run --rm -d -it -t notionassistant --name notionassistant -p 9000:9000 notionassistant
