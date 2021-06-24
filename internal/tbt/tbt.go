package tbt

import (
	"os"

	"github.com/mitch292/tbt/internal/git"
	"github.com/mitch292/tbt/internal/slack"
	"github.com/mitch292/tbt/internal/tf"
	"github.com/mitch292/tbt/internal/utils"
)

// PlanAndPostToSlack will clone a projects repo, run terraform plan and post the output to slack
func GetTerraformPlanAndPostToSlack(project string) {

	// 1: Clone the git repo
	git.Clone(utils.GetViperString(project, "git_repo_url"))

	// 2: In the git repo, run terraform plan
	output := tf.Plan(utils.GetViperString(project, "dir_name"), project)

	// 3: Output this plan to slack
	slack.Send(utils.GetViperString(project, "slack_webhook_url"), output)

	// 4 Remove the git repo
	os.RemoveAll("./" + utils.GetViperString(project, "dir_name"))
}
