.PHONY: build
build: .generate .build
.build:
		go build -o ./bin/server ./cmd/main.go

.PHONY: generate
generate: .generate
.generate:
		go generate ./...

.PHONY: deps
deps: .deps

.PHONY: .deps
.deps:
		go get github.com/jessevdk/go-assets-builder
		go install github.com/jessevdk/go-assets-builder

