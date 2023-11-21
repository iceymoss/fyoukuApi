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

func main() {
	beego.LoadAppConfig("ini", "../../conf/app.conf")
	defaultdb := beego.AppConfig.String("defaultdb")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", defaultdb, 30, 30)

	mq.Consumer("", "fyouku_top", callback)
}

// callback 获取到mq消息，然后处理业务逻辑
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
