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

type profileDetails struct {
	Login     string
	CreatedAt time.Time
	AvatarURL string `graphql:"avatarUrl(size: 72)"`
}

type repoDetails struct {
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	URL         githubv4.URI
}

// queryViewer gets metadata about the authenticated GitHub account.
func queryViewer(api *githubv4.Client) profileDetails {
	var viewerQuery struct {
		Viewer profileDetails
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

// queryRepo gets info about a given repo.
func queryRepo(api *githubv4.Client, owner string, repoName string) repoDetails {
	var repoQuery struct {
		Repository repoDetails `graphql:"repository(owner:$repositoryOwner, name:$repositoryName)"`
	}
	variables := map[string]interface{}{
		"repositoryOwner": githubv4.String(owner),
		"repositoryName":  githubv4.String(repoName),
	}

	err := api.Query(context.Background(), &repoQuery, variables)
	if err != nil {
		log.Fatalln(err)
	}

	return repoQuery.Repository
}

// printJSON pretty prints given value as JSON. It panics on an error.
func printJSON(v interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")

	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}

// Request will query the GitHub GQL API and handle the response data.
func Request() {
	api := conn.SetupAPIClient()

	profileResp := queryViewer(api)
	printJSON(profileResp)

	repoResp := queryRepo(api, "MichaelCurrin", "github-gql-go")
	printJSON(repoResp)
}
