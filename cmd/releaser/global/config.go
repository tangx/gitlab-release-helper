package global

import (
	"github.com/go-jarvis/jarvis"
	"github.com/tangx/confs3"
	"github.com/tangx/gitlab-release-helper/pkg/confgitlab"
)

var (
	GitlabHelper = &confgitlab.Server{}
	app          = jarvis.App{
		Name: "Releaser",
	}
)

func init() {
	config := &struct {
		GitlabHelper *confgitlab.Server
		S3Client     *confs3.S3Client
	}{
		GitlabHelper: GitlabHelper,
	}

	app.Conf(config)
}
