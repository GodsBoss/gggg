all: dist/examples/wasm_exec.js dist/examples/wasd/main.wasm dist/examples/wasd/index.html

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

clean:
	rm -rf dist

.PHONY: \
	all \
	clean
