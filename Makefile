DEV_DEPS = github.com/mailru/easyjson/...

NAME = rbheapleak
BINARY = bin/${NAME}
VERSION ?= $(shell cat VERSION)

#
# Main targets.
#

SOURCES = $(shell find . -name '*.go' -o -name 'Makefile' -o -name 'VERSION')

all: bootstrap generate build
bootstrap: dev-deps dep-ensure

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
	go test

.PHONY: generate
generate: dev-deps
	go generate

#
# EasyJSON targets.
#

EASYJSON_FILES = $(shell find . -name '*_easyjson.go' -not -path '*/vendor/*')

define easyjson_file
$(1): $(subst _easyjson.go,.go,$(1))
	easyjson -all $$^
endef

$(foreach file,$(EASYJSON_FILES),$(eval $(call easyjson_file,$(file))))

#
# Dependency targets.
#

.PHONY: dep-ensure
dep-ensure:
	dep ensure

.PHONY: dev-deps
dev-deps:
	@$(foreach DEP,$(DEV_DEPS),go get $(DEP);)

.PHONY: update-dev-deps
update-dev-deps:
	@$(foreach DEP,$(DEV_DEPS),go get -u $(DEP);)
