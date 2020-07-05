/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/mitch292/gimmeplan/runner"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Run the slack command for all your given projects",
	Long: `All of your projects configured in your .gimmeplan.yaml file will run the 
	slack command. Check the help command for 'slack' to learn more.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := viper.GetStringMap("projects")
		for projName := range config {
			runner.PlanAndPostToSlack(projName)
		}
	},
}

func init() {
	slackCmd.AddCommand(allCmd)
}
