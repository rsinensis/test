package routers

import (
	"test/modules/result"

	"github.com/gin-gonic/gin"
)

// @Summary ping测试
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"pong"},"msg":"ok"}"
// @Router /ping [get]
func ping(c *gin.Context) {

	ret := result.Ok("pong")

	c.JSON(200, ret)
}
