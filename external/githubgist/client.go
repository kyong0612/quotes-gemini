package githubgist

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v60/github"
)

type client struct {
	*github.Client
}

type Client interface {
	EditGist(ctx context.Context, id string, fileName, content string) error
}

func NewClient() (Client, error) {

	token := os.Getenv("GIST_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("GITHUB GIST API TOKEN is empty")
	}

	return client{
		github.NewClient(nil).WithAuthToken(token),
	}, nil
}

func (c client) EditGist(ctx context.Context, id string, fileName, content string) error {
	_, _, err := c.Client.Gists.Edit(ctx, "6016b8dd4f5d61f9cfcbbc8443d55260", &github.Gist{
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(fileName): {
				Content: github.String(content),
			},
		},
	})
	if err != nil {
		return err
	}

	// NOTE: error handling by resp if needed

	return nil
}
