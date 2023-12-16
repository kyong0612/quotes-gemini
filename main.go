package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kyong0612/quotes-gemini/external/gemini"
	"github.com/kyong0612/quotes-gemini/prompt"
)

func main() {

	ctx := context.Background()

	gemini, err := gemini.NewClient(ctx)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create gemini client: %v", err))
	}

	defer gemini.Close()

	quote, err := gemini.GenerateContent(ctx, prompt.GenQuote)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to generate content: %v", err))
	}

	fmt.Println(quote)

}
