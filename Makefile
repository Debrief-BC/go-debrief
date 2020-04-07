# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: debrief android ios debrief-cross evm all test clean
.PHONY: debrief-linux debrief-linux-386 debrief-linux-amd64 debrief-linux-mips64 debrief-linux-mips64le
.PHONY: debrief-linux-arm debrief-linux-arm-5 debrief-linux-arm-6 debrief-linux-arm-7 debrief-linux-arm64
.PHONY: debrief-darwin debrief-darwin-386 debrief-darwin-amd64
.PHONY: debrief-windows debrief-windows-386 debrief-windows-amd64

GOBIN = ./build/bin
GO ?= latest
GORUN = env GO111MODULE=on go run

debrief:
	$(GORUN) build/ci.go install ./cmd/debrief
	@echo "Done building."
	@echo "Run \"$(GOBIN)/debrief\" to launch debrief."

all:
	$(GORUN) build/ci.go install

android:
	$(GORUN) build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/debrief.aar\" to use the library."

ios:
	$(GORUN) build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/debrief.framework\" to use the library."

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
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/kevinburke/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go get -u github.com/golang/protobuf/protoc-gen-go
	env GOBIN= go install ./cmd/abigen
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

# Cross Compilation Targets (xgo)

debrief-cross: debrief-linux debrief-darwin debrief-windows debrief-android debrief-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/debrief-*

debrief-linux: debrief-linux-386 debrief-linux-amd64 debrief-linux-arm debrief-linux-mips64 debrief-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/debrief-linux-*

debrief-linux-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/debrief
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/debrief-linux-* | grep 386

debrief-linux-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/debrief
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/debrief-linux-* | grep amd64

debrief-linux-arm: debrief-linux-arm-5 debrief-linux-arm-6 debrief-linux-arm-7 debrief-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/debrief-linux-* | grep arm

debrief-linux-arm-5:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/debrief
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/debrief-linux-* | grep arm-5

debrief-linux-arm-6:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/debrief
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/debrief-linux-* | grep arm-6

debrief-linux-arm-7:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/debrief
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/debrief-linux-* | grep arm-7

debrief-linux-arm64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/debrief
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/debrief-linux-* | grep arm64

debrief-linux-mips:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/debrief
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/debrief-linux-* | grep mips

debrief-linux-mipsle:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/debrief
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/debrief-linux-* | grep mipsle

debrief-linux-mips64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/debrief
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/debrief-linux-* | grep mips64

debrief-linux-mips64le:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/debrief
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/debrief-linux-* | grep mips64le

debrief-darwin: debrief-darwin-386 debrief-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/debrief-darwin-*

debrief-darwin-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/debrief
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/debrief-darwin-* | grep 386

debrief-darwin-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/debrief
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/debrief-darwin-* | grep amd64

debrief-windows: debrief-windows-386 debrief-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/debrief-windows-*

debrief-windows-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/debrief
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/debrief-windows-* | grep 386

debrief-windows-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/debrief
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/debrief-windows-* | grep amd64
