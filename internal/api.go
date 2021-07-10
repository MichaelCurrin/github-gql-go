package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// TokenName Name of the environment variable for storing the GitHub token.
const TokenName = "GH_TOKEN"

type Viewer struct {
	Login     string
	CreatedAt time.Time
	AvatarURL string `graphql:"avatarUrl(size: 72)"`
}

func setupHTTPClient(token string) *http.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	return oauth2.NewClient(context.Background(), src)
}

func setupAPIClient() *githubv4.Client {
	token := os.Getenv(TokenName)
	if token == "" {
		log.Fatal("Must set token on environment variable: ", TokenName)
	}
	httpClient := setupHTTPClient(token)

	return githubv4.NewClient(httpClient)
}

// viewer gets metadata about the authenticated GitHub account.
func viewer(api *githubv4.Client) {
	var viewerQuery struct {
		Viewer Viewer
	}

	err := api.Query(context.Background(), &viewerQuery, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Login:", viewerQuery.Viewer.Login)
	fmt.Println("Created at:", viewerQuery.Viewer.CreatedAt)
	fmt.Println("Avatar URL:", viewerQuery.Viewer.AvatarURL)
}

// Request will query the GitHub GQL API and return data.
func Request() {
	api := setupAPIClient()
	viewer(api)
}
