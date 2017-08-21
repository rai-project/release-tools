//go:generate rice embed-go

package cmd

import (
	"fmt"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/spf13/cobra"
)

var (
	box             = rice.MustFindBox("_fixtures")
	packagePath     string
	packageName     string
	hostName        string
	memorySize      string
	diskQuota       string
	outputDirectory string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "bluemixdeployment",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	packagePath = os.ExpandEnv("${GOPATH}/src/bitbucket.org/${BITBUCKET_REPO_OWNER}/${BITBUCKET_REPO_SLUG}")
}

func Init() {
	AddDeployCommand(RootCmd)
	AddGeneratorCommand(RootCmd)
}
