package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "The version command creates a version.go file in the directory specified by the -d option.",
	Long: `The version command creates a version.go file in the directory specified by the -d option.
The version information is read from environment variables. The tool is designed to be used within release-tools.`,
	Run: func(cmd *cobra.Command, args []string) {
		versionTemplate := box.MustString("version.go.template")
		tmpl, err := template.New("version").Parse(versionTemplate)
		if err != nil {
			fmt.Println("Failed to parse version template file. Error = ", err)
			os.Exit(-1)
		}
		type Version struct {
			Package       string
			VersionInfo   string
			Repository    string
			BuildTimeInfo string
		}
		version := Version{
			Package:       getPackage(),
			VersionInfo:   getVersion(),
			Repository:    getRepository(),
			BuildTimeInfo: time.Now().Format(time.RFC3339),
		}

		filePath := filepath.Join(outputDirectory, "version.go")
		f, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Failed to open ", filePath, " for writing. Error = ", err)
			os.Exit(-1)
		}
		defer f.Close()

		err = tmpl.Execute(f, version)
		if err != nil {
			fmt.Println("Failed to execute version template file. Error = ", err)
			os.Exit(-1)
		}
	},
}

func AddVersionCommand(parent *cobra.Command) {
	parent.AddCommand(versionCmd)

	cwd, err := os.Getwd()
	if err != nil {
		cwd = "."
	}

	versionCmd.Flags().StringVarP(&packageName, "package", "p", getPackage(), "name of the package of the version.go file")
	versionCmd.Flags().StringVarP(&outputDirectory, "output", "o", cwd, "path to the output directory. "+
		"The version.go file will be written in ${output_directory}/version.go")
}
