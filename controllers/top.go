package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego"
)

type TopControllers struct {
	beego.Controller
}

// ChannelTop 根据频道获取排行榜
func (t *TopControllers) ChannelTop() {
	channelId, _ := t.GetInt("channelId")
	if channelId == 0 {
		t.Data["json"] = ReturnError(4001, "频道id不能为空")
		t.ServeJSON()
	}
	num, videos, err := models.RedisGetChannelTop(channelId)
	if err != nil {
		t.Data["json"] = ReturnError(4004, "没有相关内容")
		t.ServeJSON()
	} else {
		t.Data["json"] = ReturnSuccess(0, "success", videos, num)
		t.ServeJSON()
	}
}

// TypeTop 根据频道获取排行榜
func (t *TopControllers) TypeTop() {
	typeId, _ := t.GetInt("typeId")
	if typeId == 0 {
		t.Data["json"] = ReturnError(4001, "频道类型不能为空")
		t.ServeJSON()
		return
	}
	num, videos, err := models.RedisGetTypeTop(typeId)
	if err != nil {
		t.Data["json"] = ReturnError(4004, "没有相关内容")
		t.ServeJSON()
	} else {
		t.Data["json"] = ReturnSuccess(0, "success", videos, num)
		t.ServeJSON()
	}
}
