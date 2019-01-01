DEV_DEPS = github.com/mailru/easyjson/...

NAME = rbheap
BINARY = bin/${NAME}
VERSION ?= $(shell cat VERSION)

SOURCES = $(shell find . -name '*.go' -o -name 'Makefile' -o -name 'VERSION')

all: bootstrap generate build
bootstrap: dev-deps

$(BINARY): $(SOURCES)
	CGO_ENABLED=0 go build -a -o ${BINARY} -ldflags \ "\
		-s -w \
		-X main.version=${VERSION} \
		-X main.commit=$(shell git show --format="%h" --no-patch) \
		-X main.date=$(shell date +%Y-%m-%dT%T%z)"

.PHONY: build
build: $(BINARY)

.PHONY: clean
clean:
	$(eval BIN_DIR := $(shell dirname ${BINARY}))
	if [ -f ${BINARY} ]; then rm ${BINARY}; fi
	if [ -d ${BIN_DIR} ]; then rmdir ${BIN_DIR}; fi

.PHONY: test
test:
	go test ./...

.PHONY: generate
generate:
	go generate ./...

%_easyjson.go: %.go
	easyjson -all $^

.PHONY: dev-deps
dev-deps:
	@$(foreach DEP,$(DEV_DEPS),go get $(DEP);)

.PHONY: update-dev-deps
update-dev-deps:
	@$(foreach DEP,$(DEV_DEPS),go get -u $(DEP);)
