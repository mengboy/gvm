package main

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetVersionList(t *testing.T) {
	fmt.Println(GetVersionList(GOLang))
}

func TestOSENV(t *testing.T) {
	t.Log(runtime.GOARCH)
	t.Log(runtime.GOOS)
}

func TestFileList(t *testing.T)  {
	getLocalDown()
}

func TestLocalList(t *testing.T)  {
	localList(nil, nil)
}

func TestDownloadFileProgress(t *testing.T) {
	filename := fmt.Sprintf(FileName, "1.15.2", "darwin", "amd64")
	fmt.Println(DownloadFileProgress("https://studygolang.com/" + OSDLURL + filename, filename))
}