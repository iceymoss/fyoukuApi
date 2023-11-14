package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Video 定义表结构
type Video struct {
	Id                 int
	Title              string
	SubTitle           string
	AddTime            int64
	Img                string
	Img1               string
	EpisodesCount      int
	IsEnd              int
	ChannelId          int
	Status             int
	RegionId           int
	TypeId             int
	EpisodesUpdateTime int64
	Comment            int
	UserId             int
	IsRecommend        int
}

type VideoData struct {
	Id            int    //id
	Title         string //标题
	SubTitle      string //副标题
	AddTime       int64  //添加时间
	Img           string //横版缩略图
	Img1          string //竖版缩略图
	EpisodesCount int    //集数
	IsEnd         int    //是否完结
	Comment       int    //评论数
}

type Episodes struct {
	Id            int
	Title         string
	AddTime       int64
	Num           int
	PlayUrl       string
	Comment       int
	AliyunVideoId string
}

func init() {
	orm.RegisterModel(&Video{})
}

// GetChannelHotList 查询频道热门播放列表
func GetChannelHotList(channelId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	videos := make([]VideoData, 0)
	num, err := o.Raw("SELECT id,title,sub_title,add_time, img,img1,episodes_count,is_end FROM video WHERE status=1 AND is_hot=1 AND channel_id=? ORDER BY episodes_update_time DESC LIMIT 9", channelId).QueryRows(&videos)
	return num, videos, err
}

func GetChannelRecommendRegionList(channelId, regionId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	videos := make([]VideoData, 0)
	num, err := o.Raw("SELECT id,title,sub_title,add_time,img,img1,episodes_count,is_end FROM video WHERE status=1 AND is_recommend=1 AND region_id=? AND channel_id=? ORDER BY episodes_update_time DESC LIMIT 9", regionId, channelId).QueryRows(&videos)
	return num, videos, err
}

func GetChannelTypeList(channelId, typeId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	videos := make([]VideoData, 0)
	num, err := o.Raw("select id, title, sub_title, add_time, img, img1, episodes_count, is_end from video where status = 1 and is_recommend = 1 and channel_id = ? and type_id = ? order by episodes_update_time desc limit 9", channelId, typeId).QueryRows(&videos)
	return num, videos, err
}

func GetChannelVideoList(channelId int, regionId int, typeId int, end string, sort string, offset int, limit int) (int64, []orm.Params, error) {
	o := orm.NewOrm()
	var videos []orm.Params

	qs := o.QueryTable("video")
	qs = qs.Filter("channel_id", channelId)
	qs = qs.Filter("status", 1)
	if regionId > 0 {
		qs = qs.Filter("region_id", regionId)
	}
	if typeId > 0 {
		qs = qs.Filter("type_id", typeId)
	}
	if end == "n" {
		qs = qs.Filter("is_end", 0)
	} else if end == "y" {
		qs = qs.Filter("is_end", 1)
	}
	if sort == "episodesUpdateTime" {
		qs = qs.OrderBy("-episodes_update_time") //最近更新
	} else if sort == "comment" {
		qs = qs.OrderBy("-comment") //评论数
	} else if sort == "addTime" {
		qs = qs.OrderBy("-add_time") //上映时间
	} else {
		qs = qs.OrderBy("-add_time")
	}
	nums, _ := qs.Values(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")
	qs = qs.Limit(limit, offset)
	_, err := qs.Values(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")

	return nums, videos, err
}

// GetVideoInfo 获取视频详情
func GetVideoInfo(videoId int) (Video, error) {
	o := orm.NewOrm()
	var video Video
	err := o.Raw("SELECT * FROM video WHERE id=? LIMIT 1", videoId).QueryRow(&video)
	return video, err
}

// GetVideoEpisodesList 获取视频剧集列表
func GetVideoEpisodesList(videoId int) (int64, []Episodes, error) {
	o := orm.NewOrm()
	var episodes []Episodes
	num, err := o.Raw("SELECT id,title,add_time,num,play_url,comment FROM video_episodes WHERE video_id=? order by num asc", videoId).QueryRows(&episodes)
	return num, episodes, err
}

// GetChannelTop 频道排行榜
func GetChannelTop(channelId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,img,img1,add_time,episodes_count,is_end,comment FROM video WHERE status=1 AND channel_id=? ORDER BY comment DESC LIMIT 10", channelId).QueryRows(&videos)
	return num, videos, err
}

// GetTypeTop 类型排行榜
func GetTypeTop(typeId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,img,img1,add_time,episodes_count,is_end,comment FROM video WHERE status=1 AND type_id=? ORDER BY comment DESC LIMIT 10", typeId).QueryRows(&videos)
	return num, videos, err
}

func GetUserVideo(uid int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,img,img1,add_time,episodes_count, is_end FROM video WHERE user_id=? ORDER BY add_time DESC", uid).QueryRows(&videos)
	return num, videos, err
}

func VideoSave(title, subTitle string, channelId, regionId, typeId int, playUrl string, uid int, aliyunVideoId string) error {
	o := orm.NewOrm()
	t := time.Now().Unix()
	video := Video{
		Title:              title,
		SubTitle:           subTitle,
		AddTime:            t,
		Img:                "",
		Img1:               "",
		EpisodesCount:      1,
		IsEnd:              1,
		ChannelId:          channelId,
		Status:             1,
		RegionId:           regionId,
		TypeId:             typeId,
		EpisodesUpdateTime: t,
		Comment:            0,
		UserId:             uid,
	}
	videoId, err := o.Insert(&video)
	if err != nil {
		return err
	}
	if aliyunVideoId != "" {
		playUrl = ""
	}
	_, err = o.Raw("INSERT INTO video_episodes (title,add_time,num,video_id,play_url,status,comment,aliyun_video_id) VALUES (?,?,?,?,?,?,?,?)", subTitle, t, 1, videoId, playUrl, 1, 0, aliyunVideoId).Exec()
	//fmt.Println(err)
	return nil
}
