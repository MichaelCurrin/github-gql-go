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

func oAuth() *http.Client {
	token := os.Getenv(TokenName)
	if token == "" {
		log.Fatal("Must set token on environment variable: ", TokenName)
	}
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	return oauth2.NewClient(context.Background(), src)
}

func apiClient() *githubv4.Client {
	httpClient := oAuth()

	return githubv4.NewClient(httpClient)
}

// Request will query the GitHub GQL API and return data.
func Request() {
	client := apiClient()

	var q struct {
		Viewer struct {
			Login     string
			CreatedAt time.Time
			AvatarURL string `graphql:"avatarUrl(size: 72)"`
		}
	}

	err := client.Query(context.Background(), &q, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Login:", q.Viewer.Login)
	fmt.Println("Created at:", q.Viewer.CreatedAt)
	fmt.Println("Avatar URL:", q.Viewer.AvatarURL)
}
