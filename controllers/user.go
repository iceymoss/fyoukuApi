package controllers

import (
	"encoding/json"
	"fmt"
	"fyoukuApi/models"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

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

// SendMessageDo 批量向用户发送通知消息
func (u *UserController) SendMessageDo() {
	uids := u.GetString("uids")
	content := u.GetString("content")

	if uids == "" {
		u.Data["json"] = ReturnError(4001, "请填写接收人~")
		u.ServeJSON()
	}
	if content == "" {
		u.Data["json"] = ReturnError(4002, "请填写发送内容")
		u.ServeJSON()
	}
	meaageId, err := models.SendMessageDo(content)
	if err != nil {
		u.Data["json"] = ReturnError(5000, "发生消息失败")
		u.ServeJSON()
	} else {
		uidConfig := strings.Split(uids, ",")
		for _, id := range uidConfig {
			//发生消息给用户，就是写入数据表
			uid, _ := strconv.Atoi(id)
			models.SendMessageUser(uid, meaageId)
		}
		u.Data["json"] = ReturnSuccess(0, "发送成功", "", 1)
		u.ServeJSON()
	}
}

// UploadVideo 上传视频文件，方法一：上传到服务器本地存储
func (u *UserController) UploadVideo() {
	var (
		err error
	)
	title := make([]byte, 0)
	r := *u.Ctx.Request
	//获取表单提交的数据
	uid := r.FormValue("uid")
	//获取文件流
	file, header, _ := r.FormFile("file")
	//转换文件流为二进制
	b, _ := ioutil.ReadAll(file)

	//生成文件名
	filename := strings.Split(header.Filename, ".")
	filename[0] = GetVideoName(uid)
	//文件保存的位置
	var fileDir = "/Users/iceymoss/project/fyouku/fyouku/static/video/" + filename[0] + "." + filename[1]
	//播放地址
	var playUrl = "/static/video/" + filename[0] + "." + filename[1]
	err = ioutil.WriteFile(fileDir, b, 0777)
	if err == nil {
		title, err = json.Marshal(ReturnSuccess(0, playUrl, nil, 1))
	} else {
		title, err = json.Marshal(ReturnError(5000, "上传失败，请联系客服"))
	}
	u.Ctx.WriteString(string(title))
}
