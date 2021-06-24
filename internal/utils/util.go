package utils

import (
	"log"

	"github.com/spf13/viper"
)

// GetViperString will generate the dot notation string used to access our nested config values
func GetViperString(project, keyName string) string {
	accessString := viper.GetString("projects." + project + "." + keyName)
	if len(accessString) == 0 {
		log.Fatal("The project or key does not exist in the config file.")
	}
	return accessString
}
