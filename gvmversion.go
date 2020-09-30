package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	gvmVersionCMD = &cobra.Command{
		Use:   "version",
		Short: "gvm version",
		Run:   gvmVersion,
	}
)

func gvmVersion(cmd *cobra.Command, args []string) {
	fmt.Println(GVMVersion)
}
