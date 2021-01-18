# Maven Version Fetcher

`v1.0.1`<br>
`Status: Initial`<br>
`Created on: 30.12.2020`<br>
`Last Update: 18.01.2021`<br>

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


## Using with Path

You can use the script with adding a custom alias to your shell resource file. 

- Copy `conf.json` file to `home` directory.
- Open `.bashrc` or `.zshrc` file and add the alias below with your go binary file path.

 ```shell
 alias maven-versions="./go/src/maven-lib-versions/maven-version-fetcher"
 ```

- Restart the shell
- Run command maven-versions
