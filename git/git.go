package git

import (
	"log"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

// Clone will git clone a remote repository into the current directory
func Clone(repoURL string) (string, error) {

	cmd := exec.Command("git", "clone", repoURL)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("There was a problem cloning the repo: %s\n", err)
	}

	// convert a repo url like: https://github.com/mitch292/empty_terraform_repo.git
	// to the destination dirName like: empty_terraform_repo
	dirName := strings.TrimSuffix(path.Base(repoURL), filepath.Ext(path.Base(repoURL)))

	return dirName, nil
}
