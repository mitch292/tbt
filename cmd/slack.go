/*
Copyright © 2020 Andrew Mitchell <andrewpmitchell7@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"

	"github.com/mitch292/tbt/internal/tbt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// slackCmd represents the slack command
var slackCmd = &cobra.Command{
	Use:   "slack",
	Short: "Output the terraform plan to the slack webhooks configured in your .tbt file",
	Long: `Output the terraform plan to the slack webhooks configured in your .tbt file
You can run plan for only a single file or instead output plans for all your terraform repos.
It's good practice to output these plans to different slack channels to keep things organized,
even though it leads to managing more webhooks in slack.`,
	Run: slackCmdRun,
}

func init() {
	rootCmd.AddCommand(slackCmd)

	slackCmd.Flags().StringP("project", "p", "", "The name of the project set in your .tbt config file for this plan")
}

func slackCmdRun(cmd *cobra.Command, args []string) {
	project, err := cmd.Flags().GetString("project")
	if err != nil {
		log.Fatalf("Unable to parse the project name passed., %s\n", err)
	}

	if len(project) > 0 {
		tbt.GetTerraformPlanAndPostToSlack(project)
	} else {
		projects := viper.GetStringMap("projects")

		for project := range projects {
			tbt.GetTerraformPlanAndPostToSlack(project)
		}
	}
}
