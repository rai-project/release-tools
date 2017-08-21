package cmd

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func getEnvAny(args ...string) (string, error) {
	envVars := append([]string{}, args...)
	for _, arg := range envVars {
		if val, ok := os.LookupEnv(arg); ok {
			return val, nil
		}
	}
	return "", errors.Errorf("None of the environments were found %v", envVars)
}

func getPackage() string {
	if packageName != "" {
		return packageName
	}
	if val, err := getEnvAny("BITBUCKET_REPO_SLUG", "TRAVIS_REPO_SLUG"); err == nil {
		return val
	}
	return "package_info_not_found"
}

func getVersion() string {
	version, err := getEnvAny("BITBUCKET_COMMIT", "VERSION", "TRAVIS_COMMIT")
	if err == nil {
		return version
	}
	return "version_info_not_found" // maybe do a git revparse
}

func getRepository() string {
	repo, err := getEnvAny("BITBUCKET_REPO_SLUG", "TRAVIS_REPO_SLUG")
	if err == nil {
		return repo
	}
	return "repo_info_not_found"
}

type Runnable interface {
	GetCommand() (string, error)
	Run() ([]byte, error)
}

func runChain(runables ...Runnable) error {
	for _, runable := range runables {
		cmd, err := runable.GetCommand()
		if err != nil {
			return err
		}
		fmt.Println("Running ", cmd)
		buf, err := runable.Run()
		if err != nil {
			return errors.Wrapf(err, "Failed to run %s", cmd)
		}
		fmt.Println(string(buf))
	}
	return nil
}
