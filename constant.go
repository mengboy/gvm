package main

import (
	"os/user"
)

const (
	StudyGOLang      = "https://studygolang.com/"
	GOLang           = "https://golang.org/"
	FileName         = "go%s.%s-%s.tar.gz" // version  os GOARCH
	StudyGOLangDLURL = "dl/golang/"
	GOLangDLURL      = "dl/"
)


func GetUserHomePath() (string, error) {
	currUser, err := user.Current()
	if err != nil {
		return "", err
	}

	return currUser.HomeDir, nil
}

const (
	GVMVersion  = "1.0.1"
)