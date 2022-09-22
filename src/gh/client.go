package gh

import (
	"context"
	"sync"

	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
)

var (
	Client     *github.Client
	clientOnce sync.Once
)

func NewChient(ctx context.Context, token string) {
	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}))
	tc.Transport = &oauth2.Transport{
		Source: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}),
	}

	clientOnce.Do(func() {
		Client = github.NewClient(tc)
	})
}
