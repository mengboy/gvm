package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

type Reader struct {
	io.Reader
	Total   int64
	Current int64
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)

	r.Current += int64(n)
	percent := float64(r.Current*10000/r.Total) / 100
	Show(percent)
	return
}

// Show 显示进度
func Show(percent float64) {
	total := 50
	middle := int(percent) * total / 100.0
	arr := make([]string, total)
	for j := 0; j < total; j++ {
		if j < middle-1 {
			arr[j] = "-"
		} else if j == middle-1 {
			arr[j] = ">"
		} else {
			arr[j] = " "
		}
	}
	bar := fmt.Sprintf("[%s]", strings.Join(arr, ""))
	fmt.Printf("\r%s %% %.2f", bar, percent)
}

func DialTimeOutFunc(cTimeOut time.Duration) func(ctx context.Context, network, addr string) (net.Conn, error) {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return net.DialTimeout(network, addr, cTimeOut*time.Second)
	}
}

func DownloadFileProgress(url, backURL, filename string) error {
	home, err := GetUserHomePath()
	if err != nil {
		return err
	}
	dir := home + "/.gvm/"
	_, err = os.Stat(dir)
	if err != nil && !os.IsExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	httpCli := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			DialContext: DialTimeOutFunc(5),
		},
	}
	var r *http.Response
	r, err = httpCli.Get(url)
	if err != nil {
		fmt.Println("get from ", url, " failed")
		fmt.Println("try get from ", backURL)
		r, err = httpCli.Get(backURL)
		if err != nil {
			return err
		}
	}
	defer func() {
		_ = r.Body.Close()
	}()
	f, err := os.OpenFile(dir+filename, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	reader := &Reader{
		Reader: r.Body,
		Total:  r.ContentLength,
	}
	_, err = io.Copy(f, reader)
	if err != nil {
		return err
	}
	return nil

}
