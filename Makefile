##
#  Variables
##

BIN  = grogger
REPO = github.com/ericyoungberg/$(BIN)

GO 		:= packr
BUILD   := build/$(BIN)
DESTDIR := /usr/local/bin

LDFLAGS = 


##
# Targets
##

all: clean $(BUILD)

# Build the binary
$(BUILD): *.go
	@echo "+ $@"
	$(GO) build -v -o "${BUILD}"

# Build and run the binary
.PHONY: run
run: $(BUILD)
	@echo "+ $@"
	@$(BUILD)

# Install the built binary
.PHONY: install
install: $(BUILD)
	@echo "+ $@"
	@mkdir -p "${DESTDIR}/${BIN}"
	cp $(BUILD) "${DESTDIR}/${BIN}"

# Cleanup previous builds
.PHONY: clean
clean: $(BUILD)
	@rm -rf $(BUILD)
