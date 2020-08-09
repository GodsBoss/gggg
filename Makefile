examples=$(shell ls -I serve.go -I example.html examples)

all: \
	dist/examples/wasm_exec.js \
	$(foreach example,$(examples),dist/examples/$(example)/main.wasm dist/examples/$(example)/index.html)

dist/examples/wasm_exec.js: $(GOROOT)/misc/wasm/wasm_exec.js dist/examples
	cp $< $@

dist/examples/%/main.wasm:
	mkdir -p dist/examples/$*
	GOOS=js GOARCH=wasm go build -o $@ ./examples/$*/main.go

dist/examples/%/index.html: examples/example.html
	mkdir -p dist/examples/$*
	cp $< $@

dist/examples:
	mkdir -p $@

serve:
	SERVE_DIR=$${PWD}/dist/examples go run ./examples/serve.go

clean:
	rm -rf dist

.PHONY: \
	all \
	clean \
	serve
