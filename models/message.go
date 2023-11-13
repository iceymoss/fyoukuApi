package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Message struct {
	Id      int
	Content string
	AddTime int64
}

type MessageUser struct {
	Id        int
	MessageId int64
	AddTime   int64
	Status    int
	UserId    int
}

func init() {
	orm.RegisterModel(new(Message), new(MessageUser))
}

func SendMessageDo(content string) (int64, error) {
	o := orm.NewOrm()
	msg := Message{
		Content: content,
		AddTime: time.Now().Unix(),
	}
	return o.Insert(&msg)
}

func SendMessageUser(uid int, msgId int64) (int64, error) {
	o := orm.NewOrm()
	//获取到msg
	var msgUser MessageUser
	msgUser.UserId = uid
	msgUser.MessageId = msgId
	msgUser.AddTime = time.Now().Unix()
	msgUser.Status = 1
	return o.Insert(&msgUser)
}
