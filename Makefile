PROJECT := aiisx
REPO    := aiisx.com
HASH    := $(shell git rev-parse --short HEAD)
DATE    := $(shell date)
TAG     := $(shell git describe --tags --always --abbrev=0 --match="v[0-9]*.[0-9]*.[0-9]*" 2> /dev/null)
VERSION := $(shell echo "${TAG}" | sed 's/^.//')


LDFLAGS_DEV     := -ldflags "-X '${REPO}/src/version.CommitHash=${HASH}' -X '${REPO}/src/version.CompileDate=${DATE}' -X '${REPO}/src/version.CompileDate=${DATE}'"
LDFLAGS_RELEASE := -ldflags "-X '${REPO}/src/version.Version=${VERSION}' -X '${REPO}/src/version.CommitHash=${HASH}' -X '${REPO}/src/version.CompileDate=${DATE}'"


node-fetch:
	command -v pnpm >/dev/null >&2 || npm install \
		--no-audit \
		--no-fund \
		--quiet \
		--global pnpm
	cd src/public && \
		yarn install


node-debug:
	cd src/public && \
		yarn dev


.PHONY:
node-build:
	cd src/public && \
		yarn build

.PHONY:
node-preview:
	cd src/public && \
		yarn preview

.PHONY: build
build:
	#
	# ################################################################################
	# >>> TARGET: build
	# ################################################################################
	#
	go build ${LDFLAGS_RELEASE}

.PHONY: run
run:
	#
	# ################################################################################
	# >>> TARGET: run
	# ################################################################################
	#
	go run main.go ${LDFLAGS_DEV}
