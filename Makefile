test:
	go test -cover ./...

coverage.html: coverage.out
	go tool cover -html=$< -o $@

coverage.out:
	go test -coverprofile=$@ -coverpkg ./... ./...

.PHONY: \
	coverage.out \
	coverage.html \
	docker-unittest \
	test
