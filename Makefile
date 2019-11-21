PKG_NAME=intersight
SWAGGER_SPEC=spec/swagger.json
GENERATED_FOLDERS = $(PKG_NAME) models website

export GOOS=$(shell go env GOOS)
export GO_BUILD=env go build
export GO_RUN=env go run
export GO_INSTALL=env go install
export GO_VET=env go vet
export GO_TEST=env go test
export GO111MODULE=on

V = 0
Q = $(if $(filter 1,$V),,@)


default: install

build: vet
	@echo "building terraform-provider-intersight"
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o .build/linux_amd64/terraform-provider-intersight
	GOOS=windows GOARCH=amd64 $(GO_BUILD) -o .build/windows/terraform-provider-intersight
	GOOS=darwin GOARCH=amd64 $(GO_BUILD) -o .build/darwin/terraform-provider-intersight

install: vet build
	@echo "installing terraform-provider-intersight"
	$(GO_INSTALL)

clean:
	go clean --cache
	rm -rf vendor .build

clobber:
	go clean --cache
	rm -rf $(GENERATED_FOLDERS) vendor $(SWAGGER_SPEC) .build

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
		fi

.PHONY: build clean vet install
