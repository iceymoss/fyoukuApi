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

// ChannelVideoList 根据传入参数获取视频列表
func (v *VideoControllers) ChannelVideoList() {
	//获取频道ID
	channelId, _ := v.GetInt("channelId")
	//获取频道地区ID
	regionId, _ := v.GetInt("regionId")
	//获取频道类型ID
	typeId, _ := v.GetInt("typeId")
	//获取状态
	end := v.GetString("end")
	//获取排序
	sort := v.GetString("sort")
	//获取页码信息
	limit, _ := v.GetInt("limit")
	offset, _ := v.GetInt("offset")

	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道")
		v.ServeJSON()
	}

	if limit == 0 {
		limit = 12
	}

	num, videos, err := models.GetChannelVideoList(channelId, regionId, typeId, end, sort, offset, limit)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "没有相关内容")
		v.ServeJSON()
	}
}

// VideoInfo 视频详细信息
func (v *VideoControllers) VideoInfo() {
	videoId, _ := v.GetInt("videoId")
	if videoId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定视频ID")
		v.ServeJSON()
	}
	video, err := models.GetVideoInfo(videoId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", video, 1)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "请求数据失败，请稍后重试~")
		v.ServeJSON()
	}
}

// VideoEpisodesList 获取视频剧集列表
func (v *VideoControllers) VideoEpisodesList() {
	videoId, _ := v.GetInt("videoId")
	if videoId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定视频ID")
		v.ServeJSON()
	}
	num, episodes, err := models.GetVideoEpisodesList(videoId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", episodes, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "请求数据失败，请稍后重试~")
		v.ServeJSON()
	}
}
