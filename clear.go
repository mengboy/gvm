package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
)

var (
	clearCMD = &cobra.Command{
		Use:   "clear",
		Short: "clear downloaded go",
		Run:   clear,
	}
)

func clear(cmd *cobra.Command, args []string) {
	home, err := GetUserHomePath()
	if err != nil {
		log.Error("get home path failed", err)
		return
	}

	gvmDir := home + "/.gvm/"
	param := ""
	if len(args) > 0 {
		param = args[0]
	}

	// If "all" is specified, remove the entire .gvm directory
	if param == "all" {
		if err := os.RemoveAll(gvmDir); err != nil {
			log.Error("failed to remove all files:", err)
			return
		}
		return
	}

	// Read directory contents
	files, err := ioutil.ReadDir(gvmDir)
	if err != nil {
		log.Error("failed to read directory:", err)
		return
	}

	// Remove files matching the parameter
	for _, file := range files {
		if param == "" || strings.Contains(file.Name(), param) {
			filePath := gvmDir + file.Name()
			if err := os.RemoveAll(filePath); err != nil {
				log.Error("failed to remove file:", filePath, err)
				continue
			}
		}
	}
}
