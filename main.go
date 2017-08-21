//go:generate go generate cmd

package main

import "github.com/rai-project/release-tools/cmd"

func main() {
	cmd.Init()
	cmd.Execute()
}
