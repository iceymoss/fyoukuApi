package main

import (
	"fmt"
	"fyoukuApi/command"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net/http"
)

func initConfig() {
	beego.LoadAppConfig("ini", "./conf/app.conf")
	defaultdb := beego.AppConfig.String("defaultdb")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", defaultdb, 30, 30)
}

func main() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("\n%v\n", err)
			}
		}()
		_ = http.ListenAndServe(":7161", nil)
	}()
	command.Execute()
}
