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

// mq消费者：客服发消息只需要向mq发送，消费者进行异步处理
func main() {
	beego.LoadAppConfig("ini", "../../conf/app.conf")
	defaultdb := beego.AppConfig.String("defaultdb")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", defaultdb, 30, 30)

	mq.Consumer("", "fyouku_send_message_user", callback)
}

// callback具体的义务逻辑：将消息写入数据库，即给用户发送消息
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
