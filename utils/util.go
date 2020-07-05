package utils

// GetViperString will generate the dot notation string used to access our nested config values
func GetViperString(project, keyName string) string {
	return "projects." + project + "." + keyName
}
