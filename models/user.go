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

type UserInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	AddTime int64  `json:"addTime"`
	Avatar  string `json:"avatar"`
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

// IsMobileLogin 登录功能
func IsMobileLogin(mobile string, password string) (int, string) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("mobile", mobile).Filter("password", password).One(&user)
	if err == orm.ErrNoRows {
		return 0, ""
	} else if err == orm.ErrMissPK {
		return 0, ""
	}
	return user.Id, user.Name
}
