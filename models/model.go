package models

import (
	"fmt"
	"test/modules/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func EngineInit() {

	ty := setting.Config.Section("database").Key("type").String()
	host := setting.Config.Section("database").Key("host").String()
	port := setting.Config.Section("database").Key("port").String()
	user := setting.Config.Section("database").Key("user").String()
	password := setting.Config.Section("database").Key("password").String()
	name := setting.Config.Section("database").Key("name").String()

	var err error
	engine, err = xorm.NewEngine(ty, fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&loc=Asia%%2FShanghai", user, password, host, port, name))
	if err != nil {
		fmt.Println(fmt.Sprintf("create database engine error:%v", err))
		return
	}

	err = engine.Ping()
	if err != nil {
		fmt.Println(fmt.Sprintf("ping database engine error:%v", err))
	}

	switch setting.Mode {
	case "release":

	case "debug":
	case "test":
		engine.ShowSQL(true)
	}

	prefix := setting.Config.Section("database").Key("prefix").String()
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, prefix)

	engine.SetTableMapper(tbMapper)

	maxIdleConns := setting.Config.Section("database").Key("MaxIdleConns").MustInt(10)
	maxOpenConns := setting.Config.Section("database").Key("MaxOpenConns").MustInt(100)

	engine.SetMaxIdleConns(maxIdleConns)
	engine.SetMaxOpenConns(maxOpenConns)
}
