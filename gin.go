package tools

import (
	"github.com/gin-gonic/gin"
)

func AllGinParams(c *gin.Context) gin.Params {
	var params gin.Params

	for _, p := range c.Params {
		params = append(params, gin.Param{Key: p.Key, Value: p.Value})
	}

	req := c.Request
	req.ParseForm()

	for k, v := range req.PostForm {
		if len(v) == 0 {
			continue
		}

		params = append(params, gin.Param{Key: k, Value: v[0]})
	}

	for k, v := range req.URL.Query() {
		if len(v) == 0 {
			continue
		}
		params = append(params, gin.Param{Key: k, Value: v[0]})
	}

	return params
}
