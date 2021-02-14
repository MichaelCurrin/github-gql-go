# Configuration


## Generate token

Create Personal Access Token for authenticating with the GitHub API

1. Go GitHub and login.
2. Go to the [Personal Access Tokens](https://github.com/settings/tokens) page under Developer Settings.
3. Create a new token with appropriately scoped permissions.
    - Only **read** access is needed for this project.
    - Some recommended scopes
        * ☐ `repo` (Optionally tick the top-level one for access to _private_ repos)
            - ☑ `public_repo`
4. Find the generated token value, which you'll use in the next step.
    - Do not navigate away yet, as the token **cannot** be viewed online later. You can generate a new value for the token anytime and that will make the old value inactive.


## Store token

### Add to environment
> If using the CLI tool directly, without cloning this repo

Copy the token value and add the value to shell environment variables directly using `export`.

```sh
$ export GH_TOKEN=abcdef
```

You'll probably want to store it more permanently somewhere later.

TODO: For future development, read this value from a global file like `~/.config/ghgql/.env`.

### Store in project config
> For working in this repo directly

Create the base dotenv file.

```sh
$ cp .env.template .env
```

Then edit the new unversioned file `.env` file.

Paste your token in and save.

e.g.

- `.env`
    ```sh
    GH_TOKEN=abcdef
    ```
