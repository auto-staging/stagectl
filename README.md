# Stagectl

> Stagectl is the CLI tool to manage your auto-staging setup.

## Configuration of stagectl

Stagectl is configured by the ".stagectl.yaml" configuration file stored in your home directory.

The config file has the following format:

```yaml
tower_base_url: https://xxxxxxxxxxx.execute-api.eu-central-1.amazonaws.com/v1
```

Replace tower_base_url with your Tower API URL.

## Configuration for Repositories

Stagectl loads the configuration for a new repository or an repository update from the ".auto-staging.json" file in the current directory.

The file has the following format. Replace all values with the values matching your repository.

```json
{
    "repository": "auto-staging-demo-app",
    "infrastructureRepoURL": "https://github.com/janritter/auto-staging-demo-app.git",
    "webhook": true,
    "filters": [
        "feat(.*)"
    ],
    "codeBuildRoleARN": "arn:aws:iam::xxxxxxxxxxxxxx:role/codebuild-exec-role",
    "environmentVariables": [
        {
            "name": "TF_VAR_project",
            "type": "PLAINTEXT",
            "value": "demo-app"
        },
        {
            "name": "TF_VAR_stage",
            "type": "PLAINTEXT",
            "value": "dev"
        },
        {
            "name": "TF_VAR_instance_type",
            "type": "PLAINTEXT",
            "value": "t3.nano"
        },
        {
            "name": "TF_VAR_instance_count",
            "type": "PLAINTEXT",
            "value": "2"
        },
        {
            "name": "TF_VAR_github_owner",
            "type": "PLAINTEXT",
            "value": "janritter"
        },
        {
            "name": "TF_VAR_github_repo",
            "type": "PLAINTEXT",
            "value": "auto-staging-demo-app"
        },
        {
            "name": "GITHUB_TOKEN",
            "type": "PARAMETER_STORE",
            "value": "/CodeBuild/GITHUB_TOKEN"
        }
    ]
}
```

## Usage

### Install dependencies

```bash
make prepare
```

### Build binary

```bash
make build
```

compiles to bin folder

### Move binary to /usr/bin

```bash
sudo mv bin/stagectl /usr/bin/stagectl
```

## License and Author

Author: Jan Ritter

License: MIT