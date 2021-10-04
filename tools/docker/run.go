// Package docker contains helper calls for docker commands
package docker

import (
	"fmt"
	"os/user"

	"github.com/magefile/mage/sh"
)

// Run runs docker run command with current user's uuid and guid
func Run(dockerRunParams ...string) error {
	currentUser, err := user.Current()
	if err != nil {
		return err
	}
	dockerRunArgs := []string{
		"run",
		"--user",
		fmt.Sprintf("%s:%s", currentUser.Uid, currentUser.Gid),
	}
	dockerRunArgs = append(dockerRunArgs, dockerRunParams...)
	return sh.RunV(
		"docker",
		dockerRunArgs...,
	)
}
