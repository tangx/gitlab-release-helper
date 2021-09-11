package upload

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/tangx/gitlab-release-helper/cmd/releaser/global"
)

var (
	s3        = global.S3Client
	githelper = global.GitlabHelper
)

func Upload(folders ...string) {

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
			fileURL := githelper.FileUrl(fileName)

			fmt.Println([]string{fileURL, filePath})

			info, err := s3.UploadFile(fileURL, filePath, true)
			if err != nil {
				log.Printf("upload file %s failed: %v", fileName, err)
			}

			fmt.Println(info.Location)
		}
	}
}
