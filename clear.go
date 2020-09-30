package main

import (
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
)

var (
	clearCMD = &cobra.Command{
		Use:   "clear",
		Short: "clear downloaded go",
		Run:   clear,
	}
)

func clear(cmd *cobra.Command, args []string)  {
	home, err := GetUserHomePath()
	if err != nil{
		log.Error("get home path failed", err)
		return
	}
	param := ""
	if len(args) > 0{
		param = args[0]
	}
	gvmDir := home + "/.gvm/"
	if param == "all"{
		if err := os.RemoveAll(gvmDir); err != nil{
			log.Error( err)
			return
		}
	}
	files, err := ioutil.ReadDir(gvmDir)
	if err != nil{
		log.Error( err)
	}
	for _, v := range files{
		if strings.Contains(v.Name(), param){
			if err := os.RemoveAll(gvmDir + v.Name()); err != nil{
				log.Error( err)
				return
			}
		}
	}
}