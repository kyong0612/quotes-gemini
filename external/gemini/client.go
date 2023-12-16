package gemini

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type client struct {
	*genai.Client
	*genai.GenerativeModel
}

type Client interface {
	GenerateContent(ctx context.Context, input string) (string, error)
	Close()
}

func NewClient(ctx context.Context) (Client, error) {
	c, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return nil, err
	}

	modal := c.GenerativeModel("gemini-pro")

	return &client{c, modal}, nil
}

func (c client) GenerateContent(ctx context.Context, input string) (string, error) {
	resp, err := c.GenerativeModel.GenerateContent(ctx, genai.Text(input))
	if err != nil {
		return "", err
	}

	return fmt.Sprintln(resp.Candidates[0].Content.Parts[0]), nil
}

func (c client) Close() {
	c.Client.Close()
}
