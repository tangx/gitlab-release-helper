package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/tangx/gitlab-release-helper/cmd/releaser/cmd"
)

func main() {

	cmd.Execute()
}

func init() {
	if os.Getenv("DEBUG") != "" {
		logrus.SetReportCaller(true)
		logrus.SetLevel(logrus.DebugLevel)
	}
}
