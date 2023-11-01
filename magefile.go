//go:build mage
// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

const (
	IMAGE_NAME string = "ghcr.io/tomvictor/zephyr-build"
)

func Version() {
	fmt.Println(getLatestGitTag())
}

// Build the development docker image
func Build() error {
	return sh.Run("docker", "build", "--no-cache", "-t", getImage(), "-f", "docker/build.dockerfile", ".")
}

// Push the Pipeline docker image to gitlab container registry
func Push() {
	sh.RunV("docker", "push", getImage())
}

// Private functions

func getImage() string {
	// replace with getLatestGitTag
	return fmt.Sprintf("%s:%s", IMAGE_NAME, getLatestGitTag())
}

func getLatestGitTag() string {

	tag, err := sh.Output("git", "describe", "--tags", "--always", "--abbrev=0")
	if err != nil {
		return "v0.0.0"
	}

	return tag
}
