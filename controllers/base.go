package controllers

import (
	"fyoukuApi/models"
	beego "github.com/astaxie/beego"
)

// Operations about Users

type BaseController struct {
	beego.Controller
}

// ChannelRegion 获取频道的地区列表
func (b *BaseController) ChannelRegion() {
	channelId, _ := b.GetInt("channelId")
	if channelId == 0 {
		b.Data["json"] = ReturnError(4001, "必须指定频道")
		b.ServeJSON()
	}

	num, regions, err := models.GetChannelRegionList(channelId)
	if err != nil {
		b.Data["json"] = ReturnError(4004, "没有相关内容")
		b.ServeJSON()
	} else {
		b.Data["json"] = ReturnSuccess(0, "success", regions, num)
		b.ServeJSON()
	}
}

func (b *BaseController) ChannelType() {
	channelId, _ := b.GetInt("channelId")
	if channelId == 0 {
		b.Data["json"] = ReturnError(4001, "必须指定频道")
		b.ServeJSON()
	}

	num, types, err := models.GetChannelType(channelId)
	if err == nil {
		b.Data["json"] = ReturnSuccess(0, "success", types, num)
		b.ServeJSON()
	} else {
		b.Data["json"] = ReturnError(4004, "没有相关内容")
		b.ServeJSON()
	}
}
