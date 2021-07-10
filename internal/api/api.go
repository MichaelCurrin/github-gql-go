package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shurcooL/githubv4"

	conn "github.com/MichaelCurrin/github-gql-go/internal/conn"
)

type Viewer struct {
	Login     string
	CreatedAt time.Time
	AvatarURL string `graphql:"avatarUrl(size: 72)"`
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

// printAsJSON will accept generic response, convert to JSON and print it.
func printAsJSON(resp interface{}) {
	json, err := json.Marshal(resp)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", string(json))
}

// Request will query the GitHub GQL API and return data.
func Request() {
	api := conn.SetupAPIClient()
	resp := viewer(api)

	fmt.Fprintf(os.Stderr, "JSON data\n")

	printAsJSON(resp)
}
