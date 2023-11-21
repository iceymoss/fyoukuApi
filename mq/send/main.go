package main

import (
	"encoding/json"
	"fmt"
	"fyoukuApi/models"
	"fyoukuApi/services/mq"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.LoadAppConfig("ini", "../../conf/app.conf")
	defaultdb := beego.AppConfig.String("defaultdb")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", defaultdb, 30, 30)

	mq.Consumer("", "fyouku_send_message_user", callback)
}

func callback(msg string) {
	type data struct {
		userId int
		msgId  int64
	}
	var userMsg data
	err := json.Unmarshal([]byte(msg), &userMsg)
	if err == nil {
		models.SendMessageUser(userMsg.userId, userMsg.msgId)
		fmt.Println("消费完成：异步写入MySQL成功")
	}
}
