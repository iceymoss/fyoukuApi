package models

import "github.com/astaxie/beego/orm"

// Video 定义表结构
//type Video struct {
//	Id                 int
//	Title              string
//	SubTitle           string
//	AddTime            int64
//	Img                string
//	Img1               string
//	EpisodesCount      int
//	IsEnd              int
//	ChannelId          int
//	Status             int
//	RegionId           int
//	TypeId             int
//	EpisodesUpdateTime int64
//	Comment            int
//	UserId             int
//	IsRecommend        int
//}

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

//func init() {
//	orm.RegisterModel(&Video{})
//}

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
