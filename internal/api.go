package api

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// TokenName Name of the environment variable for storing the GitHub token.
const TokenName = "GH_TOKEN"

// Request Query the GitHub GQL API and return data.
// TODO refactor into functions.
func Request() {
	token := os.Getenv(TokenName)
	if token == "" {
		log.Fatal("Must set token on environment variable: ", TokenName)
	}
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

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

	fmt.Println(q.Viewer.Login)
	fmt.Println(q.Viewer.CreatedAt)
	fmt.Println(q.Viewer.AvatarURL)
}
