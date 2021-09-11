package confgitlab

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/xanzy/go-gitlab"
)

func (s *Server) CreateRelease(folders ...string) {
	links := s.assertLinks(folders...)
	pid := s.env("CI_PROJECT_ID")
	opts := &gitlab.CreateReleaseOptions{
		Assets: &gitlab.ReleaseAssets{
			Links: links,
		},
	}

	s.createRelease(pid, opts)
}

func (s *Server) createRelease(pid string, opts *gitlab.CreateReleaseOptions) (*gitlab.Release, error) {

	rel, _, err := s.gitlab.Releases.CreateRelease(pid, opts)

	return rel, err
}

func (s *Server) fileUrl(filename string) string {
	// https://github.com/tangx/dnsx/releases/download/v1.0.3/dnsx_v1.0.3_Darwin_x86_64
	// https://git.example.com/releases/download/:group/:project_name/:tag/:filename

	_url := `releases/%s/%s/download/%s/%s`
	url := fmt.Sprintf(_url,
		s.env("CI_PROJECT_NAMESPACE"),
		s.env("CI_PROJECT_NAME"),
		s.env("CI_COMMIT_REF_NAME"),
		filename,
	)
	return url
}

func (s *Server) assertLinks(folders ...string) []*gitlab.ReleaseAssetLink {
	links := []*gitlab.ReleaseAssetLink{}

	for _, folder := range folders {
		dirEntries, err := os.ReadDir(folder)
		if err != nil {
			log.Printf("lookup %s failed\n", folder)
			continue
		}

		for _, entry := range dirEntries {
			// not support recursive walk
			if entry.IsDir() {
				continue
			}
			filename := entry.Name()
			filePath := filepath.Join(folder, filename)
			fileUrlPostfix := s.fileUrl(filename)
			fileReleaseUrl := filepath.Join(s.ReleasePrefix, fileUrlPostfix)

			// todo: use minio-go upload
			if s3upload(filename, filePath, fileUrlPostfix, fileReleaseUrl) {
				links = append(links, &gitlab.ReleaseAssetLink{
					Name: filename,
					URL:  fileReleaseUrl,
				})
			}
		}
	}

	return links
}

// s3upload copy file into s3 object service
func s3upload(names ...string) bool {
	fmt.Println(names)
	return true
}
