package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var manifestCmd = &cobra.Command{
	Use:   "manifest",
	Short: "Generates the manifest.yml file for cloudfoundary",
	Run: func(cmd *cobra.Command, args []string) {
		manifestTemplate := box.MustString("manifest.template")
		tmpl, err := template.New("version").Parse(manifestTemplate)
		if err != nil {
			fmt.Println("Failed to parse manifest template file. Error = ", err)
			os.Exit(-1)
		}
		type Manifest struct {
			Package    string
			MemorySize string
			HostName   string
			DistQuota  string
		}
		manifest := Manifest{
			Package:    packageName,
			MemorySize: memorySize,
			HostName:   hostName,
			DistQuota:  diskQuota,
		}

		filePath := filepath.Join(outputDirectory, "manifest.yml")
		f, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Failed to open ", filePath, " for writing. Error = ", err)
			os.Exit(-1)
		}
		defer f.Close()

		err = tmpl.Execute(f, manifest)
		if err != nil {
			fmt.Println("Failed to execute manifest.yml template file. Error = ", err)
			os.Exit(-1)
		}
	},
}

func AddManifestCommand(parent *cobra.Command) {
	parent.AddCommand(manifestCmd)

	cwd, err := os.Getwd()
	if err != nil {
		cwd = "."
	}

	manifestCmd.Flags().StringVarP(&packageName, "package", "p", getPackage(), "name of the application to be placed in the manifest file")
	manifestCmd.Flags().StringVarP(&hostName, "host", "s", getPackage(), "name of the host to be placed in the manifest file")
	manifestCmd.Flags().StringVarP(&memorySize, "memory", "m", "128M", "the size of the memory to be placed in the manifest file")
	manifestCmd.Flags().StringVarP(&diskQuota, "disk", "q", "1024M", "the disk quota to be placed in the manifest file")
	manifestCmd.Flags().StringVarP(&outputDirectory, "output", "o", cwd, "path to the output directory. "+
		"The manifest.yml file will be written in ${output_directory}/manifest.yml")
}
