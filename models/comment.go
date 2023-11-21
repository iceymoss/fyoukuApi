package models

import (
	"encoding/json"
	"fyoukuApi/services/mq"
	"github.com/astaxie/beego/orm"
	"time"
)

type Comment struct {
	Id          int
	Content     string
	AddTime     int64
	UserId      int
	Stamp       int
	Status      int
	PraiseCount int
	EpisodesId  int
	VideoId     int
}

func init() {
	orm.RegisterModel(new(Comment))
}

func GetCommentList(episodesId int, offset int, limit int) (int64, []Comment, error) {
	o := orm.NewOrm()
	var comments []Comment
	num, _ := o.Raw("SELECT id FROM comment WHERE status=1 AND episodes_id=?", episodesId).QueryRows(&comments)
	_, err := o.Raw("SELECT id,content,add_time,user_id,stamp,praise_count,episodes_id FROM comment WHERE status=1 AND episodes_id=? ORDER BY add_time DESC LIMIT ?,?", episodesId, offset, limit).QueryRows(&comments)
	return num, comments, err
}

func SaveComment(episodesId int, videoId int, uid int, content string) error {
	o := orm.NewOrm()
	var comment Comment
	comment.EpisodesId = episodesId
	comment.UserId = uid
	comment.VideoId = videoId
	comment.Content = content
	comment.Stamp = 0
	comment.Status = 1
	comment.AddTime = time.Now().Unix()
	_, err := o.Insert(&comment)
	if err != nil {
		return err
	}
	o.Raw("UPDATE video SET comment=comment+1 WHERE id=?", videoId).Exec()
	//修改视频剧集的评论数
	o.Raw("UPDATE video_episodes SET comment=comment+1 WHERE id=?", episodesId).Exec()

	//逻辑：排行榜是依据评论数来设计的，为了实时获取排行榜数据，新增评论的视频id通过mq来转发到Redis进行累加
	//更新redis排行榜 - 通过MQ来实现
	//创建一个简单模式的MQ
	//把要传递的数据转换为json字符串
	videoObj := map[string]int{
		"VideoId": videoId,
	}
	videoJson, _ := json.Marshal(videoObj)
	mq.Publish("", "fyouku_top", string(videoJson))

	//延迟增加评论数
	videoCountObj := map[string]int{
		"VideoId":    videoId,
		"EpisodesId": episodesId,
	}
	videoCountJson, _ := json.Marshal(videoCountObj)
	mq.PublishDlx("fyouku.comment.count", string(videoCountJson))

	return err
}
