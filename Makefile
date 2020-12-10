all: build

build: format $(patsubst %, build-%, $(PLUGINS))
	cd cmd/kvm && go build

format:
	go fmt ./pkg/...
	go vet ./pkg/...

test:
	go test ./cmd/... ./pkg/...

test-%:
	go test ./$(subst -,/,$*)/...

functest:
	pytest tests

dep:
	dep ensure -v

clean-dep:
	rm -f ./Gopkg.lock
	rm -rf ./vendor

.PHONY: format build test docker-build docker-push docker-local-push dep clean-dep

