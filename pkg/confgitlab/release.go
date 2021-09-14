package confgitlab

import (
	"fmt"

	"github.com/xanzy/go-gitlab"
)

func (s *Server) CreateRelease(links ...AssertLink) (string, error) {
	// CI_JOB_TOKEN 只能创建 release, 可以携带 link
	// 不能单独创建 link
	releaseLinks := []*gitlab.ReleaseAssetLink{}
	for _, link := range links {
		// logrus.Debugln(link)
		releaseLinks = append(releaseLinks, &gitlab.ReleaseAssetLink{
			Name: link.Name,
			URL:  link.URL,
		})
	}

	pid := s.env("CI_PROJECT_ID")
	ref := s.env("CI_COMMIT_REF_NAME")
	rName := fmt.Sprintf("Release %s", ref)
	opts := &gitlab.CreateReleaseOptions{
		Name:    &rName,
		TagName: &ref,
		Ref:     &ref,
		Assets: &gitlab.ReleaseAssets{
			Links: releaseLinks,
		},
	}

	release, _, err := s.gitlab.Releases.CreateRelease(pid, opts)
	if err != nil {
		return "", err
	}

	return release.TagName, nil
}

func (s *Server) Object(filename string) string {
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

	return links
}
