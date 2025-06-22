TOOL := $(shell basename $(shell pwd))

.PHONY: dev-install
dev-install:
	go install .

.PHONY: dev-doc
dev-doc:
	go run main.go doc

.PHONY: dev-man
dev-man:
	go run main.go man

.PHONY: build
build:
	GOOS=darwin GOARCH=arm64 go build -o $(TOOL) main.go

.PHONY: doc
doc:
	$(TOOL) doc

.PHONY: man
man:
	$(TOOL) man
