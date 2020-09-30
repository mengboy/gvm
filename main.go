package main

import "github.com/spf13/cobra"

var rootCMD = &cobra.Command{
	Use:   "gvm",
	Short: "go version manager",
}

func init() {
	rootCMD.AddCommand(installCMD)
	rootCMD.AddCommand(remoteCMD)
	rootCMD.AddCommand(localListCMD)
	rootCMD.AddCommand(useCMD)
	rootCMD.AddCommand(clearCMD)
	rootCMD.AddCommand(gvmVersionCMD)
}

func main() {
	_ = rootCMD.Execute()
}
