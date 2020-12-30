# Maven Version Fetcher

`v1.0.0`<br>
`Status: Initial`<br>
`Created on: 30.12.2020`<br>
`Last Update: 30.12.2020`<br>

## Configuration

You use the `conf.json` file to configure authentication and repository URLs. You need to get a personal access token from GitLab in order to authenticate with your GitLab projects. 

>GitLab -> Settings -> Access Tokens -> Enable `api`, `read_registry` Scopes -> Create Personal Access Token

```json
{
  "token": "your-token",
  "urls": [
    "https://gitlab.com/api/v4/projects/project-id/packages",
    "https://gitlab.com/api/v4/projects/project-id/packages",
    "https://gitlab.com/api/v4/projects/project-id/packages",
  ]
}
```

## Usage

Build the go project with executing `go build` in project root directory, then run the executable `./maven-version-fetcher`.