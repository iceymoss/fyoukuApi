package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// User 表结构
type User struct {
	Id       int
	Name     string
	Password string
	Status   int
	Mobile   string
	AddTime  int64
	Avatar   string
}

func init() {
	orm.RegisterModel(new(User))
}

// GetUserByMobile 查询用户
func GetUserByMobile(mobile string) bool {
	o := orm.NewOrm()
	user := User{Mobile: mobile}
	err := o.Read(&user, "mobile")
	if err == orm.ErrNoRows {
		return false
	} else if err == orm.ErrMissPK {
		return false
	}
	return true
}

func UserSave(mobile string, password string) error {
	o := orm.NewOrm()
	user := User{
		Name:     "",
		Password: password,
		Status:   1,
		Mobile:   mobile,
		AddTime:  time.Now().Unix(),
	}
	_, err := o.Insert(&user)
	return err
}
