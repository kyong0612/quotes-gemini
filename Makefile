.PHONY: run
run:
	@go build -v -o quotes-gemini . && ./quotes-gemini
