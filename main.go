package main

import (
	"flag"
	"fmt"
	"net/http"
	"test/models"
	"test/modules/setting"
	"test/routers"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	// BuildVersion from git tag
	BuildVersion string
	// BuildTime from make time
	BuildTime string
	// BuildMode from make mode
	BuildMode string
)

func main() {

	var v bool
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.Parse()

	if v {
		fmt.Println(fmt.Sprintf("\nBuildVersion: %v\n   BuildTime: %v\n   BuildMode: %v", BuildVersion, BuildTime, BuildMode))
		return
	}

	gin.SetMode(BuildMode)

	setting.SettingInit(BuildMode)

	engine := ginInit()

	models.EngineInit()

	routers.RouterInit(engine)

	port := setting.Config.Section("server").Key("port").MustInt(80)
	readTimeout := setting.Config.Section("server").Key("ReadTimeout").MustInt(10)
	writeTimeout := setting.Config.Section("server").Key("WriteTimeout").MustInt(10)
	maxHeaderBytes := setting.Config.Section("server").Key("MaxHeaderBytes").MustInt(1)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%v", port),
		Handler:        engine,
		ReadTimeout:    time.Duration(readTimeout) * time.Second,
		WriteTimeout:   time.Duration(writeTimeout) * time.Second,
		MaxHeaderBytes: maxHeaderBytes << 20,
	}

	fmt.Println(server.ListenAndServe())
}

func ginInit() *gin.Engine {

	engine := gin.New()

	// switch gin.Mode() {
	// case gin.ReleaseMode:
	// 	engine.Use(gin.Recovery())
	// case gin.TestMode:
	// case gin.DebugMode:
	// 	engine.Use(gin.Logger(), gin.Recovery())
	// }

	switch setting.Mode {
	case "release":
		engine.Use(gin.Recovery())
	case "debug":
	case "test":
		engine.Use(gin.Logger(), gin.Recovery())
	}

	return engine
}
