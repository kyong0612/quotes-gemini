.PHONY: run
run:
	@go build -v -o quotes-gemini . && ./quotes-gemini

.PHONY: lint
lint:
	go tool golangci-lint run --fix
	go tool govulncheck ./...

