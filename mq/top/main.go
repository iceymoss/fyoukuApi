package main

import (
	"encoding/json"
	"fmt"
	"fyoukuApi/models"
	"fyoukuApi/services/mq"
	redisClient "fyoukuApi/services/redis"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

// mq消费者：基于评论控制排行榜，为了提高性能和实时，当评论增加时通过MQ将视频评论总数异步+1，从而使排行榜接口进行排行
func main() {
	beego.LoadAppConfig("ini", "../../conf/app.conf")
	defaultdb := beego.AppConfig.String("defaultdb")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", defaultdb, 30, 30)

	mq.Consumer("", "fyouku_top", callback)
}

// callback 获取到mq消息，查询视频id然后评论数+1
func callback(s string) {
	type Data struct {
		VideoId int
	}
	var data Data
	err := json.Unmarshal([]byte(s), &data)
	videoInfo, err := models.RedisGetVideoInfo(data.VideoId)
	if err == nil {
		conn := redisClient.PoolConnect()
		defer conn.Close()
		//更新排行榜
		redisChannelKey := "video:top:channel:channelId:" + strconv.Itoa(videoInfo.ChannelId)
		//video:top:type:type:
		//video:top:channel:channelId
		redisTypeKey := "video:top:type:typeId:" + strconv.Itoa(videoInfo.TypeId)
		//将对应视频
		conn.Do("zincrby", redisChannelKey, 1, data.VideoId)
		conn.Do("zincrby", redisTypeKey, 1, data.VideoId)
	}
	fmt.Printf("msg is :%s\n", s)
}
