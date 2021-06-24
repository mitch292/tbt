package tf

import (
	"log"
	"os"
	"os/exec"

	"github.com/mitch292/tbt/internal/utils"
)

// Plan will execute a terraform plan command in a given repo
func Plan(repoName, project string) []byte {

	initCmd := exec.Command("terraform", "init")
	initCmd.Env = append(os.Environ(),
		"AWS_ACCESS_KEY_ID="+utils.GetViperString(project, "aws_api_key"),
		"AWS_SECRET_ACCESS_KEY="+utils.GetViperString(project, "aws_secret"),
		"AWS_DEFAULT_REGION="+utils.GetViperString(project, "aws_default_region"))

	initCmd.Dir = "./" + repoName
	if initErr := initCmd.Run(); initErr != nil {
		log.Fatalf("There was a problem running terraform init: %s\n", initErr)
	}

	planCmd := exec.Command("terraform", "plan", "-no-color")
	planCmd.Dir = "./" + repoName
	planCmd.Env = append(os.Environ(),
		"AWS_ACCESS_KEY_ID="+utils.GetViperString(project, "aws_api_key"),
		"AWS_SECRET_ACCESS_KEY="+utils.GetViperString(project, "aws_secret"),
		"AWS_DEFAULT_REGION="+utils.GetViperString(project, "aws_default_region"))

	out, err := planCmd.Output()
	if err != nil {
		log.Fatalf("Didn't get the output: %s\n", err)
	}
	return out
}
