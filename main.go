package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kyong0612/quotes-gemini/external/gemini"
	"github.com/kyong0612/quotes-gemini/external/githubgist"
	"github.com/kyong0612/quotes-gemini/prompt"
)

func main() {
	ctx := context.Background()

	// generate quote by gemini
	gemini, err := gemini.NewClient(ctx)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create gemini client: %w", err))
	}

	defer gemini.Close()

	quote, err := gemini.GenerateContent(ctx, prompt.GenQuote)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to generate content: %w", err))
	}

	fmt.Println(quote)

	// post to github gist
	githubgist, err := githubgist.NewClient()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create githubgist client: %w", err))
	}

	if err := githubgist.EditGist(ctx, "6016b8dd4f5d61f9cfcbbc8443d55260", "quote.md", quote); err != nil {
		log.Fatal(fmt.Errorf("failed to edit gist: %w", err))
	}

	fmt.Println("done")
}
