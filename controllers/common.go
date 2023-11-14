package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type CommonController struct {
	beego.Controller
}

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Items interface{} `json:"items"`
	Count int64       `json:"count"`
}

func ReturnSuccess(code int, msg interface{}, items interface{}, count int64) (json *JsonStruct) {
	json = &JsonStruct{
		Code:  code,
		Msg:   msg,
		Items: items,
		Count: count,
	}
	return
}

func ReturnError(code int, msg interface{}) (json *JsonStruct) {
	json = &JsonStruct{
		Code: code,
		Msg:  msg,
	}
	return
}

func MD5V(password string) string {
	h := md5.New()
	h.Write([]byte(password + beego.AppConfig.String("md5code")))
	return hex.EncodeToString(h.Sum(nil))
}

// DateFormat 格式化时间
func DateFormat(times int64) string {
	video_time := time.Unix(times, 0)
	return video_time.Format("2006-01-02")
}

// GetVideoName 视频文件名生成函数
func GetVideoName(uid string) string {
	//用户ID+精确到毫秒的时间戳
	h := md5.New()
	h.Write([]byte(uid + strconv.FormatInt(time.Now().UnixNano(), 10)))
	return hex.EncodeToString(h.Sum(nil))
}
