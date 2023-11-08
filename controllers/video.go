package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego"
)

type VideoControllers struct {
	beego.Controller
}

// ChannelAdvert 频道-顶部广告
func (v *VideoControllers) ChannelAdvert() {
	channelId, _ := v.GetInt("channelId")
	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "频道id不能为空")
		v.ServeJSON()
	}

	num, adverts, err := models.GetChannelAdvert(channelId)
	if err != nil {
		v.Data["json"] = ReturnError(4004, "请求数据失败，情稍后再试")
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnSuccess(0, "success", adverts, num)
		v.ServeJSON()
	}
}

// GetChannelHotList 获取频道热播视频
func (v *VideoControllers) GetChannelHotList() {
	channelId, _ := v.GetInt("channelId")
	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "频道id不能为空")
		v.ServeJSON()
	}

	//获取当前频道热播视频
	num, videos, err := models.GetChannelHotList(channelId)
	if err != nil {
		v.Data["json"] = ReturnError(4004, "请求数据失败，情稍后再试")
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnSuccess(0, "success", videos, num)
		v.ServeJSON()
	}
}

// GetChannelRecommendRegionList 根据频道和地区获取视频列表
func (v *VideoControllers) GetChannelRecommendRegionList() {
	channelId, _ := v.GetInt("channelId")
	regionId, _ := v.GetInt("regionId")
	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道")
		v.ServeJSON()
	}
	if regionId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道地区")
		v.ServeJSON()
	}

	num, videos, err := models.GetChannelRecommendRegionList(channelId, regionId)
	if err != nil {
		v.Data["json"] = ReturnError(4004, "请求数据失败，情稍后再试")
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnSuccess(0, "success", videos, num)
		v.ServeJSON()
	}
}

// GetChannelTypeList 根据频道和类型推荐视频
func (v *VideoControllers) GetChannelTypeList() {
	channelId, _ := v.GetInt("channelId")
	typeId, _ := v.GetInt("typeId")
	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道")
		v.ServeJSON()
	}
	if typeId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定类型")
		v.ServeJSON()
	}

	num, videos, err := models.GetChannelTypeList(channelId, typeId)
	if err != nil {
		v.Data["json"] = ReturnError(4004, "请求数据失败，情稍后再试")
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnSuccess(0, "success", videos, num)
		v.ServeJSON()
	}
}
