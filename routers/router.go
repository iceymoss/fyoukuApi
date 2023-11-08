// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"fyoukuApi/controllers"
	"github.com/astaxie/beego"
)

// 路由配置
func init() {
	//用户模块
	beego.Router("/register/save", &controllers.UserController{}, "post:SaveRegister")
	beego.Router("/login/do", &controllers.UserController{}, "post:LoginDo")

	//视频列表
	beego.Router("/channel/advert", &controllers.VideoControllers{}, "get:ChannelAdvert")
	beego.Router("/channel/hot", &controllers.VideoControllers{}, "get:GetChannelHotList")
	beego.Router("/channel/recommend/region", &controllers.VideoControllers{}, "get:GetChannelRecommendRegionList")
	beego.Router("/channel/recommend/type", &controllers.VideoControllers{}, "get:GetChannelTypeList")

	//视频筛选条件列表
	beego.Router("/channel/region", &controllers.BaseController{}, "get:ChannelRegion")
	beego.Router("/channel/type", &controllers.BaseController{}, "get:ChannelType")

	//筛选视频列表
	beego.Router("/channel/video", &controllers.VideoControllers{}, "get:ChannelVideoList")

	//视频详情
	beego.Router("/video/info", &controllers.VideoControllers{}, "get:VideoInfo")
	beego.Router("/video/episodes/list", &controllers.VideoControllers{}, "get:VideoEpisodesList")

	//beego.Include(&controllers.UserController{})
}
