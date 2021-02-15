# GitHub GraphQL Go
> An easy way to query GitHub's GraphQL API from the command-line

<!-- Badges generated with Badge Generator - https://michaelcurrin.github.io/badge-generator/ -->

[![Go CI](https://github.com/MichaelCurrin/github-gql-go/workflows/Go%20CI/badge.svg)](https://github.com/MichaelCurrin/github-gql-go/actions?query=workflow:"Go+CI")
[![GitHub tag](https://img.shields.io/github/tag/MichaelCurrin/github-gql-go?include_prereleases=&sort=semver)](https://github.com/MichaelCurrin/github-gql-go/releases/)
[![License](https://img.shields.io/badge/License-MIT-blue)](#license)

[![Made with Go](https://img.shields.io/github/go-mod/go-version/MichaelCurrin/github-gql-go?logo=go&logoColor=white)](https://golang.org)


A CLI tool which wraps another Go package and that can query the [GitHub GraphQL](https://michaelcurrin.github.io/dev-resources/resources/version-control/github/graphql.html) API.


## Sample usage

Create your GH auth token and set it on the environment.

```sh
$ export GH_TOKEN=abcdef
```

Run the CLI app to do a query and print results.

```sh
$ ghgql
```
```
Login: MichaelCurrin
Created at: 2016-04-30 11:19:17 +0000 UTC
Avatar URL: https://avatars.githubusercontent.com/u/18750745?s=72&u=ec21949f76c6d8f152f3d8c8f8204d86d6fceba5&v=4
```

**The query is hardcoded as a Go object and printed but in future this could be set up in a `.gql` input file and output saved as a `.json` file.**


## Documentation
> How to install and run this project locally and release it

<div align="center">

[![View - Documentation](https://img.shields.io/badge/View-Documentation-blue?style=for-the-badge)](/docs/)

</div>


## License

Released under [MIT](/LICENSE) by [@MichaelCurrin](https://github.com/MichaelCurrin).
