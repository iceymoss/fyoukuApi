package controllers

import (
	"fmt"
	"fyoukuApi/models"
	"github.com/astaxie/beego"
)

type CommentControllers struct {
	beego.Controller
}

// CommentInfo 返回数据类型
type CommentInfo struct {
	Id           int             `json:"id"`
	Content      string          `json:"content"`
	AddTime      int64           `json:"addTime"`
	AddTimeTitle string          `json:"addTimeTitle"`
	UserId       int             `json:"userId"`
	Stamp        int             `json:"stamp"`
	PraiseCount  int             `json:"praiseCount"`
	UserInfo     models.UserInfo `json:"userinfo"`
	EpisodesId   int             `json:"episodesId"`
}

// List 视频评论列表
func (c *CommentControllers) List() {
	//获取剧集数
	episodesId, _ := c.GetInt("episodesId")
	//获取页码信息
	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")

	if episodesId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定视频剧集")
		c.ServeJSON()
	}
	if limit == 0 {
		limit = 12
	}
	num, comments, err := models.GetCommentList(episodesId, offset, limit)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
	} else {
		var data []CommentInfo
		var commentInfo CommentInfo
		for _, v := range comments {
			commentInfo.Id = v.Id
			commentInfo.Content = v.Content
			commentInfo.AddTime = v.AddTime
			commentInfo.AddTimeTitle = DateFormat(v.AddTime)
			commentInfo.UserId = v.UserId
			commentInfo.Stamp = v.Stamp
			commentInfo.PraiseCount = v.PraiseCount
			commentInfo.EpisodesId = v.EpisodesId
			//获取用户信息
			commentInfo.UserInfo, _ = models.RedisGetUserInfo(v.UserId)
			data = append(data, commentInfo)
		}
		c.Data["json"] = ReturnSuccess(0, "success", data, num)
		c.ServeJSON()
	}
}

// ListV2 视频评论列表
func (c *CommentControllers) ListV2() {
	//获取剧集数
	episodesId, _ := c.GetInt("episodesId")
	//获取页码信息
	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")

	if episodesId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定视频剧集")
		c.ServeJSON()
	}
	if limit == 0 {
		limit = 12
	}
	num, comments, err := models.GetCommentList(episodesId, offset, limit)
	if err != nil {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
	} else {
		var data []CommentInfo
		var commentInfo CommentInfo

		//并发获取用户用户信息
		uidChan := make(chan int, 12)
		closeChan := make(chan struct{}, 5)
		resChan := make(chan models.UserInfo, 12)

		//uidChan生产者
		go func() {
			for _, v := range comments {
				uidChan <- v.UserId
			}
			close(uidChan)
		}()

		for i := 0; i < 5; i++ {
			go chanGetUserInfo(uidChan, closeChan, resChan)
		}

		go func() {
			for i := 0; i < 5; i++ {
				<-closeChan
			}
			close(resChan)
			close(closeChan)
		}()

		userInfoMap := make(map[int]models.UserInfo)
		for r := range resChan {
			userInfoMap[r.Id] = r
		}
		for _, v := range comments {
			commentInfo.Id = v.Id
			commentInfo.Content = v.Content
			commentInfo.AddTime = v.AddTime
			commentInfo.AddTimeTitle = DateFormat(v.AddTime)
			commentInfo.UserId = v.UserId
			commentInfo.Stamp = v.Stamp
			commentInfo.PraiseCount = v.PraiseCount
			commentInfo.EpisodesId = v.EpisodesId
			commentInfo.UserInfo = userInfoMap[v.UserId]
			data = append(data, commentInfo)
		}
		c.Data["json"] = ReturnSuccess(0, "success", data, num)
		c.ServeJSON()
	}
}

// chanGetUserInfo 为uidChan消费者和resChan的生产者
func chanGetUserInfo(uidChan chan int, closeChan chan struct{}, resChan chan models.UserInfo) {
	for uid := range uidChan {
		userInfo, _ := models.RedisGetUserInfo(uid)
		resChan <- userInfo
	}
	closeChan <- struct{}{}
}

func (c *CommentControllers) Save() {
	//评论内容
	content := c.GetString("content")
	//用户id
	uid, _ := c.GetInt("uid")
	//视频id
	episodesId, _ := c.GetInt("episodesId")
	//所属视频id
	videoId, _ := c.GetInt("videoId")

	fmt.Printf("cont:%s, uid:%d, vid:%d, epid:%d\n", content, uid, videoId, episodesId)
	if content == "" {
		c.Data["json"] = ReturnError(4002, "评论内容不能为空")
		c.ServeJSON()
		return
	}
	if uid == 0 {
		c.Data["json"] = ReturnError(4002, "请先登录")
		c.ServeJSON()
		return
	}
	if videoId == 0 {
		c.Data["json"] = ReturnError(4003, "必须填写视频所属ID")
		c.ServeJSON()
		return
	}
	if episodesId == 0 {
		c.Data["json"] = ReturnError(4005, "必须填写视频ID")
		c.ServeJSON()
		return
	}

	//保存评论
	err := models.SaveComment(episodesId, videoId, uid, content)
	if err != nil {
		c.Data["json"] = ReturnError(5000, "保存评论失败")
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnSuccess(0, "success", nil, 1)
		c.ServeJSON()
	}
}
