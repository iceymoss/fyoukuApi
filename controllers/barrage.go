package controllers

import (
	"encoding/json"
	"fyoukuApi/models"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type BarrageControllers struct {
	beego.Controller
}

type VideoStats struct {
	views map[int]int
	mu    sync.Mutex
}

var videoStats = NewVideoStats()

func NewVideoStats() *VideoStats {
	return &VideoStats{
		views: make(map[int]int),
	}
}

func (vs *VideoStats) AddView(videoID int) {
	vs.mu.Lock()
	defer vs.mu.Unlock()
	vs.views[videoID]++
}

func (vs *VideoStats) ReduceView(videoID int) {
	vs.mu.Lock()
	defer vs.mu.Unlock()
	vs.views[videoID]--
}

func (vs *VideoStats) GetViews(videoID int) int {
	vs.mu.Lock()
	defer vs.mu.Unlock()
	return vs.views[videoID]
}

// WsData 数据格式
type WsData struct {
	CurrentTime int //当前时间
	EpisodesId  int //视频ID
}

type Barrage struct {
	BarrageMsg []models.BarrageData `json:"barrage_msg"`
	UserOnline int                  `json:"user_online"`
}

// 设置websocket跨域问题
var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// BarrageWs 获取弹幕websocket 核心逻辑：使用ws协议，从视频播放开始获取60s的弹幕内容，60s播放结束后再次请求后60s的弹幕内容，
// 前端这要在60s循环对比弹幕时间和视频播放时间对应，渲染到屏幕即可。
func (b *BarrageControllers) BarrageWs() {
	var (
		conn    *websocket.Conn
		err     error
		data    []byte
		barrage Barrage
	)

	//将http转为websocket
	if conn, err = upgrader.Upgrade(b.Ctx.ResponseWriter, b.Ctx.Request, nil); err != nil {
		conn.Close()
	}

	if _, data, err = conn.ReadMessage(); err != nil {
		conn.Close()
	}

	var wsDataInfo WsData
	json.Unmarshal(data, &wsDataInfo)

	videoStats.AddView(wsDataInfo.EpisodesId)
	ResBarrageMsg(conn, wsDataInfo, barrage)

	//用户退出视频
	defer func(Eid int) {
		videoStats.ReduceView(Eid)
		conn.Close()
	}(wsDataInfo.EpisodesId)

	//监听消息
	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			conn.Close()
		}
		var wsData WsData
		json.Unmarshal(data, &wsData)
		ResBarrageMsg(conn, wsData, barrage)
	}
}

func ResBarrageMsg(conn *websocket.Conn, wsData WsData, barrage Barrage) {
	var err error

	//当前时间开始后的60s
	endTime := wsData.CurrentTime + 60

	//获取弹幕数据
	_, barrage.BarrageMsg, err = models.BarrageList(wsData.EpisodesId, wsData.CurrentTime, endTime)

	//返回在线人数
	barrage.UserOnline = videoStats.GetViews(wsData.EpisodesId)
	if err == nil {
		if err := conn.WriteJSON(barrage); err != nil {
			conn.Close()
		}
	}
}

func (b *BarrageControllers) Save() {
	uid, _ := b.GetInt("uid")
	content := b.GetString("content")
	currentTime, _ := b.GetInt("currentTime")
	episodesId, _ := b.GetInt("episodesId")
	videoId, _ := b.GetInt("videoId")

	if content == "" {
		b.Data["json"] = ReturnError(4001, "弹幕不能为空")
		b.ServeJSON()
	}
	if uid == 0 {
		b.Data["json"] = ReturnError(4002, "请先登录")
		b.ServeJSON()
	}
	if episodesId == 0 {
		b.Data["json"] = ReturnError(4003, "必须指定剧集ID")
		b.ServeJSON()
	}
	if videoId == 0 {
		b.Data["json"] = ReturnError(4005, "必须指定视频ID")
		b.ServeJSON()
	}

	if currentTime == 0 {
		b.Data["json"] = ReturnError(4006, "必须指定视频播放时间")
		b.ServeJSON()
	}
	err := models.SaveBarrage(episodesId, videoId, currentTime, uid, content)
	if err == nil {
		b.Data["json"] = ReturnSuccess(0, "success", "", 1)
		b.ServeJSON()
	} else {
		b.Data["json"] = ReturnError(5000, err)
		b.ServeJSON()
	}
}
