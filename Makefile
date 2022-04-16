docker-unittest:
	docker run \
		--rm \
		-v $${PWD}:/root/gggg:ro \
		--workdir /root/gggg \
		golang:1.18.0-alpine3.15 \
		go test -v -cover ./pkg/...

test:
	go test -cover ./...

.PHONY: \
	docker-unittest \
	test
