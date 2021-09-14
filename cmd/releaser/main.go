package main

import (
	"github.com/tangx/gitlab-release-helper/cmd/releaser/cmd"
)

func main() {

	cmd.Execute()
}

func init() {
	// logrus.SetReportCaller(true)
	// logrus.SetLevel(logrus.DebugLevel)
}
