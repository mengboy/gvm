package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

var remoteCMD = &cobra.Command{
	Use:   "remote",
	Short: "go version remote",
	Run:   remote,
}

func remote(cmd *cobra.Command, args []string) {
	f := ""
	if len(args) > 0 {
		f = args[0]
	}
	var vlist []string
	var err error
	if f == "g" {
		vlist, err = GetVersionList(GOLang)
		if err != nil {
			log.Error("get version remote failed form", GOLang)
			return
		}
	}
	vlist, err = GetVersionList(StudyGOLang)
	if err != nil {
		log.Error("get version remote failed form", StudyGOLang)
		return
	}
	for i := len(vlist) - 1; i >= 0 ; i--{
		fmt.Println(vlist[i])
		time.Sleep(10 * time.Millisecond)
	}
	return
}

func GetVersionList(url string) ([]string, error) {
	res, err := http.Get(url + "dl")
	if err != nil {
		return nil, err
	}
	pageContent, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return parseVersionList(pageContent)
}

func parseVersionList(pageContent []byte) ([]string, error) {
	stableReg, err := regexp.Compile(`<div class="toggleVisible" id="go(.*?)">`)
	if err != nil {
		return nil, err
	}
	stable := stableReg.FindAllStringSubmatch(string(pageContent), -1)
	unstableReg, err := regexp.Compile(`<div class="toggle" id="go(.*?)">`)
	if err != nil {
		return nil, err
	}
	unstable := unstableReg.FindAllStringSubmatch(string(pageContent), -1)
	versionList := make([]string, len(stable)+len(unstable))
	vListSource := append(stable, unstable...)
	for i := len(vListSource) - 1; i >= 0; i-- {
		versionList[i] = vListSource[i][1]
	}
	return versionList, nil
}
