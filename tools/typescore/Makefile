TOOL := $(shell basename $(shell pwd))

.PHONY: dev-install
dev-install:
	go install .

.PHONY: doc
doc:
	rm -fr ./docs/
	go run main.go doc

.PHONY: man
man:
	go run main.go man

.PHONY: completions
completions:
	go run main.go completion bash > $(TOOL).bash
	go run main.go completion zsh > $(TOOL).zsh
	go run main.go completion fish > $(TOOL).fish

.PHONY: build
build:
	GOOS=darwin GOARCH=arm64 go build -o $(TOOL) main.go

.PHONY: package
package: man completions build
	mkdir -p dist/$(TOOL)
	mv *.1 dist/$(TOOL)
	mv $(TOOL).bash dist/$(TOOL)
	mv $(TOOL).zsh dist/$(TOOL)
	mv $(TOOL).fish dist/$(TOOL)
	mv $(TOOL) dist/$(TOOL)
	tar -czf dist/$(TOOL).tar.gz -C dist $(TOOL)
	mv dist/$(TOOL).tar.gz .
	rm -fr dist/
