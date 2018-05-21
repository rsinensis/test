package setting

import (
	"fmt"

	"github.com/gin-gonic/gin"
	ini "gopkg.in/ini.v1"
)

// Config is app config handle
var Config *ini.File

var Mode string

func SettingInit(mode string) {

	config, err := ini.Load(fmt.Sprintf("configs/config_%v.ini", gin.Mode()))
	if err != nil {
		fmt.Println(err)
		return
	}
	Config = config
	Mode = mode
}
