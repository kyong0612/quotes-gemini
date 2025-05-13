package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kyong0612/quotes-gemini/internal/gemini"
	"github.com/kyong0612/quotes-gemini/internal/githubgist"
	"github.com/kyong0612/quotes-gemini/prompt"
)

func main() {
	ctx := context.Background()

	// Load Asia/Tokyo location
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to load location: %w", err))
	}

	// generate quote by gemini
	geminiClient, err := gemini.NewClient(ctx)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create gemini client: %w", err))
	}

	defer geminiClient.Close()

	quote, err := geminiClient.GenerateContent(ctx, fmt.Sprintf(prompt.GenPositiveMessage, time.Now().In(jst).Format("2006年1月2日")))
	if err != nil {
		log.Fatal(fmt.Errorf("failed to generate content: %w", err))
	}

	fmt.Println(quote)

	// post to github gist
	gistClient, err := githubgist.NewClient()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create githubgist client: %w", err))
	}

	if err := gistClient.EditGist(ctx, "6016b8dd4f5d61f9cfcbbc8443d55260", "quote.md", quote); err != nil {
		log.Fatal(fmt.Errorf("failed to edit gist: %w", err))
	}

	fmt.Println("done")
}
