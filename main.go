package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
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

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
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

	go func() {
		// service connections
		fmt.Println(server.ListenAndServe())
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
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
