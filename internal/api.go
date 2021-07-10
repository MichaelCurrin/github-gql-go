package api

import (
	"context"
	"encoding/json"
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
func viewer(api *githubv4.Client) Viewer {
	var viewerQuery struct {
		Viewer Viewer
	}

	err := api.Query(context.Background(), &viewerQuery, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintln(os.Stderr, "Pretty view:")
	fmt.Fprintln(os.Stderr, "Login:", viewerQuery.Viewer.Login)
	fmt.Fprintln(os.Stderr, "Created at:", viewerQuery.Viewer.CreatedAt)
	fmt.Fprintln(os.Stderr, "Avatar URL:", viewerQuery.Viewer.AvatarURL)
	fmt.Fprint(os.Stderr, "\n\n")

	return viewerQuery.Viewer
}

// process will accept generic response, convert to JSON and print it.
func process(resp interface{}) {
	json, err := json.Marshal(resp)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", string(json))
}

// Request will query the GitHub GQL API and return data.
func Request() {
	api := setupAPIClient()
	resp := viewer(api)

	fmt.Fprintf(os.Stderr, "JSON data\n")

	process(resp)
}
