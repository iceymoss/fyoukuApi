package mq

import (
	"encoding/json"
	"fmt"
	"fyoukuApi/models"
	"fyoukuApi/services/mq"
	redisClient "fyoukuApi/services/redis"
	"github.com/spf13/cobra"
	"strconv"
)

var Top = &cobra.Command{
	Use:   "rank",
	Short: "rank Starts the MQ consumer service for Ranking list",
	Run:   rankingTop,
}

// mq消费者：客服发消息只需要向mq发送，消费者进行异步处理
func rankingTop(cmd *cobra.Command, args []string) {
	mq.Consumer("", "fyouku_top", func(s string) {
		// callback 获取到mq消息，查询视频id然后评论数+1
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
	})
}

var Send = &cobra.Command{
	Use:   "send",
	Short: "send Starts the MQ consumer service for massage to user",
	Run:   sendMsgUser,
}

func sendMsgUser(cmd *cobra.Command, args []string) {
	mq.Consumer("", "fyouku_send_message_user", func(msg string) {
		// callback具体的义务逻辑：将消息写入数据库，即给用户发送消息
		type data struct {
			userId int
			msgId  int64
		}
		var userMsg data
		err := json.Unmarshal([]byte(msg), &userMsg)
		fmt.Println("userMsg:", userMsg)
		if err == nil {
			models.SendMessageUser(userMsg.userId, userMsg.msgId)
			fmt.Println("消费完成：异步写入MySQL成功")
		}
	})
}
