package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
)

var (
	localListCMD = &cobra.Command{
		Use:   "list",
		Short: "local has download",
		Run:   localList,
	}
)

func localList(cmd *cobra.Command, args []string)  {
	cache, err := getLocalDown()
	if err != nil{
		log.Error(err)
		return
	}
	fmt.Println("download: ")
	fmt.Println(strings.Join(cache, "\n"))
	goCMD := exec.Command("go", "version")
	goVersionCMDRes, err := goCMD.Output()
	if err != nil{
		log.Error("get now use failed", err)
		return
	}
	reg := regexp.MustCompile("go.*?go([0-9.]*)[\\s].*?")
	verionRes := reg.FindStringSubmatch(string(goVersionCMDRes))
	if len(verionRes) > 1{
		fmt.Println("now use: ", verionRes[1])
	}
}

func getLocalDown() ([]string, error) {
	home, err := GetUserHomePath()
	if err != nil{
		return nil, err
	}
	gvmDir := home + "/.gvm/"
	files, err := ioutil.ReadDir(gvmDir)
	if err != nil{
		return nil, err
	}
	if len(files) == 0{
		return nil, nil
	}
	reg := regexp.MustCompile("go([0-9.]*)[.].*?.tar.gz")
	versionList := []string{}
	for i := 0; i < len(files); i++{
		version := reg.FindStringSubmatch(files[i].Name())
		if len(version) > 1{
			versionList = append(versionList, version[1])
		}
	}
	return versionList, nil
}