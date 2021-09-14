package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)

func Test_put(t *testing.T) {
	f, err := os.Open("main_test.go") // 创建文件句柄对象
	Panic(err)
	defer f.Close()
	body, _ := ioutil.ReadAll(f)

	now := time.Now().Unix()
	req, err := http.NewRequest(
		"PUT",
		`http://127.0.0.1:8088/app/v0/object/tmp/demo/README2.md`+fmt.Sprint(now),
		bytes.NewBuffer(body), // 问题在这里， 好像不能直接传 文件句柄 对象。
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "multipart/form-data")

	proxyUrl, _ := url.Parse(`http://127.0.0.1:8080`)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}

	resp, err := client.Do(req)
	Panic(err)
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	Panic(err)
	fmt.Printf("%s\n", data)
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func Test_get(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8088/app/v0/object/tmp/demo/README2.md", nil)
	Panic(err)

	// client := &http.Client{}
	proxyUrl, _ := url.Parse(`http://127.0.0.1:8080`)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
		// CheckRedirect: ,
	}

	client.Do(req)
	resp, err := client.Do(req)
	// resp, err := http.Get("http://127.0.0.1:8089/baidu")
	Panic(err)
	defer resp.Body.Close()

	ReadBody(resp)
}

func ReadBody(resp *http.Response) {
	data, err := ioutil.ReadAll(resp.Body)
	Panic(err)

	fmt.Printf("%s\n", data)
}
