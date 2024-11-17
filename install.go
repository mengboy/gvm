package main

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
)

var (
	installCMD = &cobra.Command{
		Use:   "install",
		Short: "install go",
		Run:   install,
	}
	toOS   string
	toArch string
	from   string
)

func init() {
	installCMD.PersistentFlags().StringVarP(&toOS, "toOS", "o", "", "操作系统如linux、darwin")
	installCMD.PersistentFlags().StringVarP(&toArch, "toArch", "a", "", "架构如386、amd64")
	installCMD.PersistentFlags().StringVarP(&from, "from", "f", "", "安装源默认go官方, 可以设置studygolang走国内")
}

func install(cmd *cobra.Command, args []string) {
	var version string
	if len(args) > 0 {
		version = args[0]
	}
	if len(args) > 1 {
		from = "studygolang"
	}
	if err := Install(version, toOS, toArch, from); err != nil {
		log.Error("install failed", err)
	}

}

func Install(version string, os string, arch string, form string) error {
	installOS := runtime.GOOS
	installArch := runtime.GOARCH
	if os != "" {
		installOS = os
	}
	if arch != "" {
		installArch = arch
	}
	fileName := fmt.Sprintf(FileName, version, installOS, installArch)
	defaultDownloadURL := GOLang + GOLangDLURL + fileName
	backupDownloadURL := StudyGOLang + StudyGOLangDLURL + fileName
	if from == "studygolang" {
		defaultDownloadURL, backupDownloadURL = backupDownloadURL, defaultDownloadURL
	}
	home, err := GetUserHomePath()
	if err != nil {
		return err
	}
	err = DownloadFileProgress(defaultDownloadURL, backupDownloadURL, fileName)
	if err != nil {
		return err
	}
	fmt.Println()
	// TODO check sum
	tarCmd := exec.Command("/bin/sh", "-c", "sudo rm -rf /usr/local/go && sudo tar  -zxf "+home+"/.gvm/"+fileName+" -C "+"/usr/local")
	if _, err := tarCmd.Output(); err != nil {
		return err
	}
	fmt.Println("succ")
	fmt.Println("add following to your shell config if you haven't")
	fmt.Println("export PATH=$PATH:/usr/local/go/bin")
	return nil
}
