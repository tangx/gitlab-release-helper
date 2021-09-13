package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
	"github.com/tangx/gitlab-release-helper/cmd/server/global"
)

func BaseRoute(base *gin.RouterGroup) {
	v0Route := base.Group("/v0")

	objectRoute := v0Route.Group("/object")

	objectRoute.GET("/*object", getHandler)
	objectRoute.PUT("/*object", putHandler)
}

type Params struct {
	Object string `uri:"object"`
}

func getHandler(c *gin.Context) {
	params := &Params{}

	err := ginbinder.ShouldBindRequest(c, params)
	if err != nil {
		c.String(http.StatusBadRequest, "bind params failed: %v", err)
		return
	}

	u, err := global.S3.PreSignedGetURL(params.Object)
	if err != nil {
		c.String(http.StatusInternalServerError, "internal error: %v", err)
		return
	}

	c.String(http.StatusFound, u.String())
}

func putHandler(c *gin.Context) {
	params := &Params{}

	err := ginbinder.ShouldBindRequest(c, params)
	if err != nil {
		c.String(http.StatusBadRequest, "bind params failed: %v", err)
		return
	}

	u, err := global.S3.PreSignedPutURL(params.Object, false)
	if err != nil {
		c.String(http.StatusInternalServerError, "internal error: %v", err)
		return
	}

	c.String(http.StatusTemporaryRedirect, u.String())
}
