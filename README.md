# GitHub GraphQL Go
> An easy way to query GitHub's GraphQL API from the command-line

<!-- Badges generated with Badge Generator - https://michaelcurrin.github.io/badge-generator/ -->

[![Go CI](https://github.com/MichaelCurrin/github-gql-go/workflows/Go%20CI/badge.svg)](https://github.com/MichaelCurrin/github-gql-go/actions?query=workflow:"Go+CI")
[![GitHub tag](https://img.shields.io/github/tag/MichaelCurrin/github-gql-go?include_prereleases=&sort=semver)](https://github.com/MichaelCurrin/github-gql-go/releases/)
[![License](https://img.shields.io/badge/License-MIT-blue)](#license)

[![Made with Go](https://img.shields.io/github/go-mod/go-version/MichaelCurrin/github-gql-go?logo=go&logoColor=white)](https://golang.org)
[![dependency - githubv4](https://img.shields.io/badge/dependency-githubv4-blue)](https://pkg.go.dev/github.com/shurcooL/githubv4)


A CLI tool to query the [GitHub GraphQL](https://michaelcurrin.github.io/dev-resources/resources/version-control/github/graphql.html) API by acting as a wrapper on another Go package.


## Sample usage

Create your GH auth token and set it on the environment.

```sh
$ export GH_TOKEN=abcdef
```

Run the CLI app to do a query and print results.

```console
$ ghgql
Login: MichaelCurrin
Created at: 2016-04-30 11:19:17 +0000 UTC
Avatar URL: https://avatars.githubusercontent.com/u/18750745?s=72&u=ec21949f76c6d8f152f3d8c8f8204d86d6fceba5&v=4
```


## Purpose

The intention for this project is:

- Be easy to use for developers who are not familiar with Go.
- Produce an binary executable for downloading (without needing Node or Python installed).
- Allow use interactively in the CLI or in project pipeline to handle (such as to download GitHub data as JSON data and use it to make a site).


## Limitations ⚠️

- This project is a proof-of-concept for extending. It only has short GQL query is **hardcoded** in this repo as a Go object and the result printed as text.
- The aim for future development:
    - Use a selection of `.gql` input file given by the user, or choose from queries known to the app.
    - And then output saved as a `.json` file.
- For a more full-fledged Python project that takes GQL files and writes CSV reports for each, see my [GH Reporting Tool](https://github.com/MichaelCurrin/github-reporting-py).


## Documentation
> How to install and run this project locally and release it

<div align="center">

[![View - Documentation](https://img.shields.io/badge/View-Documentation-blue?style=for-the-badge)](/docs/)

</div>


## License

Released under [MIT](/LICENSE) by [@MichaelCurrin](https://github.com/MichaelCurrin).
