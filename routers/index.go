package routers

import (
	"test/modules/result"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {

	ret := result.Ok(gin.H{
		"message": "pong",
	})

	c.JSON(200, ret)
}
