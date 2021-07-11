.PHONY: build
build: .generate .build
.build:
		go build -o ./bin/server ./cmd/server/main.go

.PHONY: surf
surf: .surf
.surf:
		go build -o ./bin/surf ./cmd/surf/main.go

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

