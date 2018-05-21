package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"test/routers"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
)

var (
	// BuildVersion from git tag
	BuildVersion string
	// BuildTime from make time
	BuildTime string
	// BuildMode from make mode
	BuildMode string
)

// Config is app config handle
var Config *ini.File

func main() {

	var v bool
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.Parse()

	if v {
		fmt.Println(fmt.Sprintf(" BuildVersion:%v\n BuildTime:%v\n BuildMode:%v", BuildVersion, BuildTime, BuildMode))
		return
	}

	gin.SetMode(BuildMode)

	config, err := ini.Load(fmt.Sprintf("configs/config_%v.ini", gin.Mode()))
	if err != nil {
		log.Fatal(err)
		return
	}
	Config = config

	engine := ginInit()

	routers.RouterInit(engine)

	addr := Config.Section("server").Key("addr").MustInt(80)
	readTimeout := Config.Section("server").Key("ReadTimeout").MustInt(10)
	writeTimeout := Config.Section("server").Key("WriteTimeout").MustInt(10)
	maxHeaderBytes := Config.Section("server").Key("MaxHeaderBytes").MustInt(1)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%v", addr),
		Handler:        engine,
		ReadTimeout:    time.Duration(readTimeout) * time.Second,
		WriteTimeout:   time.Duration(writeTimeout) * time.Second,
		MaxHeaderBytes: maxHeaderBytes << 20,
	}

	server.ListenAndServe()
}

func ginInit() *gin.Engine {

	engine := gin.New()

	switch gin.Mode() {
	case gin.ReleaseMode:
		engine.Use(gin.Recovery())
	case gin.TestMode:
	case gin.DebugMode:
		engine.Use(gin.Logger(), gin.Recovery())
	}

	return engine
}
