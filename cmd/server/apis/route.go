package apis

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
	"github.com/tangx/gitlab-release-helper/cmd/server/global"
	"github.com/tangx/gitlab-release-helper/pkg/confgin/response"
)

func BaseRoute(base *gin.RouterGroup) {
	v0Route := base.Group("/v0")

	objectRoute := v0Route.Group("/object")

	objectRoute.GET("/*object", getHandler)
	objectRoute.POST("/*object", putHandler)
}

type Params struct {
	Object string `uri:"object"`
}

func getHandler(c *gin.Context) {
	params := &Params{}

	err := ginbinder.ShouldBindRequest(c, params)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	_object := strings.Trim(params.Object, "/")
	u, err := global.S3.PreSignedGetURL(_object)
	if err != nil {
		// c.String(http.StatusInternalServerError, "internal error: %v", err)
		response.Common(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusFound, u.String())
}

func putHandler(c *gin.Context) {
	params := &Params{}

	err := ginbinder.ShouldBindRequest(c, params)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	_object := strings.Trim(params.Object, "/")
	u, err := global.S3.PreSignedPutURL(_object, false)
	if err != nil {
		msg := fmt.Sprintf("internal error: %v", err)
		response.Error(c, http.StatusInternalServerError, msg)
		return
	}

	// 这里其实不应该直接使用 c.Redirect
	// client 会把 body 发过来， 是一种浪费。
	// 应该拆分为两步
	// 1. c -> server : 请求一个 presign url
	// 2. c -> s3 : put upload 文件
	// c.Redirect(http.StatusTemporaryRedirect, u.String())
	data := Data{
		PermanentiLink:    filepath.Join(c.Request.Host, _object),
		TemporaryRedirect: u.String(),
	}
	response.Common(c, http.StatusOK, data)
}

type Data struct {
	PermanentiLink    string
	TemporaryRedirect string
}
