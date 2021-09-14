package global

import (
	"github.com/go-jarvis/jarvis"
	"github.com/tangx/gitlab-release-helper/pkg/confgitlab"
	"github.com/tangx/gitlab-release-helper/pkg/httpclient"
)

var (
	GitlabHelper = &confgitlab.Server{}
	HttpClient   = &httpclient.Client{}
	app          = jarvis.App{
		Name: "Releaser",
	}
)

func init() {
	config := &struct {
		GitlabHelper *confgitlab.Server
		HttpClient   *httpclient.Client
	}{
		GitlabHelper: GitlabHelper,
		HttpClient:   HttpClient,
	}

	app.Conf(config)
}
