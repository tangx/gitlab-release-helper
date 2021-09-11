package global

import "github.com/tangx/gitlab-release-helper/pkg/confgitlab"

var (
	GitlabHelper = confgitlab.Server{
		ReleasePrefix: "https://git-dl.example.com/xxxx",
	}
)

func init() {
	GitlabHelper.Init()
}
