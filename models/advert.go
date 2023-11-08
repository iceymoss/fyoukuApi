package models

import "github.com/astaxie/beego/orm"

type Advert struct {
	Id       int    //主键
	Title    string //标题
	SubTitle string //子标题
	AddTime  int64  //添加时间戳
	Img      string //图片链接
	Url      string //频道跳转地址
}

func init() {
	orm.RegisterModel(&Advert{})
}

func GetChannelAdvert(channelId int) (int64, []Advert, error) {
	o := orm.NewOrm()
	var adverts []Advert
	num, err := o.Raw("SELECT id, title, sub_title,img,add_time,url FROM advert WHERE status=1 AND channel_id=? ORDER BY sort DESC LIMIT 1", channelId).QueryRows(&adverts)
	return num, adverts, err
}
