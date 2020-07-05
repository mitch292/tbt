package utils

import "strings"

// GetViperString will generate the dot notation string used to access our nested config values
func GetViperString(project, keyName string) string {
	return "projects." + project + "." + keyName
}

// RemoveRefreshData will just take the plan data that we want to see
func RemoveRefreshData(output string) string {
	sliceOfString := strings.Split(output, "------------------------------------------------------------------------")
	return sliceOfString[len(sliceOfString)-2]
}
