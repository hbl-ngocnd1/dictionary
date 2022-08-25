NAME := get-started
MOCKGEN_BUILD_FLAGS ?= -mod=mod
GO ?= go
GO_ENV ?= GOPRIVATE="github.com/hbl-ngocnd1/dictionary"

deploy:
	ibmcloud cf push

go.list:
	go list -json -m all > go.list

mockgen:
	mockgen -build_flags=$(MOCKGEN_BUILD_FLAGS) -destination=./services/mock_services/mock_dictionary.go github.com/hbl-ngocnd1/dictionary/services DictionaryService
	mockgen -build_flags=$(MOCKGEN_BUILD_FLAGS) -destination=./services/mock_services/mock_translate.go github.com/hbl-ngocnd1/dictionary/services TranslateService

.PHONY: test
test: FLAGS ?= -parallel 3
test:
	$(GO_ENV) CI_TEST=test $(GO) test $(FLAGS) ./... -cover

coverage.out:
	go test -v -covermode=count -coverprofile=coverage.out ./...