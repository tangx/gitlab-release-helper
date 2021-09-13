package global

import (
	"github.com/go-jarvis/jarvis"
	"github.com/tangx/confs3"
	"github.com/tangx/gitlab-release-helper/pkg/confgin"
)

var (
	S3     = &confs3.S3Client{}
	Server = &confgin.Server{}
)

func init() {
	app := jarvis.App{
		Name: "Server",
	}
	config := &struct {
		S3     *confs3.S3Client
		Server *confgin.Server
	}{
		S3:     S3,
		Server: Server,
	}

	app.Conf(config)
}
