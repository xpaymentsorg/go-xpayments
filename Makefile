.PHONY: gpay gpay-cross evm all test clean
.PHONY: gpay-linux gpay-linux-386 gpay-linux-amd64 gpay-linux-mips64 gpay-linux-mips64le
.PHONY: gpay-darwin gpay-darwin-386 gpay-darwin-amd64

GOBIN = $(shell pwd)/build/bin
GOFMT = gofmt
GO ?= latest
GO_PACKAGES = .
GO_FILES := $(shell find $(shell go list -f '{{.Dir}}' $(GO_PACKAGES)) -name \*.go)

GIT = git

gpay:
	build/env.sh go run build/ci.go install ./cmd/gpay
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gpay\" to launch gpay."

gc:
	build/env.sh go run build/ci.go install ./cmd/gc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gc\" to launch gc."

bootnode:
	build/env.sh go run build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch a bootnode."

puppeth:
	build/env.sh go run build/ci.go install ./cmd/puppeth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/puppeth\" to launch puppeth."

all:
	build/env.sh go run build/ci.go install

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# Cross Compilation Targets (xgo)

gpay-cross: gpay-linux gpay-darwin
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gpay-*

gpay-linux: gpay-linux-386 gpay-linux-amd64 gpay-linux-mips64 gpay-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-*

gpay-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/gpay
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep 386

gpay-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/gpay
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep amd64

gpay-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/gpay
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep mips

gpay-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/gpay
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep mipsle

gpay-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/gpay
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep mips64

gpay-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/gpay
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep mips64le

gpay-darwin: gpay-darwin-386 gpay-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gpay-darwin-*

gpay-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/gpay
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-darwin-* | grep 386

gpay-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/gpay
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-darwin-* | grep amd64

gofmt:
	$(GOFMT) -s -w $(GO_FILES)
	$(GIT) checkout vendor
