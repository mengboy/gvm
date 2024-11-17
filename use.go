package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
)

var (
	useCMD = &cobra.Command{
		Use:   "use",
		Short: "change go version",
		Run:   use,
	}
)

func use(cmd *cobra.Command, args []string) {
	if len(args) <= 0 {
		fmt.Println("pls input version")
		return
	}
	version := args[0]
	arch := runtime.GOARCH
	goos := runtime.GOOS
	fileName := fmt.Sprintf(FileName, version, goos, arch)
	home, err := GetUserHomePath()
	if err != nil {
		log.Error(err)
		return
	}
	path := home + "/.gvm/" + fileName
	_, err = os.Stat(path)
	// 当前版本不存在
	if err != nil && !os.IsExist(err) {
		fmt.Println("download from origin")
		if err := Install(version, goos, arch, ""); err != nil {
			log.Error("install failed", err)
			return
		}
		return
	}
	tarCmd := exec.Command("/bin/sh", "-c", "sudo rm -rf /usr/local/go &&  sudo tar  -zxf "+path+" -C "+"/usr/local")
	if _, err := tarCmd.Output(); err != nil {
		log.Error("change go version failed", err)
	}
	fmt.Println("succ")
	fmt.Println("add following to your shell config if you haven't")
	fmt.Println("export PATH=$PATH:/usr/local/go/bin")
}
