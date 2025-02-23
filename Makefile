LOCAL_BIN:=$(CURDIR)/bin
PATH:=$(PATH):$(LOCAL_BIN)

SPEC_VERSION := v0.0.5
OGEN_VERSION := v1.10.0
LINT_VERSION := v1.64.5

SPEC_URL := https://raw.githubusercontent.com/reaport/docs/$(SPEC_VERSION)/openAPI/ground_control.yaml

TARGET_DIR := pkg/api

download-ogen:
	GOBIN=$(LOCAL_BIN) go install github.com/ogen-go/ogen/cmd/ogen@$(OGEN_VERSION)

download-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(LINT_VERSION)

generate-api: download-ogen
	$(LOCAL_BIN)/ogen --package api --target $(TARGET_DIR) --clean $(SPEC_URL)

clean:
	rm -rf $(TARGET_DIR)

lint: download-golangci-lint
	$(LOCAL_BIN)/golangci-lint run --fix

.PHONY: download-ogen download-golangci-lint generate-api clean lint