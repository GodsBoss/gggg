all: dist/examples/wasm_exec.js

dist/examples/wasm_exec.js: $(GOROOT)/misc/wasm/wasm_exec.js dist/examples
	cp $< $@

dist/examples:
	mkdir -p $@

clean:
	rm -rf dist

.PHONY: \
	all \
	clean
