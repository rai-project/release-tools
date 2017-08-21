package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	shellwords "github.com/junegunn/go-shellwords"
	"github.com/spf13/cobra"
)

type cf struct {
	Args string
}

func (c cf) GetCommand() (string, error) {
	cfExec, err := exec.LookPath("cf")
	if err != nil {
		fmt.Println("cf command not found. Error = ", err)
		return "", err
	}
	sh := os.ExpandEnv(c.Args)
	args, err := shellwords.Parse(sh)
	if err != nil {
		fmt.Println("cf failed to parse command shell. Error = ", err)
		return "", err
	}
	args = append([]string{cfExec}, args...)
	return strings.Join(args, " "), nil
}

func (c cf) Run() ([]byte, error) {
	joinedArgs, err := c.GetCommand()
	if err != nil {
		return nil, err
	}
	args := strings.Split(joinedArgs, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = packagePath
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("cf command failed to run. Error = ", err)
		return nil, err
	}
	return out, nil
}

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"push"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := runChain(
			cf{"api https://api.ng.bluemix.net"},
			cf{"auth ${BLUEMIX_USER} ${BLUEMIX_PASSWORD}"},
			cf{"a"},
			cf{`push "${BLUEMIX_PROJECT}" -b https://github.com/cloudfoundry/go-buildpack.git`},
			cf{"logs ${BLUEMIX_PROJECT} --recent"},
		)
		if err != nil {
			fmt.Println("error running deploy command. Error = ", err)
			os.Exit(-1)
		}
	},
}

func AddDeployCommand(parent *cobra.Command) {
	parent.AddCommand(deployCmd)
}
