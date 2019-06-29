#-- Define 
BUILD_DIR = build
PKG 	  = github.com/ericyoungberg/grogger
ARCH 	 ?= darwin/amd64
BIN       = grogger

BUILD_CONTAINER = "${BIN}-builder"

# Set default compiler
GO := go


#-- Generate flags
go_ld_flags = -ldflags "-w -X ${PKG}/$(1)/version.VERSION=$(shell cat VERSION)"
go_ld_flags_static = -ldflags "-w -X ${PKG}/$(1)/version.VERSION=$(shell cat VERSION) -extldflags -static"

GOOSARCHES = darwin/amd64 darwin/386 freebsd/amd64 freebsd/386 linux/arm linux/arm64 linux/amd64 linux/386 solaris/amd64 windows/amd64 windows/386


#-- Build 
all: clean test build install

define build
@${GO} build $(call go_ld_flags,${1}) -o ${BUILD_DIR}/${1} ./${1}
endef

build: $(BIN)/*.go
		@echo "+ $@"
		$(call build,${BIN})

.PHONY: with-docker
with-docker:
		@echo "+ $@"
		docker build -t $(BUILD_CONTAINER) .
		docker run --rm -t -v $(PWD):/go/src/$(PKG) -e "ARCH=$(ARCH)" $(BUILD_CONTAINER)
		docker rmi $(BUILD_CONTAINER)
		@chown -R $(whomai):$(whoami) $(BUILD_DIR)

define buildrelease
GOOS=$(3) GOARCH=$(4) CGO_ENABLED=0 go build \
	 -o $(BUILD_DIR)/$(1)-$(3)-$(4) \
	 -a -tags "static_build netgo" \
	 -installsuffix netgo ${GO_LDFLAGS_STATIC} ./$(2);
md5sum $(BUILD_DIR)/$(1)-$(3)-$(4) > $(BUILD_DIR)/$(1)-$(3)-$(4).md5;
sha256sum $(BUILD_DIR)/$(1)-$(3)-$(4) > $(BUILD_DIR)/$(1)-$(3)-$(4).sha256;
endef

.PHONY: static
static:
		@echo "+ $@"
		$(call buildrelease,$(PRQL_BIN),$(PRQL_DIR),$(subst /,,$(dir $(ARCH))),$(notdir $(ARCH)))
		$(call buildrelease,$(PRQLD_BIN),$(PRQLD_DIR),$(subst /,,$(dir $(ARCH))),$(notdir $(ARCH)))

.PHONY: test
test:
		@echo "+ $@"
		@echo "No tests exist for grogger"

.PHONY: staticcheck
staticcheck:
		@echo "+ $@"
		@staticcheck $(shell go list ./... | grep -v vendor) | grep -v '.pb.go:' | tee /dev/stderr

.PHONY: install
install:
		@echo "+ $@"
		@cp $(BUILD_DIR)/* $(GOPATH)/bin

.PHONY: clean
clean:
		@echo "+ $@"
		rm -rf $(BUILD_DIR)
		rm -rf $(GOPATH)/bin/$(PRQL_BIN) $(GOPATH)/bin/$(PRQLD_BIN)
