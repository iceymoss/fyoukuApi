package models

import (
	"encoding/json"
	"fmt"
	redisClient "fyoukuApi/services/redis"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"strconv"
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
	Status        int
	AliyunVideoId string
}

type VideoEpisodes struct {
	Id            int
	Title         string
	AddTime       int64
	Num           int
	PlayUrl       string
	Comment       int
	Status        int
	AliyunVideoId string
}

func init() {
	orm.RegisterModel(&Video{}, &VideoEpisodes{})
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

func RedisGetVideoInfo(videoId int) (Video, error) {
	//使用Redis做缓存，先到Redis里面查询，没有查询到就去MySQL查询，然后写入到Redis
	var video Video
	conn := redisClient.PoolConnect()
	defer conn.Close()
	//构建key
	key := "video:id:" + strconv.Itoa(videoId)
	//判断Redis是否存在
	exists, err := redis.Bool(conn.Do("exists", key))
	if exists { //存在
		res, _ := redis.Values(conn.Do("hgetall", key))
		err = redis.ScanStruct(res, &video)
	} else { //不存在， 到MySQL里面查询
		o := orm.NewOrm()
		err = o.Raw("SELECT * FROM video WHERE id=? LIMIT 1", videoId).QueryRow(&video)
		if err == nil {
			//保存redis
			_, err = conn.Do("hmset", redis.Args{key}.AddFlat(video)...)
			if err == nil {
				conn.Do("expire", key, 86400) //设置过期时间
			}
		}
	}
	return video, nil
}

// UpdateVideoInfo 更新影视信息
func UpdateVideoInfo(id int, title string, subTitle string, channelId int, typeId int, regionId int, img string, imgSub string, isEnd int) error {
	o := orm.NewOrm()
	var video Video
	err := o.Raw("SELECT * FROM video WHERE id=? LIMIT 1", id).QueryRow(&video)
	if err != nil {
		return err
	}

	if title != "" {
		video.Title = title
	}
	fmt.Println("subtitle: ", subTitle)
	if subTitle != "" {
		video.SubTitle = subTitle
	}
	if channelId != 0 {
		video.ChannelId = channelId
	}
	if typeId != 0 {
		video.TypeId = typeId
	}
	if regionId != 0 {
		video.RegionId = regionId
	}
	if img != "" {
		video.Img = img
	}
	if imgSub != "" {
		video.Img1 = imgSub
	}
	video.IsEnd = isEnd
	_, err = o.Update(&video)
	return err
}

func EpisodesInfoUpdate(episodeId int, title string, num int, status int) error {
	o := orm.NewOrm()
	var episode VideoEpisodes
	err := o.Raw("select * from fyouku.video_episodes where id = ? LIMIT 1", episodeId).QueryRow(&episode)
	if err != nil {
		return err
	}
	fmt.Println("data:", episode)
	if title != "" {
		episode.Title = title
	}
	if num > 0 {
		episode.Num = num
	}
	episode.Status = status
	_, err = o.Update(&episode)
	fmt.Println("err: ", err)
	return err
}

func EpisodeDelete(episodesId int) (int64, error) {
	o := orm.NewOrm()
	var episode VideoEpisodes
	episode.Id = episodesId
	return o.Delete(&episode)
}

func VideoDelete(videoId int) error {
	o := orm.NewOrm()
	var video Video
	video.Id = videoId
	_, err := o.Delete(&video)
	if err != nil {
		return err
	}
	_, err = o.Raw("delete from video_episodes where video_id = ?", videoId).Exec()
	return err
}

// GetVideoEpisodesList 获取视频剧集列表
func GetVideoEpisodesList(videoId int) (int64, []Episodes, error) {
	o := orm.NewOrm()
	var episodes []Episodes
	num, err := o.Raw("SELECT id,title,add_time,num,play_url,comment,aliyun_video_id  FROM video_episodes WHERE video_id=? order by num asc", videoId).QueryRows(&episodes)
	return num, episodes, err
}

// RedisVideoEpisodesList redis作为缓存提高系统并发和稳定性，将影视的所有剧集存储到redis的list中
func RedisVideoEpisodesList(videoId int) (int64, []Episodes, error) {
	var episodes []Episodes
	var err error
	var num int64
	conn := redisClient.PoolConnect()
	defer conn.Close()

	key := "video:episodes:videoId:" + strconv.Itoa(videoId)
	exists, err := redis.Bool(conn.Do("exists", key))
	if exists {
		num, err = redis.Int64(conn.Do("llen", key))
		if err == nil {
			values, _ := redis.Values(conn.Do("lrange", key, "0", "-1"))
			var episodeInfo Episodes
			for _, value := range values {
				err = json.Unmarshal(value.([]byte), &episodeInfo) //value.([]byte) interface{}类型断言
				if err == nil {
					episodes = append(episodes, episodeInfo)
				}
			}
		}
	} else {
		o := orm.NewOrm()
		num, err = o.Raw("SELECT id,title,add_time,num,play_url,comment,aliyun_video_id  FROM video_episodes WHERE video_id=? order by num asc", videoId).QueryRows(&episodes)
		//将数组写到Redis的list中
		for _, episode := range episodes {
			jsonValue, err := json.Marshal(episode)
			if err == nil {
				conn.Do("rpush", key, jsonValue)
			}
		}
		conn.Do("expire", key, 86400*time.Second)
	}
	return num, episodes, err
}

// GetChannelTop 频道排行榜
func GetChannelTop(channelId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,img,img1,add_time,episodes_count,is_end,comment FROM video WHERE status=1 AND channel_id=? ORDER BY comment DESC LIMIT 10", channelId).QueryRows(&videos)
	return num, videos, err
}

// RedisGetChannelTop 增加缓存redis-频道排行榜
func RedisGetChannelTop(channelId int) (int64, []VideoData, error) {
	//基于评论热度进行排行，这里使用zset数据类型
	var videos []VideoData
	var num int64
	var err error
	conn := redisClient.PoolConnect()
	defer conn.Close()

	key := "video:top:channel:channelId:" + strconv.Itoa(channelId)
	exists, err := redis.Bool(conn.Do("exists", key))
	if exists {
		num = 0
		//zrevrange第一条获取到的是id,第二条是分数
		res, _ := redis.Values(conn.Do("zrevrange", key, "0", "10", "WITHSCORES"))
		for k, v := range res {
			if k%2 == 0 {
				videoId, err := strconv.Atoi(string(v.([]byte)))
				videosInfo, err := RedisGetVideoInfo(videoId)
				if err == nil {
					video := VideoData{
						Id:            videosInfo.Id,
						Title:         videosInfo.Title,
						SubTitle:      videosInfo.SubTitle,
						AddTime:       videosInfo.AddTime,
						Img:           videosInfo.Img,
						Img1:          videosInfo.Img1,
						EpisodesCount: videosInfo.EpisodesCount,
						IsEnd:         videosInfo.IsEnd,
						Comment:       videosInfo.Comment,
					}
					videos = append(videos, video)
					num++
				}
			}
		}
	} else {
		o := orm.NewOrm()
		num, err = o.Raw("SELECT id,title,sub_title,img,img1,add_time,episodes_count,is_end,comment FROM video WHERE status=1 AND channel_id=? ORDER BY comment DESC LIMIT 10", channelId).QueryRows(&videos)
		for _, video := range videos {
			conn.Do("zadd", key, video.Comment, video.Id)
		}
		conn.Do("expire", key, 86400*time.Second)
	}
	return num, videos, err
}

// GetTypeTop 类型排行榜
func GetTypeTop(typeId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,img,img1,add_time,episodes_count,is_end,comment FROM video WHERE status=1 AND type_id=? ORDER BY comment DESC LIMIT 10", typeId).QueryRows(&videos)
	return num, videos, err
}

// RedisGetTypeTop 增加缓存redis-影视类型排行榜
func RedisGetTypeTop(typeId int) (int64, []VideoData, error) {
	//基于评论热度进行排行，这里使用zset数据类型
	var videos []VideoData
	var num int64
	var err error
	conn := redisClient.PoolConnect()
	defer conn.Close()

	key := "video:top:type:type:" + strconv.Itoa(typeId)
	exists, err := redis.Bool(conn.Do("exists", key))
	if exists {
		num = 0
		//zrevrange第一条获取到的是id,第二条是分数
		res, _ := redis.Values(conn.Do("zrevrange", key, "0", "10", "WITHSCORES"))
		for k, v := range res {
			if k%2 == 0 {
				videoId, err := strconv.Atoi(string(v.([]byte)))
				videosInfo, err := RedisGetVideoInfo(videoId)
				if err == nil {
					video := VideoData{
						Id:            videosInfo.Id,
						Title:         videosInfo.Title,
						SubTitle:      videosInfo.SubTitle,
						AddTime:       videosInfo.AddTime,
						Img:           videosInfo.Img,
						Img1:          videosInfo.Img1,
						EpisodesCount: videosInfo.EpisodesCount,
						IsEnd:         videosInfo.IsEnd,
						Comment:       videosInfo.Comment,
					}
					videos = append(videos, video)
					num++
				}
			}
		}
	} else {
		o := orm.NewOrm()
		num, err = o.Raw("SELECT id,title,sub_title,img,img1,add_time,episodes_count,is_end,comment FROM video WHERE status=1 AND type_id=? ORDER BY comment DESC LIMIT 10", typeId).QueryRows(&videos)
		for _, video := range videos {
			conn.Do("zadd", key, video.Comment, video.Id)
		}
		conn.Do("expire", key, 86400*time.Second)
	}
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

func SaveAliyunVideo(videoId string, log string) error {
	o := orm.NewOrm()
	_, err := o.Raw("INSERT INTO aliyun_video (video_id, log, add_time) VALUES (?,?,?)", videoId, log, time.Now().Unix()).Exec()
	fmt.Println(err)
	return err
}
