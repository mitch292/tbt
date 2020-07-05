package tf

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/viper"
)

// Plan will execute a terraform plan command in a given repo
func Plan(repoName, project string) ([]byte, error) {
	// TODO: this will all come as arugments that come from a config file

	initCmd := exec.Command("terraform", "init")
	initCmd.Env = append(os.Environ(),
		"AWS_ACCESS_KEY_ID="+viper.GetString(project+".aws_api_key"),
		"AWS_SECRET_ACCESS_KEY="+viper.GetString(project+".aws_secret"),
		"AWS_DEFAULT_REGION="+viper.GetString(project+".aws_default_region"))

	initCmd.Dir = "./" + repoName
	if initErr := initCmd.Run(); initErr != nil {
		log.Fatalf("There was a problem running terraform init: %s\n", initErr)
	}

	cmd := exec.Command("terraform", "plan", "-no-color")
	cmd.Dir = "./" + repoName
	cmd.Env = append(os.Environ(),
		"AWS_ACCESS_KEY_ID="+viper.GetString(project+".aws_api_key"),
		"AWS_SECRET_ACCESS_KEY="+viper.GetString(project+".aws_secret"),
		"AWS_DEFAULT_REGION="+viper.GetString(project+".aws_default_region"))

	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Didn't get the output: %s\n", err)
	}
	return out, nil
}
