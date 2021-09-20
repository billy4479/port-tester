GO ?= go

all:
	mkdir -p build
	$(GO) build -o build

windows:
	mkdir -p build
	GOOS=windows $(GO) build -o build

.PHONY: all windows