package routers

import "github.com/gin-gonic/gin"

// RouterInit creates routers
func RouterInit(engine *gin.Engine) {
	engine.GET("/ping", ping)
}
