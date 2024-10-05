run:
	@redocly bundle ./docs/openapi.spec.yaml -o ./docs/spec.yaml
	@go run .


BIN := $(shell pwd)/.bin
TOOLS := $(shell pwd)/tools

# install needed tools
install-tools:
	@echo Installing tools from $(TOOLS)/tools.go
	@cd $(TOOLS) && cat tools.go | grep _ | awk -F'"' '{print $$2}' | GOBIN=$(BIN) xargs -tI % go install %

BIN_PATH := PATH="$(abspath $(BIN)):$$PATH"

generate:
	@echo Generating go files...
	@$(BIN_PATH) go generate ./...
