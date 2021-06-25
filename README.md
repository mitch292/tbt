## What is tbt (To Be Terraformed)?

tbt is a cli tool that you can use to keep track of the state of your infrastructure compared to your terraform configuration. It'll clone your git repo, run terraform plan, output that to a slack channel.

## What's required to use tbt?

tbt is very much in MVP / single use case stage of development. To run today, you'll need a couple dependencies resolved.

-   Your own slack application with access to your workspace
-   A machine with git installed and SSH access to your git repo
-   Valid AWS credentials to run `terraform plan` - ONLY AWS IS SUPPORTED
-   Terraform state is managed through shared state

## How do I use tbt?

You configure a `.tbt.yaml` file that details the projects you want generate slack updates for.

To output a plan to slack for a single project you can use:

```
tbt slack -p=my_infra
```

or

```
tbt slack --project=my_infra
```

## What should my .tbt.yaml file look like?

The `tbt.yaml` file should be very simple. Below is a sample file structure.

```yaml
projects:
    web_infra:
        slack_webhook_url: "https://hooks.slack.com/services/YOUR_WEBHOOK_URL"
        git_repo_url: "git@github.com:user/repo.git"
        dir_name: "NAME_OF_DIR_TO_CREATE_AND_DELETE"
        aws_api_key: "YOUR_AWS_API_KEY"
        aws_secret: "YOUR_AWS_API_KEY_SECRET"
        aws_default_region: "YOUR_AWS_DEFAULT_REGION"
    reporting_worker:
        slack_webhook_url: "https://hooks.slack.com/services/YOUR_WEBHOOK_URL"
        git_repo_url: "git@github.com:user/repo.git"
        dir_name: "NAME_OF_DIR_TO_CREATE_AND_DELETE"
        aws_api_key: "YOUR_AWS_API_KEY"
        aws_secret: "YOUR_AWS_API_KEY_SECRET"
        aws_default_region: "YOUR_AWS_DEFAULT_REGION"
```
