package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	api "github.com/MichaelCurrin/github-gql-go/internal"
)

const (
	// VERSION is the release number.
	VERSION = "v0.1.0"
)

// run does the overall app logic.
func run(c *cli.Context) error {
	api.Request()

	return nil
}

func main() {
	app := &cli.App{
		Name:        "GitHub GQL Tool",
		HelpName:    "ghgql",
		Usage:       "Provide a GQL query to read and a JSON path to write to. Required env variable: GH_TOKEN",
		UsageText:   "ghgql [-h] [-v]",
		Version:     VERSION,
		Description: "Query the GitHub GQL API and return data as JSON",

		Action: run,

		Compiled: time.Time{},
		Authors: []*cli.Author{
			{
				Name:  "MichaelCurrin",
				Email: "",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
