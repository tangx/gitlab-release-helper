package releaser

import (
	"log"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/tangx/gitlab-release-helper/cmd/releaser/global"
	"github.com/tangx/gitlab-release-helper/pkg/confgitlab"
)

var githelper = global.GitlabHelper
var httpclient = global.HttpClient

func CreateRelease(folders ...string) {
	links := upload(folders...)
	_, err := githelper.CreateRelease(links...)
	if err != nil {
		logrus.Fatalf("Create Release failed: %v\n", err)
		return
	}
	logrus.Info("Create Release Success")
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
			object := githelper.Object(fileName)

			// upload
			// 1. get redirect link and permernt link
			permlink, s3link, err := httpclient.PresignPutURL(object)
			if err != nil {
				logrus.Fatalf("generate presign put url failed: %v", err)
			}
			logrus.Debugln(permlink)
			// 2. put file into redirect link
			err = httpclient.PutFile(s3link, filePath)
			// fmt.Println(fileName)
			if err != nil {
				logrus.Fatalf("upload file %s to %s failed: %v", filePath, s3link, err)
			}

			// 3. create gitlab assert link
			links = append(links, confgitlab.AssertLink{Name: fileName, URL: permlink})
		}
	}

	return links
}
