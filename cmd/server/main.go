package main

import (
	"github.com/tangx/gitlab-release-helper/cmd/server/apis"
	"github.com/tangx/gitlab-release-helper/cmd/server/global"
)

func main() {
	global.Server.Run()
}

func init() {
	global.Server.RegisterRoutes(apis.BaseRoute)
}
