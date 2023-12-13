# fyoukuApi为仿优酷视频项目(fyouku)的后台服务

---
title: fyoukuApi v1.0.0
language_tabs:
- shell: Shell
- http: HTTP
- javascript: JavaScript
- ruby: Ruby
- python: Python
- php: PHP
- java: Java
- go: Go
  toc_footers: []
  includes: []
  search: true
  code_clipboard: true
  highlight_theme: darkula
  headingLevel: 2
  generator: "@tarslib/widdershins v4.0.17"

---

[toc]



# fyoukuApi

> v1.0.0

Base URLs:

* <a href="http://127.0.0.1:8099">开发环境: http://127.0.0.1:8099</a>

# 用户模块/登录/注册

## POST 用户注册

POST /register/save

> Body 请求参数

```yaml
mobile: "17585610985"
password: "123456"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» mobile|body|string| 是 |none|
|» password|body|string| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "注册账号成功",
  "items": null,
  "count": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» items|null|true|none||none|
|» count|integer|true|none||none|

## POST 手机号登录

POST /login/do

> Body 请求参数

```yaml
mobile: "17585610985"
password: "123456"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» mobile|body|string| 是 |none|
|» password|body|string| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "登录成功",
  "items": {
    "uid": 3,
    "username": ""
  },
  "count": 1
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» items|object|true|none||none|
|»» uid|integer|true|none||none|
|»» username|string|true|none||none|
|» count|integer|true|none||none|

# 用户模块/消息通知

## POST 通知用户

POST /send/message

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|uids|query|string| 是 |none|
|content|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 用户模块/影视上传/编辑

## GET 用户视频列表

GET /user/video

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|uid|query|number| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "items": [
    {
      "Id": 0,
      "Title": "string",
      "SubTitle": "string",
      "AddTime": 0,
      "Img": "string",
      "Img1": "string",
      "EpisodesCount": 0,
      "IsEnd": 0,
      "Comment": 0
    }
  ],
  "count": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» items|[object]|true|none||none|
|»» Id|integer|true|none||none|
|»» Title|string|true|none||none|
|»» SubTitle|string|true|none||none|
|»» AddTime|integer|true|none||none|
|»» Img|string|true|none||none|
|»» Img1|string|true|none||none|
|»» EpisodesCount|integer|true|none||none|
|»» IsEnd|integer|true|none||none|
|»» Comment|integer|true|none||none|
|» count|integer|true|none||none|

## POST 上传视频

POST /uploadDo

> Body 请求参数

```yaml
uid: 用户ID
file: file:///Users/iceymoss/Desktop/1406740804.mp4

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|uid|query|number| 是 |none|
|body|body|object| 否 |none|
|» uid|body|number| 是 |none|
|» file|body|string(binary)| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "items": null,
  "count": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» items|null|true|none||none|
|» count|integer|true|none||none|

## POST 添加视频

POST /video/save

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|playUrl|query|string| 是 |none|
|title|query|string| 是 |none|
|subTitl|query|string| 是 |none|
|channelId|query|number| 是 |none|
|typeId|query|number| 是 |none|
|regionId|query|number| 是 |none|
|uid|query|number| 是 |none|
|aliyunVideoId|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 封面图片

POST /video/img

> Body 请求参数

```yaml
uid: "2"
file: file:///Users/iceymoss/Downloads/3565777165.jpg

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» uid|body|number| 是 |none|
|» file|body|string(binary)| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "items": null,
  "count": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» items|null|true|none||none|
|» count|integer|true|none||none|

## POST 编辑影视

POST /video/update

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|videoId|query|number| 是 |none|
|uid|query|number| 是 |none|
|title|query|string| 是 |none|
|subTitle|query|string| 是 |none|
|channelId|query|number| 是 |none|
|typeId|query|number| 是 |none|
|regionId|query|number| 是 |none|
|img|query|string| 是 |none|
|img1|query|string| 是 |none|
|is_end|query|number| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "success",
  "items": null,
  "count": 1
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 编辑视频

POST /video/episodes/update

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|episodeId|query|number| 是 |none|
|uid|query|number| 是 |none|
|title|query|string| 是 |none|
|status|query|number| 是 |none|
|num|query|number| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE 删除视频

DELETE /video/episodes/delete

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|episodeId|query|number| 是 |none|
|uid|query|number| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "success",
  "items": null,
  "count": 1
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE 删除影视

DELETE /video/delete

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|videoId|query|number| 是 |none|
|uid|query|number| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 视频模块/视频筛选

## GET 频道-地区视频推荐列表

GET /channel/recommend/region

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|channelId|query|string| 是 |none|
|regionId|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 频道-类型视频推荐列表

GET /channel/recommend/type

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|channelId|query|string| 否 |none|
|typeId|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取频道地区列表

GET /channel/region

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|channelId|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取频道类型列表

GET /channel/type

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|channelId|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 筛选视频列表

GET /channel/video

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|channelId|query|string| 是 |none|
|regionId|query|string| 是 |none|
|typeId|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 视频模块/视频排行

## GET 频道-热门视频

GET /channel/hot

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|channelId|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 频道排名

GET /channel/top

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|channelId|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 类型排名

GET /type/top

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|typeId|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 视频模块/广告

## GET 频道-顶部广告

GET /channel/advert

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|channelId|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 视频模块/视频播放

## GET 视频详细信息

GET /video/info

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|videoId|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 视频播放信息

GET /video/episodes/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|videoId|query|string| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "success",
  "items": [
    {
      "Id": 1,
      "Title": "漩涡鸣人登场",
      "AddTime": 1581423724,
      "Num": 1,
      "PlayUrl": "/static/video/coverr-sparks-of-bonfire-1573980240958.mp4",
      "Comment": 10058215,
      "Status": 0,
      "AliyunVideoId": ""
    },
    {
      "Id": 2,
      "Title": "我是木叶丸",
      "AddTime": 1581423724,
      "Num": 2,
      "PlayUrl": "http://valipl.cp31.ott.cibntv.net/6975FF784E44171ADFFEC2C8F/03000500005DFA480C8BB7800000001B2BD8F5-D527-4D34-B6BF-9421CC1386AE-1-114.m3u8",
      "Comment": 0,
      "Status": 0,
      "AliyunVideoId": ""
    },
    {
      "Id": 3,
      "Title": "宿敌 佐助和小樱",
      "AddTime": 1581423724,
      "Num": 3,
      "PlayUrl": "http://valipl.cp31.ott.cibntv.net/65732AA09E93E71999F465E15/03000300005C3EE91602EA4011BA6ABCA65B2E-C656-4A40-A2DF-0D5615BBDC9A-1-114.m3u8?ccode=0502&duration=1422&expire=18000&psid=cc2bdefaf31ebf17a11d7a0d8f481af3&ups_client_netip=3b6ddb32&ups_ts=1584294982&ups_userid=1074229826&utid=iibzDpga4F4CAW%2FKOdow37WF&vid=XNTI4NjEwODg4&vkey=B3d057b782ea13d7eb639640801d05569&sm=1&operate_type=1&dre=u37&si=73&eo=0&dst=1&iv=0&s=cc001f06962411de83b1&bc=2",
      "Comment": 0,
      "Status": 0,
      "AliyunVideoId": ""
    }
  ],
  "count": 3
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» items|[object]|true|none||none|
|»» Id|integer|true|none||none|
|»» Title|string|true|none||none|
|»» AddTime|integer|true|none||none|
|»» Num|integer|true|none||none|
|»» PlayUrl|string|true|none||none|
|»» Comment|integer|true|none||none|
|»» Status|integer|true|none||none|
|»» AliyunVideoId|string|true|none||none|
|» count|integer|true|none||none|

# 视频模块/视频弹幕

## POST 发起弹幕

POST /barrage/save

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|content|query|string| 是 |弹幕内容|
|uid|query|number| 是 |弹幕发起者|
|currentTime|query|number| 是 |发起弹幕对应的播放时间|
|episodesId|query|number| 是 |视频ID|
|videoId|query|number| 是 |视频归属ID|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 视频模块/视频评论

## GET 评论列表

GET /comment/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|episodesId|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "items": [
    {
      "id": 0,
      "content": "string",
      "addTime": 0,
      "addTimeTitle": "string",
      "userId": 0,
      "stamp": 0,
      "praiseCount": 0,
      "userinfo": {
        "id": 0,
        "name": "string",
        "addTime": 0,
        "avatar": "string"
      },
      "episodesId": 0
    }
  ],
  "count": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» items|[object]|true|none||none|
|»» id|integer|true|none||none|
|»» content|string|true|none||none|
|»» addTime|integer|true|none||none|
|»» addTimeTitle|string|true|none||none|
|»» userId|integer|true|none||none|
|»» stamp|integer|true|none||none|
|»» praiseCount|integer|true|none||none|
|»» userinfo|object|true|none||none|
|»»» id|integer|true|none||none|
|»»» name|string|true|none||none|
|»»» addTime|integer|true|none||none|
|»»» avatar|string|true|none||none|
|»» episodesId|integer|true|none||none|
|» count|integer|true|none||none|

## POST 发表评论

POST /comment/save

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|content|query|string| 是 |none|
|uid|query|number| 是 |none|
|videoId|query|number| 是 |none|
|episodesId|query|number| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# ip相关

## GET 获取ip归属国

GET /geoip/country

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|ip|query|string| 是 |none|
|lan|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取ip归属市

GET /geoip/city

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|ip|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

