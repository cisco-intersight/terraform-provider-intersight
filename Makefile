PKG_NAME=intersight
SWAGGER_SPEC=codegen/resource/swagger.json
MANAGED_SOURCE_FILES = $(PKG_NAME)/resource_*.go $(PKG_NAME)/provider_resource_map.go
DATA_SOURCE_FILES = $(PKG_NAME)/data_source_*.go $(PKG_NAME)/provider_data_source_map.go
OTHER_GENERATED_FILES = $(PKG_NAME)/flatten_functions.go models
GENERATED_FILES= $(MANAGED_SOURCE_FILES) $(DATA_SOURCE_FILES) $(OTHER_GENERATED_FILES)
GO_FILES=$$(find . -name '*.go' | grep -v vendor)

export GOOS=$(shell go env GOOS)
export GO_BUILD=env go build
export GO_RUN=env go run
export GO_INSTALL=env go install
export GO_VET=env go vet
export GO_TEST=env go test
export GO111MODULE=on

V = 0
Q = $(if $(filter 1,$V),,@)


default: build

fmt: $(PKG_NAME)
	@echo "running gofmt..."
	goimports -w $(PKG_NAME)

models: $(SWAGGER_SPEC)
	@echo "generating model structs..."
	$Q go get github.com/go-swagger/go-swagger/cmd/swagger
	swagger generate model --spec $(SWAGGER_SPEC)

build: models fmt fmtcheck vet
	@echo "building terraform-provider-intersight"
	$(GO_INSTALL)

clean:
	rm -rf $(GENERATED_FILES) vendor $(SWAGGER_SPEC)
	go clean --cache

clobber: clean
	rm -rf glide.lock ~/.glide/cache

fmtcheck:
	gofmt -l $(GO_FILES)

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
		fi

.PHONY: build fmt clean clobber fmtcheck vet
