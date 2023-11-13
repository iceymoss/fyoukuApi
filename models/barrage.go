package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Barrage 弹幕表结构
type Barrage struct {
	Id          int    //主键
	Content     string //弹幕内容
	CurrentTime int    //当前时间,秒
	AddTime     int64  //添加时间
	UserId      int    //添加用户
	Status      int    //弹幕状态
	EpisodesId  int    //弹幕视频
	VideoId     int    //归属视频
}

// BarrageData 弹幕返回结构
type BarrageData struct {
	Id          int    `json:"id"`
	Content     string `json:"content"`
	CurrentTime int    `json:"currentTime"`
}

func init() {
	orm.RegisterModel(new(Barrage))
}

// BarrageList 获取指定时间范围弹幕内容
func BarrageList(episodesId int, startTime int, endTime int) (int64, []BarrageData, error) {
	o := orm.NewOrm()
	var barrages []BarrageData
	num, err := o.Raw("SELECT id,content,`current_time` FROM barrage WHERE status=1 AND episodes_id=? AND `current_time`>=? AND `current_time`<? ORDER BY `current_time` ASC", episodesId, startTime, endTime).QueryRows(&barrages)
	return num, barrages, err
}

// SaveBarrage 保存弹幕
func SaveBarrage(episodesId int, videoId int, currentTime int, uid int, content string) error {
	o := orm.NewOrm()
	var barrage Barrage
	barrage.Content = content
	barrage.CurrentTime = currentTime
	barrage.AddTime = time.Now().Unix()
	barrage.UserId = uid
	barrage.Status = 1
	barrage.EpisodesId = episodesId
	barrage.VideoId = videoId
	_, err := o.Insert(&barrage)
	return err
}
