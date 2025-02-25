.PHONY: test test-integration

test:
	go test ./...

test-integration:
	go test ./test/integration/... -v

# convert this make config to use mise, delete makefile AI!

