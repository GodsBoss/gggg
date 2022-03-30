examples=$(shell ls -I serve.go -I example.html -I assets examples)
assets=$(shell ls examples/assets)
wasms=$(foreach example,$(examples),dist/examples/$(example)/main.wasm)
webfiles=$(foreach example,$(examples),dist/examples/$(example)/index.html)

all: \
	dist/examples/wasm_exec.js \
	$(foreach asset,$(assets),dist/examples/assets/$(asset)) \
	$(wasms) \
	$(webfiles)

dist/examples/wasm_exec.js: $(shell go env GOROOT)/misc/wasm/wasm_exec.js dist/examples
	cp $< $@

dist/examples/%/main.wasm: FORCE
	mkdir -p dist/examples/$*
	GOOS=js GOARCH=wasm go build -o $@ ./examples/$*/main.go

dist/examples/%/index.html: examples/example.html
	mkdir -p dist/examples/$*
	cp $< $@

dist/examples/assets/%:
	mkdir -p dist/examples/assets
	cp examples/assets/$* $@

dist/examples:
	mkdir -p $@

serve:
	SERVE_DIR=$${PWD}/dist/examples go run ./examples/serve.go

docker-unittest:
	docker run \
		--rm \
		-v $${PWD}:/root/gggg:ro \
		--workdir /root/gggg \
		golang:1.15.2 \
		go test -v -cover ./pkg/...

clean:
	rm -rf dist

.PHONY: \
	all \
	clean \
	docker-unittest \
	serve

FORCE:
