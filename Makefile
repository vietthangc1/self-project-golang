GOPATH              := $(or $(GOPATH), $(HOME)/go)
WIRE                := $(GOPATH)/bin/wire

ROOT_DIR                := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
DIST_DIR                := out
SERVICES                := $(shell find services/ -mindepth 1 -maxdepth 1 -type d)
MOCK_TARGETS            := $(addprefix mock/,$(SERVICES))

CMDS                    := $(shell find $(ROOT_DIR)/cmd -mindepth 1 -maxdepth 1 -type d)
STATIC_TARGETS          := $(addprefix static-,$(CMDS))
DI_TARGETS              := $(addprefix di-,$(CMDS))

SERVICES                := $(shell find $(ROOT_DIR)/services/ -mindepth 1 -maxdepth 1 -type d)
CI_INTEGRATIONS_TARGETS := $(addprefix ci-,$(SERVICES))

$(WIRE):
	GOPATH=$(GOPATH) go install -mod=mod github.com/google/wire/cmd/wire

di: $(WIRE) $(DI_TARGETS) $(DI_TEST_TARGETS)
$(DI_TARGETS):
	$(WIRE) $(subst di-,,$@)

run-server:
	go run ./cmd/server

run-product:
	go run ./cmd/product_server

lint:
	golangci-lint run --timeout 10m
