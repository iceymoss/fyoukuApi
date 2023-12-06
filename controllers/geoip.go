package controllers

import (
	"github.com/astaxie/beego"
	"github.com/oschwald/geoip2-golang"
	"net"
)

type GeoIpControllers struct {
	beego.Controller
}

func (g *GeoIpControllers) GetCountryInfo() {
	ip := g.GetString("ip")
	lan := g.GetString("lan")
	if lan != "en" && lan != "cn" && lan != "" {
		g.Data["json"] = ReturnError(4001, "lan参数非法")
		g.ServeJSON()
		return
	}
	realIp := net.ParseIP(ip)
	if realIp == nil || realIp.String() != ip {
		g.Data["json"] = ReturnError(4001, "ip参数非法")
		g.ServeJSON()
		return
	}
	// 打开GeoLite2-Country数据库文件
	db, err := geoip2.Open("./resource/libs/GeoLite2-Country.mmdb")
	if err != nil {
		g.Data["json"] = ReturnError(5000, err.Error())
		g.ServeJSON()
		return
	}
	defer db.Close()

	// 查询IP所属国家
	record, err := db.Country(realIp)
	if err != nil {
		g.Data["json"] = ReturnError(5000, err.Error())
		g.ServeJSON()
		return
	}
	data := make(map[string]interface{})
	data["country"] = record.Country.Names["zh-CN"]
	data["code"] = record.Country.IsoCode
	data["Continent"] = record.Continent.Names["zh-CN"]
	g.Data["json"] = ReturnSuccess(0, "success", data, 0)
	g.ServeJSON()
}

func (g *GeoIpControllers) GetCityInfo() {
	ip := g.GetString("ip")
	realIp := net.ParseIP(ip)
	if realIp == nil || realIp.String() != ip {
		g.Data["json"] = ReturnError(4001, "ip参数非法")
		g.ServeJSON()
		return
	}
	// 打开GeoLite2-Country数据库文件
	db, err := geoip2.Open("./resource/libs/GeoLite2-City.mmdb")
	if err != nil {
		g.Data["json"] = ReturnError(5000, err.Error())
		g.ServeJSON()
		return
	}
	defer db.Close()

	// 查询IP所属国家
	record, err := db.City(realIp)
	if err != nil {
		g.Data["json"] = ReturnError(5000, err.Error())
		g.ServeJSON()
		return
	}
	data := make(map[string]interface{})
	data["country"] = record.Country.Names["zh-CN"]
	data["code"] = record.Country.IsoCode
	data["Continent"] = record.Continent.Names["zh-CN"]
	data["city"] = record.City.Names["zh-CN"]
	data["Location"] = record.Location
	g.Data["json"] = ReturnSuccess(0, "success", data, 0)
	g.ServeJSON()
}
