GOPATH:=$(shell go env GOPATH)
ROOT_DIR=$(abspath ../../../../kratos-app)
VERSION=$(shell git describe --tags --always)
PROTO_FILES=$(shell find $(ROOT_DIR)/pkg/infra/config -name *.proto)

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=$(ROOT_DIR) \
	       --proto_path=$(ROOT_DIR)/third_party \
 	       --go_out=paths=source_relative:$(ROOT_DIR) \
	       $(PROTO_FILES)

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
