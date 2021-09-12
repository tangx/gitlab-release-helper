package releaser

import (
	"log"
	"os"
	"path/filepath"

	"github.com/tangx/gitlab-release-helper/cmd/releaser/global"
	"github.com/tangx/gitlab-release-helper/pkg/confgitlab"
)

var s3 = global.S3Client
var githelper = global.GitlabHelper

func CreateRelease(folders ...string) {
	links := upload(folders...)
	tag, err := githelper.CreateRelease()
	if err != nil {
		log.Fatalf("Create Release failed: %v\n", err)
	}

	for _, link := range links {
		err := githelper.CreateReleaseLink(tag, link)
		if err != nil {
			log.Fatalf("create release link failed: %v\n", link)
		}
		log.Printf("create release link success: %v\n", link)
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

			_, err := s3.UploadFile(releaseName, filePath, true)
			if err != nil {
				log.Fatalf("upload file %s failed: %v", filePath, err)
			}

			links = append(links, confgitlab.AssertLink{Name: fileName, URL: releaseName})
		}
	}

	return links
}
