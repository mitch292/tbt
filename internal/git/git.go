package git

import (
	"log"
	"os/exec"
)

// Clone will git clone a remote repository into the current directory
func Clone(repoURL string) {

	cmd := exec.Command("git", "clone", repoURL)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("There was a problem cloning the repo. Please make sure the repo exists. %s\n", err)
	}
}
