.PHONY: test test-integration

test:
	go test ./...

test-integration:
	go test ./test/integration/... -v
