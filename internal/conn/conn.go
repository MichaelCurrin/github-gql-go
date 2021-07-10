package conn

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// TokenName Name of the environment variable for storing the GitHub token.
const TokenName = "GH_TOKEN"

func setupHTTPClient(token string) *http.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	return oauth2.NewClient(context.Background(), src)
}

func SetupAPIClient() *githubv4.Client {
	token := os.Getenv(TokenName)
	if token == "" {
		log.Fatal("Must set token on environment variable: ", TokenName)
	}
	httpClient := setupHTTPClient(token)

	return githubv4.NewClient(httpClient)
}
