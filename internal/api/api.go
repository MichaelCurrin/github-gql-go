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

// printJSON prints v as JSON encoded with indent to stdout. It panics on any error.
func printJSON(v interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")

	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}

// Request will query the GitHub GQL API and return data.
func Request() {
	api := conn.SetupAPIClient()
	resp := viewer(api)
	printJSON(resp)
}
