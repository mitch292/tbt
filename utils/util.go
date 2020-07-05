package utils

import "strings"

const (
	tfOutputDelimiter = "------------------------------------------------------------------------"
	noChanges         = "No changes. Infrastructure is up-to-date."
)

// GetViperString will generate the dot notation string used to access our nested config values
func GetViperString(project, keyName string) string {
	return "projects." + project + "." + keyName
}

// GetMeaningfulTfOutput will just take the plan data that we want to see
func GetMeaningfulTfOutput(output string) string {
	// TODO: This is verrry hacky
	sliceOfString := strings.Split(output, tfOutputDelimiter)
	if strings.Contains(sliceOfString[len(sliceOfString)-1], noChanges) {
		return "The state of your infrastructure and Terraform are in sync!"
	}
	return sliceOfString[len(sliceOfString)-2]
}
