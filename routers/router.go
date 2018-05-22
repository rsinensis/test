package routers

import (
	_ "test/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// RouterInit creates routers
func RouterInit(engine *gin.Engine) {
	engine.GET("/ping", ping)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
