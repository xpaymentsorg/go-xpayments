.PHONY: gpay android ios gpay-cross swarm evm all test clean
.PHONY: gpay-linux gpay-linux-386 gpay-linux-amd64 gpay-linux-mips64 gpay-linux-mips64le
.PHONY: gpay-linux-arm gpay-linux-arm-5 gpay-linux-arm-6 gpay-linux-arm-7 gpay-linux-arm64
.PHONY: gpay-darwin gpay-darwin-386 gpay-darwin-amd64
.PHONY: gpay-windows gpay-windows-386 gpay-windows-amd64
.PHONY: docker release

GOBIN = $(shell pwd)/build/bin
GO ?= latest

# Compare current go version to minimum required version. Exit with \
# error message if current version is older than required version.
# Set min_ver to the mininum required Go version such as "1.12"
min_ver := 1.12
ver = $(shell go version)
ver2 = $(word 3, ,$(ver))
cur_ver = $(subst go,,$(ver2))
ver_check := $(filter $(min_ver),$(firstword $(sort $(cur_ver) \
$(min_ver))))
ifeq ($(ver_check),)
$(error Running Go version $(cur_ver). Need $(min_ver) or higher. Please upgrade Go version)
endif

gpay:
	cd cmd/gpay; go build -o ../../bin/gpay
	@echo "Done building."
	@echo "Run \"bin/gpay\" to launch gpay."

bootnode:
	cd cmd/bootnode; go build -o ../../bin/gpay-bootnode
	@echo "Done building."
	@echo "Run \"bin/gpay-bootnode\" to launch gpay."

docker:
	docker build -t gpay/gpay .

all: bootnode gpay

release:
	./release.sh

install: all
	cp bin/gpay-bootnode $(GOPATH)/bin/gpay-bootnode
	cp bin/gpay $(GOPATH)/bin/gpay

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gpay.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gpay.framework\" to use the library."

test:
	go test ./...

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/kevinburke/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go get -u github.com/golang/protobuf/protoc-gen-go
	env GOBIN= go install ./cmd/abigen
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
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
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/gpay
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep 386

gpay-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/gpay
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep amd64

gpay-linux-arm: gpay-linux-arm-5 gpay-linux-arm-6 gpay-linux-arm-7 gpay-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep arm

gpay-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/gpay
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep arm-5

gpay-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/gpay
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep arm-6

gpay-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/gpay
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep arm-7

gpay-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/gpay
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-linux-* | grep arm64

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

gpay-windows: gpay-windows-386 gpay-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gpay-windows-*

gpay-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/gpay
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-windows-* | grep 386

gpay-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/gpay
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gpay-windows-* | grep amd64
