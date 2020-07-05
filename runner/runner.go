package runner

import (
	"log"
	"os"

	"github.com/mitch292/gimmeplan/git"
	"github.com/mitch292/gimmeplan/slack"
	"github.com/mitch292/gimmeplan/tf"
	"github.com/mitch292/gimmeplan/utils"
	"github.com/spf13/viper"
)

// PlanAndPostToSlack will clone a projects repo, run terraform plan and post the output to slack
func PlanAndPostToSlack(project string) {

	// 1: Clone the git repo
	repoName, err := git.Clone(viper.GetString(utils.GetViperString(project, "git_repo_url")))
	if err != nil {
		log.Fatalf("There was an error cloning the git repo: %s\n", err)
	}

	// 2: In the git repo, run terraform plan
	output, err := tf.Plan(repoName, project)

	// 3: Output this plan to slack
	slack.Send(viper.GetString(utils.GetViperString(project, "slack_webhook_url")), output)

	// 4 Remove the git repo
	os.RemoveAll("./" + repoName)
}
