package response

import "github.com/gin-gonic/gin"

func Common(c *gin.Context, status int, data interface{}) {
	resp := Response{
		Status: status,
		Data:   data,
	}

	c.JSON(status, resp)
}

func Error(c *gin.Context, status int, err error) {
	resp := Response{
		Status: status,
		Error:  err,
	}

	c.JSON(status, resp)
}

type Response struct {
	Status int
	Data   interface{}
	Error  error
}
