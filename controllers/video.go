package controllers

import (
	"encoding/json"
	"fmt"
	"fyoukuApi/models"
	"fyoukuApi/services/es"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strconv"
	"strings"
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
	video, err := models.RedisGetVideoInfo(videoId)
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
	num, episodes, err := models.RedisVideoEpisodesList(videoId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", episodes, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "请求数据失败，请稍后重试~")
		v.ServeJSON()
	}
}

// UserVideos 用户视频管理-视频列表
func (v *VideoControllers) UserVideos() {
	//获取用户id
	uid, _ := v.GetInt("uid")
	if uid == 0 {
		v.Data["json"] = ReturnError(4002, "请登录账号")
		v.ServeJSON()
	}
	num, videos, err := models.GetUserVideo(uid)
	if err != nil {
		v.Data["json"] = ReturnError(4004, "您还没有发布作品")
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnSuccess(0, "success", videos, num)
		v.ServeJSON()
	}
}

// UploadVideo 上传视频文件，方法一：上传到服务器本地存储
func (v *VideoControllers) UploadVideo() {
	var (
		err error
	)
	title := make([]byte, 0)
	r := *v.Ctx.Request
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
	v.Ctx.WriteString(string(title))
}

// VideoImgSave 保存封面文件
func (v *VideoControllers) VideoImgSave() {
	var err error
	title := make([]byte, 0)
	r := *v.Ctx.Request
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
	var fileDir = "/Users/iceymoss/project/fyouku/fyouku/static/data/img/" + filename[0] + "." + filename[1]
	err = ioutil.WriteFile(fileDir, b, 0777)
	var playUrl = "/static/data/img/" + filename[0] + "." + filename[1]
	err = ioutil.WriteFile(fileDir, b, 0777)
	if err == nil {
		title, err = json.Marshal(ReturnSuccess(0, playUrl, nil, 1))
	} else {
		title, err = json.Marshal(ReturnError(5000, "上传失败，请联系客服"))
	}
	v.Ctx.WriteString(string(title))
}

// VideoSave 保存视频信息
func (v *VideoControllers) VideoSave() {
	playUrl := v.GetString("playUrl")
	title := v.GetString("title")
	subTitle := v.GetString("subTitle")
	channelId, _ := v.GetInt("channelId")
	typeId, _ := v.GetInt("typeId")
	regionId, _ := v.GetInt("regionId")
	uid, _ := v.GetInt("uid")
	aliyunVideoId := v.GetString("aliyunVideoId")
	fmt.Println("播放地址：", playUrl, "视频标题：", title)
	if uid == 0 {
		v.Data["json"] = ReturnError(4001, "请先登录")
		v.ServeJSON()
	}
	if playUrl == "" {
		v.Data["json"] = ReturnError(4002, "视频地址不能为空")
		v.ServeJSON()
	}

	err := models.VideoSave(title, subTitle, channelId, regionId, typeId, playUrl, uid, aliyunVideoId)
	if err != nil {
		v.Data["json"] = ReturnError(5000, err)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnSuccess(0, "success", nil, 1)
		v.ServeJSON()
	}
}

// VideoInfoUpdate 更新影视
func (v *VideoControllers) VideoInfoUpdate() {
	id, _ := v.GetInt("videoId")
	uid, _ := v.GetInt("uid")
	title := v.GetString("title")
	subTitle := v.GetString("subTitle")
	channelId, _ := v.GetInt("channelId")
	typeId, _ := v.GetInt("typeId")
	regionId, _ := v.GetInt("regionId")
	img := v.GetString("img")
	imgSub := v.GetString("img1")
	isEnd, _ := v.GetInt("is_end")
	if id == 0 {
		v.Data["json"] = ReturnError(40001, "请选择编辑的影视")
		v.ServeJSON()
	}
	if uid == 0 {
		v.Data["json"] = ReturnError(40003, "请登录")
		v.ServeJSON()
	}
	err := models.UpdateVideoInfo(id, title, subTitle, channelId, typeId, regionId, img, imgSub, isEnd)
	if err != nil {
		v.Data["json"] = ReturnError(5000, "更新失败"+err.Error())
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnSuccess(0, "success", nil, 1)
		v.ServeJSON()
	}
}

// EpisodesInfoUpdate 更新视频
func (v *VideoControllers) EpisodesInfoUpdate() {
	episodeId, _ := v.GetInt("episodeId")
	uid, _ := v.GetInt("uid")
	title := v.GetString("title")
	num, _ := v.GetInt("num")
	status, _ := v.GetInt("status")
	if episodeId == 0 {
		v.Data["json"] = ReturnError(4004, "请选中视频")
		v.ServeJSON()
	}
	if uid == 0 {
		v.Data["json"] = ReturnError(40003, "请登录")
		v.ServeJSON()
	}
	err := models.EpisodesInfoUpdate(episodeId, title, num, status)
	if err != nil {
		v.Data["json"] = ReturnError(5000, "更新失败")
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnSuccess(0, "success", nil, 1)
		v.ServeJSON()
	}
}

// EpisodeDelete 删除视频
func (v *VideoControllers) EpisodeDelete() {
	episodeId, _ := v.GetInt("episodeId")
	uid, _ := v.GetInt("uid")
	if episodeId == 0 {
		v.Data["json"] = ReturnError(4004, "请选中视频")
		v.ServeJSON()
	}
	if uid == 0 {
		v.Data["json"] = ReturnError(40003, "请登录")
		v.ServeJSON()
	}
	num, err := models.EpisodeDelete(episodeId)
	if err != nil {
		v.Data["json"] = ReturnError(5000, "删除失败"+err.Error())
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnSuccess(0, "success", nil, num)
		v.ServeJSON()
	}
}

// VideoDelete 删除影视
func (v *VideoControllers) VideoDelete() {
	videoId, _ := v.GetInt("videoId")
	uid, _ := v.GetInt("uid")
	if videoId == 0 {
		v.Data["json"] = ReturnError(4004, "请选中影视")
		v.ServeJSON()
	}
	if uid == 0 {
		v.Data["json"] = ReturnError(40003, "请登录")
		v.ServeJSON()
	}
	err := models.VideoDelete(videoId)
	if err != nil {
		v.Data["json"] = ReturnError(5000, "删除失败"+err.Error())
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnSuccess(0, "success", nil, 1)
		v.ServeJSON()
	}
}

// SendEs 导入ES脚本
func (vc *VideoControllers) SendEs() {
	_, data, _ := models.GetAllList()
	for _, v := range data {
		body := map[string]interface{}{
			"id":                   v.Id,
			"title":                v.Title,
			"sub_title":            v.SubTitle,
			"add_time":             v.AddTime,
			"img":                  v.Img,
			"img1":                 v.Img1,
			"episodes_count":       v.EpisodesCount,
			"is_end":               v.IsEnd,
			"channel_id":           v.ChannelId,
			"status":               v.Status,
			"region_id":            v.RegionId,
			"type_id":              v.TypeId,
			"episodes_update_time": v.EpisodesUpdateTime,
			"comment":              v.Comment,
			"user_id":              v.UserId,
			"is_recommend":         v.IsRecommend,
		}
		es.EsAdd("fyouku_video", "video-"+strconv.Itoa(v.Id), body)
	}
	vc.Data["json"] = ReturnSuccess(0, "success", nil, 1)
}
