package releaser

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/tangx/gitlab-release-helper/cmd/releaser/global"
	"github.com/tangx/gitlab-release-helper/pkg/confgitlab"
)

var githelper = global.GitlabHelper

func CreateRelease(folders ...string) {
	links := upload(folders...)
	_, err := githelper.CreateRelease(links...)
	if err != nil {
		log.Fatalf("Create Release failed: %v\n", err)
	}
}

func upload(folders ...string) []confgitlab.AssertLink {

	links := []confgitlab.AssertLink{}

	for _, folder := range folders {
		entries, err := os.ReadDir(folder)
		if err != nil {
			log.Printf("walk folder %s failed: %v\n", folder, err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			fileName := entry.Name()
			filePath := filepath.Join(folder, fileName)
			releaseName := githelper.ReleaseName(fileName)

			// upload
			// 1. get redirect link and permernt link
			permlink, s3link := getLinks(releaseName)
			// 2. put file into redirect link
			err := putFile(filePath, s3link)
			if err != nil {
				continue
			}

			// 3. create gitlab assert link
			links = append(links, confgitlab.AssertLink{Name: fileName, URL: permlink})
		}
	}

	return links
}

func getLinks(u string) (permlink string, s3link string) {
	resp, err := http.Post(u, "", nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	r := &ResponseData{}
	err = json.Unmarshal(data, r)
	if err != nil {
		return
	}

	return r.Data.PermanentiLink, r.Data.TemporaryRedirect
}

func putFile(file string, s3link string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	req, _ := http.NewRequest("PUT", s3link, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "multipart/form-data")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

type ResponseData struct {
	Status int
	Error  string
	Data   struct {
		PermanentiLink    string
		TemporaryRedirect string
	}
}
