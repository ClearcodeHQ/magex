package adr

import (
	"fmt"
	"os"
	"strings"

	"github.com/magefile/mage/sh"
)

// Install installs adr tools in a version specified in a ADR_VERSION environment variable.
// If the variable is not passed, latest version is used. To force updating the tool
// use ADR_VERSION=latest.
func Install() error {
	var url string
	version := os.Getenv("ADR_VERSION")

	// URL with version must have "v" in the beginning
	if version == "latest" {
		url = "github.com/marouni/adr@latest"
	} else {
		url = fmt.Sprintf("github.com/marouni/adr@v%s", version)
	}

	// check exit code and version of adr tools
	output, err := sh.Output("adr", "--version")
	if sh.ExitStatus(err) != 0 || !strings.Contains(output, version) {
		fmt.Println(
			"adr tools not found or found in a different than requested version, installing...",
		)
		return sh.RunV("go", "install", url)
	}

	fmt.Println("adr tools already installed, skipping...")
	return nil
}

// Init initializes adr in a given directory. Requires one argument - directory
func Init(dir string) error {
	return sh.RunV("adr", "init", dir)
}

// New creates a new adr entry file. Requires one argument - entry title
func New(title string) error {
	return sh.RunV("adr", "new", title)
}
