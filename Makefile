.PHONY: run
run:
	@go build -v -o quotes-gemini . && ./quotes-gemini

.PHONY: lint
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run --fix
	go run golang.org/x/vuln/cmd/govulncheck ./...

