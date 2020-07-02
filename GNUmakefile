PKG_NAME=intersight
VERSION=0.1.3
TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
WEBSITE_REPO=github.com/hashicorp/terraform-website
SWAGGER_SPEC=spec/swagger.json
GENERATED_FOLDERS = $(PKG_NAME) models website

export GOOS=$(shell go env GOOS)
export GO_BUILD=env go build
export GO_RUN=env go run
export GO_VET=env go vet
export GO_TEST=env go test
export GO111MODULE=on


default: build

build: fmtcheck
	go mod vendor
	go install

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'scripts/gofmtcheck.sh'"

lint:
	@echo "==> Checking source code against linters..."
	tfproviderlint ./intersight
	golangci-lint run ./...

test: fmtcheck
#	go test -i $(TEST) || exit 1
#	echo $(TEST) | \
#		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc: fmtcheck
	#TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

test-compile:
#	@if [ "$(TEST)" = "./..." ]; then \
#		echo "ERROR: Set TEST to a specific package. For example,"; \
#		echo "  make test-compile TEST=./$(PKG_NAME)"; \
#		exit 1; \
#	fi
#	go test -c $(TEST) $(TESTARGS)

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

website:
	ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
		echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
		git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
	endif
		@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

website-lint:
	@echo "==> Checking website against linters..."
	@misspell -error -source=text website/

website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

clean:
	go clean --cache
	rm -rf vendor .build

clobber:
	go clean --cache
	rm -rf $(GENERATED_FOLDERS) vendor $(SWAGGER_SPEC) .build

.PHONY: build test testacc vet fmt fmtcheck lint tools test-compile website website-lint website-test