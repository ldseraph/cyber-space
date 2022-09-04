MODULE     = $(shell env GO111MODULE=on $(GO) list -m)
BUILDTIME ?= $(shell date +%FT%T%z)
GITCOMMIT ?= $(shell git rev-parse --short HEAD 2> /dev/null)
GITBRANCH ?= $(shell git symbolic-ref -q --short HEAD 2> /dev/null || echo HEAD)
GITSTATE  ?= $(shell git status --porcelain 2> /dev/null | grep -q -v -e "^$$" && echo dirty || echo clean )
VERSION   ?= $(shell git describe --tags --dirty --match=v* 2> /dev/null || echo v0)
PKGS       = $(or $(PKG),$(shell env GO111MODULE=on $(GO) list ./...))
TESTPKGS   = $(shell env GO111MODULE=on $(GO) list -f \
			        '{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' \
			         $(PKGS))
BIN        = $(CURDIR)/bin
ARCH       = linux/amd64 linux/arm64 windows/amd64
GO         = go
GODOC      = godoc
TIMEOUT    = 15
V          = 0
Q          = $(if $(filter 1,$V),,@)
M          = $(shell printf "\033[34;1m▶\033[0m")
SUFFIX 	   =

ifeq ($(OS),Windows_NT)
SUFFIX   =.exe
endif

GOX      = gox$(SUFFIX)
GO2XUNIT = go2xunit$(SUFFIX)
GOLINT   = golint$(SUFFIX)
GOCOV    = gocov$(SUFFIX)
GOCOVXML = gocov-xml$(SUFFIX)

export GO111MODULE=on

.PHONY: all
all: fmt lint | $(BIN) $(GOX); $(info $(M) building executable…) @ ## Build program binary
	$Q $(GOX) \
		-tags="release" \
		-ldflags="-X $(MODULE)/pkg/version.BUILDTIME=$(BUILDTIME) \
		          -X $(MODULE)/pkg/version.VERSION=$(VERSION) \
							-X $(MODULE)/pkg/version.GITSTATE=$(GITSTATE) \
							-X $(MODULE)/pkg/version.GITCOMMIT=$(GITCOMMIT) \
							-X $(MODULE)/pkg/version.GITBRANCH=$(GITBRANCH) " \
		-osarch="$(ARCH)" \
		-output="$(BIN)/{{.Dir}}_{{.OS}}_{{.Arch}}" .

$(GOLINT):
	$Q $(GO) install golang.org/x/lint/golint@latest

$(GOCOV):
	$Q $(GO) install github.com/axw/gocov/gocov@latest

$(GOCOVXML): 
	$Q $(GO) install github.com/AlekSi/gocov-xml@latest

$(GO2XUNIT):
	$Q $(GO) install github.com/tebeka/go2xunit@latest

$(GOX):
	$Q $(GO) install github.com/mitchellh/gox@latest

# Tools

$(BIN):
	@mkdir -p $@

# Tests

TEST_TARGETS := test-default test-bench test-short test-verbose test-race
.PHONY: $(TEST_TARGETS) test-xml check test tests
test-bench:   ARGS=-run=__absolutelynothing__ -bench=. ## Run benchmarks
test-short:   ARGS=-short        ## Run only short tests
test-verbose: ARGS=-v            ## Run tests in verbose mode with coverage reporting
test-race:    ARGS=-race         ## Run tests with race detector
$(TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
$(TEST_TARGETS): test
check test tests: fmt lint ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests
	$Q $(GO) test -timeout $(TIMEOUT)s $(ARGS) $(TESTPKGS)

test-xml: fmt lint | $(GO2XUNIT) ; $(info $(M) running xUnit tests…) @ ## Run tests with xUnit output
	$Q mkdir -p test
	$Q 2>&1 $(GO) test -timeout $(TIMEOUT)s -v $(TESTPKGS) | tee test/tests.output
	$(GO2XUNIT) -fail -input test/tests.output -output test/tests.xml

COVERAGE_MODE    = atomic
COVERAGE_PROFILE = $(COVERAGE_DIR)/profile.out
COVERAGE_XML     = $(COVERAGE_DIR)/coverage.xml
COVERAGE_HTML    = $(COVERAGE_DIR)/index.html
.PHONY: test-coverage test-coverage-tools
test-coverage-tools: | $(GOCOV) $(GOCOVXML)
test-coverage: COVERAGE_DIR := $(CURDIR)/test/coverage.$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
test-coverage: fmt lint test-coverage-tools ; $(info $(M) running coverage tests…) @ ## Run coverage tests
	$Q mkdir -p $(COVERAGE_DIR)
	$Q $(GO) test \
		-coverpkg=$$($(GO) list -f '{{ join .Deps "\n" }}' $(TESTPKGS) | \
					grep '^$(MODULE)/' | \
					tr '\n' ',' | sed 's/,$$//') \
		-covermode=$(COVERAGE_MODE) \
		-coverprofile="$(COVERAGE_PROFILE)" $(TESTPKGS)
	$Q $(GO) tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	$Q $(GOCOV) convert $(COVERAGE_PROFILE) | $(GOCOVXML) > $(COVERAGE_XML)

.PHONY: lint
lint: | $(GOLINT) ; $(info $(M) running golint…) @ ## Run golint
	$Q $(GOLINT) -set_exit_status $(PKGS)

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	$Q $(GO) fmt $(PKGS)

# Misc
.PHONY: gen
gen: ; $(info $(M) generate…)	@ ## generate
	$Q $(GO) generate

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf $(BIN)
	@rm -rf test/tests.* test/coverage.*

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)
