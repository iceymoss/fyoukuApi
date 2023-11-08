package controllers

import (
	"fmt"
	"fyoukuApi/models"
	"regexp"

	beego "github.com/astaxie/beego"
)

// Operations about Users

type UserController struct {
	beego.Controller
}

// SaveRegister 用户注册功能
// @router /register/save [get]
func (u *UserController) SaveRegister() {
	var (
		mobile   string
		password string
		err      error
	)

	mobile = u.GetString("mobile")
	password = u.GetString("password")

	//校验手机号和密码的合法性
	if mobile == "" {
		u.Data["json"] = ReturnError(4001, "手机号不能为空")
		u.ServeJSON()
	}
	fmt.Println("mobile:", mobile)
	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	fmt.Println("结果：", isorno)
	if !isorno {
		u.Data["json"] = ReturnError(4002, "手机格式不正确")
		u.ServeJSON()
	}

	if password == "" {
		u.Data["json"] = ReturnError(4001, "密码不能为空")
		u.ServeJSON()
	}

	//查询账上是否存在
	isHasUser := models.GetUserByMobile(mobile)
	if isHasUser {
		u.Data["json"] = ReturnError(4005, "手机号已经注册")
		u.ServeJSON()
	} else {
		err = models.UserSave(mobile, MD5V(password))
		if err == nil {
			u.Data["json"] = ReturnSuccess(0, "注册账号成功", nil, 0)
			u.ServeJSON()
		} else {
			u.Data["json"] = ReturnError(5000, err)
			u.ServeJSON()
		}
	}
}

func (u *UserController) LoginDo() {
	mobile := u.GetString("mobile")
	password := u.GetString("password")

	if mobile == "" {
		u.Data["json"] = ReturnError(4001, "手机号不能为空")
		u.ServeJSON()
	}
	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isorno {
		u.Data["json"] = ReturnError(4002, "手机号格式不正确")
		u.ServeJSON()
	}
	if password == "" {
		u.Data["json"] = ReturnError(4003, "密码不能为空")
		u.ServeJSON()
	}
	uid, name := models.IsMobileLogin(mobile, MD5V(password))
	if uid != 0 {
		u.Data["json"] = ReturnSuccess(0, "登录成功", map[string]interface{}{"uid": uid, "username": name}, 1)
		u.ServeJSON()
	} else {
		u.Data["json"] = ReturnError(4004, "手机号或密码不正确")
		u.ServeJSON()
	}
}
