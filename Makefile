# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: gpay android ios gpay-cross evm all test clean
.PHONY: gpay-linux gpay-linux-386 gpay-linux-amd64 gpay-linux-mips64 gpay-linux-mips64le
.PHONY: gpay-linux-arm gpay-linux-arm-5 gpay-linux-arm-6 gpay-linux-arm-7 gpay-linux-arm64
.PHONY: gpay-darwin gpay-darwin-386 gpay-darwin-amd64
.PHONY: gpay-windows gpay-windows-386 gpay-windows-amd64

GOBIN = ./build/bin
GO ?= latest
GORUN = env GO111MODULE=on go run

gpay:
	$(GORUN) build/ci.go install ./cmd/gpay
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gpay\" to launch gpay."

all:
	$(GORUN) build/ci.go install

android:
	$(GORUN) build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gpay.aar\" to use the library."
	@echo "Import \"$(GOBIN)/gpay-sources.jar\" to add javadocs"
	@echo "For more info see https://stackoverflow.com/questions/20994336/android-studio-how-to-attach-javadoc"

ios:
	$(GORUN) build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Gpay.framework\" to use the library."

test: all
	$(GORUN) build/ci.go test

lint: ## Run linters.
	$(GORUN) build/ci.go lint

clean:
	env GO111MODULE=on go clean -cache
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go install golang.org/x/tools/cmd/stringer@latest
	env GOBIN= go install github.com/kevinburke/go-bindata/go-bindata@latest
	env GOBIN= go install github.com/fjl/gencodec@latest
	env GOBIN= go install github.com/golang/protobuf/protoc-gen-go@latest
	env GOBIN= go install ./cmd/abigen
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

# Cross Compilation Targets (xgo)

gpay-cross: gpay-linux gpay-darwin gpay-windows gpay-android gpay-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gpay-*

gpay-linux: gpay-linux-386 gpay-linux-amd64 gpay-linux-arm gpay-linux-mips64 gpay-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-*

gpay-linux-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/gpay
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep 386

gpay-linux-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/gpay
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep amd64

gpay-linux-arm: gpay-linux-arm-5 gpay-linux-arm-6 gpay-linux-arm-7 gpay-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep arm

gpay-linux-arm-5:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/gpay
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep arm-5

gpay-linux-arm-6:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/gpay
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep arm-6

gpay-linux-arm-7:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/gpay
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep arm-7

gpay-linux-arm64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/gpay
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep arm64

gpay-linux-mips:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/gpay
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep mips

gpay-linux-mipsle:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/gpay
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep mipsle

gpay-linux-mips64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/gpay
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep mips64

gpay-linux-mips64le:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/gpay
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep mips64le

gpay-darwin: gpay-darwin-386 gpay-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gpay-darwin-*

gpay-darwin-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/gpay
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-darwin-* | grep 386

gpay-darwin-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/gpay
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-darwin-* | grep amd64

gpay-windows: gpay-windows-386 gpay-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gpay-windows-*

gpay-windows-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/gpay
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-windows-* | grep 386

gpay-windows-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/gpay
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-windows-* | grep amd64