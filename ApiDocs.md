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
mobile: "17553217643"
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

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» items|null|true|none||none|
|» count|integer|true|none||none|

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

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» items|null|true|none||none|
|» count|integer|true|none||none|

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

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» items|null|true|none||none|
|» count|integer|true|none||none|

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

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» items|null|true|none||none|
|» count|integer|true|none||none|

## DELETE 删除影视

DELETE /video/delete

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|videoId|query|number| 是 |none|
|uid|query|number| 否 |none|

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

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» items|null|true|none||none|
|» count|integer|true|none||none|

# 视频模块/视频筛选

## GET 频道-地区视频推荐列表

GET /channel/recommend/region

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|channelId|query|string| 是 |none|
|regionId|query|string| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "success",
  "items": [
    {
      "Id": 42,
      "Title": "哪吒-我命由我不由天",
      "SubTitle": "哪吒-我命由我不由天",
      "AddTime": 1587653999,
      "Img": "/static/data/new/5.png",
      "Img1": "/static/data/new/5.png",
      "EpisodesCount": 1,
      "IsEnd": 0,
      "Comment": 0
    },
    {
      "Id": 2,
      "Title": "阿衰 第二季",
      "SubTitle": "怕踢学校搞笑故事再次归来",
      "AddTime": 1581423724,
      "Img": "/static/data/new/2.png",
      "Img1": "/static/data/new/2a.png",
      "EpisodesCount": 15,
      "IsEnd": 0,
      "Comment": 0
    },
    {
      "Id": 28,
      "Title": "罗小黑战记",
      "SubTitle": "同名大电影热映中",
      "AddTime": 1581423724,
      "Img": "/static/data/new/28.png",
      "Img1": "/static/data/new/28a.png",
      "EpisodesCount": 28,
      "IsEnd": 1,
      "Comment": 0
    },
    {
      "Id": 29,
      "Title": "秦时明月之君临天下",
      "SubTitle": "沧海横流 江湖再见",
      "AddTime": 1581423724,
      "Img": "/static/data/new/29.png",
      "Img1": "/static/data/new/29a.png",
      "EpisodesCount": 75,
      "IsEnd": 1,
      "Comment": 0
    },
    {
      "Id": 30,
      "Title": "镇魂街",
      "SubTitle": "吸纳亡灵镇压恶灵之地",
      "AddTime": 1581423724,
      "Img": "/static/data/new/30.png",
      "Img1": "/static/data/new/30a.png",
      "EpisodesCount": 24,
      "IsEnd": 1,
      "Comment": 0
    },
    {
      "Id": 31,
      "Title": "秦时明月之诸子百家",
      "SubTitle": "出机关城伪装入桑海",
      "AddTime": 1581423724,
      "Img": "/static/data/new/31.png",
      "Img1": "/static/data/new/31a.png",
      "EpisodesCount": 34,
      "IsEnd": 1,
      "Comment": 0
    },
    {
      "Id": 32,
      "Title": "少年锦衣卫",
      "SubTitle": "三盗大闹皇宫劫公主",
      "AddTime": 1581423724,
      "Img": "/static/data/new/32.png",
      "Img1": "/static/data/new/32a.png",
      "EpisodesCount": 28,
      "IsEnd": 1,
      "Comment": 0
    },
    {
      "Id": 33,
      "Title": "十万个冷笑话 第三季",
      "SubTitle": "经典角色悉数回归爆笑玩梗",
      "AddTime": 1581423724,
      "Img": "/static/data/new/33.png",
      "Img1": "/static/data/new/33a.png",
      "EpisodesCount": 52,
      "IsEnd": 1,
      "Comment": 0
    },
    {
      "Id": 34,
      "Title": "超神学院",
      "SubTitle": "英雄联盟改编人气动画",
      "AddTime": 1581423724,
      "Img": "/static/data/new/34.png",
      "Img1": "/static/data/new/34a.png",
      "EpisodesCount": 30,
      "IsEnd": 1,
      "Comment": 0
    }
  ],
  "count": 9
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

## GET 频道-类型视频推荐列表

GET /channel/recommend/type

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|channelId|query|string| 否 |none|
|typeId|query|string| 否 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "success",
  "items": [
    {
      "Id": 42,
      "Title": "哪吒-我命由我不由天",
      "SubTitle": "哪吒-我命由我不由天",
      "AddTime": 1587653999,
      "Img": "/static/data/new/5.png",
      "Img1": "/static/data/new/5.png",
      "EpisodesCount": 1,
      "IsEnd": 0,
      "Comment": 0
    },
    {
      "Id": 2,
      "Title": "阿衰 第二季",
      "SubTitle": "怕踢学校搞笑故事再次归来",
      "AddTime": 1581423724,
      "Img": "/static/data/new/2.png",
      "Img1": "/static/data/new/2a.png",
      "EpisodesCount": 15,
      "IsEnd": 0,
      "Comment": 0
    },
    {
      "Id": 11,
      "Title": "猎人 重制版",
      "SubTitle": "主人公寻父的冒险路",
      "AddTime": 1581423724,
      "Img": "/static/data/new/11.png",
      "Img1": "/static/data/new/11a.png",
      "EpisodesCount": 148,
      "IsEnd": 1,
      "Comment": 0
    },
    {
      "Id": 12,
      "Title": "银魂",
      "SubTitle": "跟银魂比吐槽你输定了",
      "AddTime": 1581423724,
      "Img": "/static/data/new/18.png",
      "Img1": "/static/data/new/18a.png",
      "EpisodesCount": 368,
      "IsEnd": 1,
      "Comment": 0
    },
    {
      "Id": 17,
      "Title": "BLEACH 境&middot;界",
      "SubTitle": "古惑仔之人在地府",
      "AddTime": 1581423724,
      "Img": "/static/data/new/17.png",
      "Img1": "/static/data/new/17a.png",
      "EpisodesCount": 366,
      "IsEnd": 1,
      "Comment": 0
    },
    {
      "Id": 28,
      "Title": "罗小黑战记",
      "SubTitle": "同名大电影热映中",
      "AddTime": 1581423724,
      "Img": "/static/data/new/28.png",
      "Img1": "/static/data/new/28a.png",
      "EpisodesCount": 28,
      "IsEnd": 1,
      "Comment": 0
    },
    {
      "Id": 29,
      "Title": "秦时明月之君临天下",
      "SubTitle": "沧海横流 江湖再见",
      "AddTime": 1581423724,
      "Img": "/static/data/new/29.png",
      "Img1": "/static/data/new/29a.png",
      "EpisodesCount": 75,
      "IsEnd": 1,
      "Comment": 0
    },
    {
      "Id": 30,
      "Title": "镇魂街",
      "SubTitle": "吸纳亡灵镇压恶灵之地",
      "AddTime": 1581423724,
      "Img": "/static/data/new/30.png",
      "Img1": "/static/data/new/30a.png",
      "EpisodesCount": 24,
      "IsEnd": 1,
      "Comment": 0
    },
    {
      "Id": 31,
      "Title": "秦时明月之诸子百家",
      "SubTitle": "出机关城伪装入桑海",
      "AddTime": 1581423724,
      "Img": "/static/data/new/31.png",
      "Img1": "/static/data/new/31a.png",
      "EpisodesCount": 34,
      "IsEnd": 1,
      "Comment": 0
    }
  ],
  "count": 9
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

## GET 获取频道地区列表

GET /channel/region

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|channelId|query|string| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "success",
  "items": [
    {
      "Id": 2,
      "Name": "中国大陆"
    },
    {
      "Id": 1,
      "Name": "日本"
    },
    {
      "Id": 3,
      "Name": "美国"
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
|»» Name|string|true|none||none|
|» count|integer|true|none||none|

## GET 获取频道类型列表

GET /channel/type

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|channelId|query|string| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "success",
  "items": [
    {
      "Id": 1,
      "Name": "少女"
    },
    {
      "Id": 2,
      "Name": "热血"
    },
    {
      "Id": 3,
      "Name": "青春"
    },
    {
      "Id": 4,
      "Name": "社会"
    },
    {
      "Id": 5,
      "Name": "科幻"
    },
    {
      "Id": 6,
      "Name": "搞笑"
    }
  ],
  "count": 6
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
|»» Name|string|true|none||none|
|» count|integer|true|none||none|

## GET 筛选视频列表

GET /channel/video

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|channelId|query|string| 是 |none|
|regionId|query|string| 是 |none|
|typeId|query|string| 否 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "success",
  "items": [
    {
      "AddTime": 1699950217,
      "EpisodesCount": 1,
      "Id": 38,
      "Img": "",
      "Img1": "",
      "IsEnd": 1,
      "SubTitle": "",
      "Title": "小姐姐"
    },
    {
      "AddTime": 1581423724,
      "EpisodesCount": 873,
      "Id": 10,
      "Img": "/static/data/new/10.png",
      "Img1": "/static/data/new/10a.png",
      "IsEnd": 1,
      "SubTitle": "人小鬼大的野原新之助",
      "Title": "蜡笔小新 第二季"
    },
    {
      "AddTime": 1581423724,
      "EpisodesCount": 156,
      "Id": 13,
      "Img": "/static/data/new/13.png",
      "Img1": "/static/data/new/13a.png",
      "IsEnd": 1,
      "SubTitle": "野原新之助的日常生活",
      "Title": "蜡笔小新 第三季"
    },
    {
      "AddTime": 1581423724,
      "EpisodesCount": 178,
      "Id": 14,
      "Img": "/static/data/new/14.png",
      "Img1": "/static/data/new/14a.png",
      "IsEnd": 1,
      "SubTitle": "运动美少年热血故事",
      "Title": "网球王子"
    },
    {
      "AddTime": 1581423724,
      "EpisodesCount": 51,
      "Id": 15,
      "Img": "/static/data/new/15.png",
      "Img1": "/static/data/new/15a.png",
      "IsEnd": 1,
      "SubTitle": "能实现愿望的魔法蛋",
      "Title": "守护甜心 第一季"
    },
    {
      "AddTime": 1581423724,
      "EpisodesCount": 144,
      "Id": 16,
      "Img": "/static/data/new/16.png",
      "Img1": "/static/data/new/16a.png",
      "IsEnd": 1,
      "SubTitle": "星球上不可思议的生物",
      "Title": "精灵宝可梦 第四季"
    },
    {
      "AddTime": 1581423724,
      "EpisodesCount": 167,
      "Id": 18,
      "Img": "/static/data/new/38.png",
      "Img1": "/static/data/new/38a.png",
      "IsEnd": 1,
      "SubTitle": "寻找四魂碎片的旅程",
      "Title": "犬夜叉"
    },
    {
      "AddTime": 1581423724,
      "EpisodesCount": 920,
      "Id": 21,
      "Img": "/static/data/new/21.png",
      "Img1": "/static/data/new/21a.png",
      "IsEnd": 0,
      "SubTitle": "路飞再踏热血征程",
      "Title": "海贼王"
    },
    {
      "AddTime": 1581423724,
      "EpisodesCount": 13,
      "Id": 22,
      "Img": "/static/data/new/22.png",
      "Img1": "/static/data/new/22a.png",
      "IsEnd": 1,
      "SubTitle": "软萌少女结缘狐神",
      "Title": "元气少女缘结神 第一季"
    },
    {
      "AddTime": 1581423724,
      "EpisodesCount": 15,
      "Id": 23,
      "Img": "/static/data/new/23.png",
      "Img1": "/static/data/new/23a.png",
      "IsEnd": 1,
      "SubTitle": "高中生成后宫之王",
      "Title": "爱神巧克力"
    },
    {
      "AddTime": 1581423724,
      "EpisodesCount": 24,
      "Id": 24,
      "Img": "/static/data/new/24.png",
      "Img1": "/static/data/new/24a.png",
      "IsEnd": 1,
      "SubTitle": "骨叔新娘养成计划",
      "Title": "魔法使的新娘"
    },
    {
      "AddTime": 1581423724,
      "EpisodesCount": 1,
      "Id": 25,
      "Img": "/static/data/new/25.png",
      "Img1": "/static/data/new/25a.png",
      "IsEnd": 1,
      "SubTitle": "宫崎骏版小美人鱼",
      "Title": "悬崖上的金鱼公主"
    }
  ],
  "count": 14
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
|»» AddTime|integer|true|none||none|
|»» EpisodesCount|integer|true|none||none|
|»» Id|integer|true|none||none|
|»» Img|string|true|none||none|
|»» Img1|string|true|none||none|
|»» IsEnd|integer|true|none||none|
|»» SubTitle|string|true|none||none|
|»» Title|string|true|none||none|
|» count|integer|true|none||none|

# 视频模块/视频排行

## GET 频道-热门视频

GET /channel/hot

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|channelId|query|string| 否 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "success",
  "items": [
    {
      "Id": 42,
      "Title": "哪吒-我命由我不由天",
      "SubTitle": "哪吒-我命由我不由天",
      "AddTime": 1587653999,
      "Img": "/static/data/new/5.png",
      "Img1": "/static/data/new/5.png",
      "EpisodesCount": 1,
      "IsEnd": 0,
      "Comment": 0
    },
    {
      "Id": 1,
      "Title": "火影忍者",
      "SubTitle": "孤独少年忍者世界英雄梦",
      "AddTime": 1581423724,
      "Img": "/static/data/new/14.png",
      "Img1": "/static/data/new/14a.png",
      "EpisodesCount": 720,
      "IsEnd": 1,
      "Comment": 0
    },
    {
      "Id": 2,
      "Title": "阿衰 第二季",
      "SubTitle": "怕踢学校搞笑故事再次归来",
      "AddTime": 1581423724,
      "Img": "/static/data/new/2.png",
      "Img1": "/static/data/new/2a.png",
      "EpisodesCount": 15,
      "IsEnd": 0,
      "Comment": 0
    },
    {
      "Id": 3,
      "Title": "名侦探柯南",
      "SubTitle": "万年小学生的推理生涯",
      "AddTime": 1581423724,
      "Img": "/static/data/new/3.png",
      "Img1": "/static/data/new/3a.png",
      "EpisodesCount": 1021,
      "IsEnd": 0,
      "Comment": 0
    },
    {
      "Id": 4,
      "Title": "博人传 火影忍者新时代",
      "SubTitle": "鸣人之子续写忍者传奇",
      "AddTime": 1581423724,
      "Img": "/static/data/new/4.png",
      "Img1": "/static/data/new/4a.png",
      "EpisodesCount": 137,
      "IsEnd": 0,
      "Comment": 0
    },
    {
      "Id": 5,
      "Title": "一拳超人 第二季",
      "SubTitle": "一拳埼玉热血回归",
      "AddTime": 1581423724,
      "Img": "/static/data/new/5.png",
      "Img1": "/static/data/new/5a.png",
      "EpisodesCount": 13,
      "IsEnd": 0,
      "Comment": 0
    },
    {
      "Id": 6,
      "Title": "排球少年 第四季",
      "SubTitle": "排球部迎接更艰巨的挑战",
      "AddTime": 1581423724,
      "Img": "/static/data/new/6.png",
      "Img1": "/static/data/new/6a.png",
      "EpisodesCount": 5,
      "IsEnd": 0,
      "Comment": 0
    },
    {
      "Id": 7,
      "Title": "秦时明月之沧海横流",
      "SubTitle": "十年国漫十年秦时",
      "AddTime": 1581423724,
      "Img": "/static/data/new/7.png",
      "Img1": "/static/data/new/7a.png",
      "EpisodesCount": 1,
      "IsEnd": 0,
      "Comment": 0
    },
    {
      "Id": 8,
      "Title": "猫和老鼠",
      "SubTitle": "汤姆和杰瑞的搞笑日常",
      "AddTime": 1581423724,
      "Img": "/static/data/new/39.png",
      "Img1": "/static/data/new/39a.png",
      "EpisodesCount": 140,
      "IsEnd": 1,
      "Comment": 0
    }
  ],
  "count": 9
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

## GET 频道排名

GET /channel/top

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|channelId|query|string| 否 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "success",
  "items": [
    {
      "Id": 1,
      "Title": "火影忍者",
      "SubTitle": "孤独少年忍者世界英雄梦",
      "AddTime": 1581423724,
      "Img": "/static/data/new/14.png",
      "Img1": "/static/data/new/14a.png",
      "EpisodesCount": 720,
      "IsEnd": 1,
      "Comment": 10077219
    },
    {
      "Id": 5,
      "Title": "一拳超人 第二季",
      "SubTitle": "一拳埼玉热血回归",
      "AddTime": 1581423724,
      "Img": "/static/data/new/5.png",
      "Img1": "/static/data/new/5a.png",
      "EpisodesCount": 13,
      "IsEnd": 0,
      "Comment": 98761
    },
    {
      "Id": 3,
      "Title": "名侦探柯南",
      "SubTitle": "万年小学生的推理生涯",
      "AddTime": 1581423724,
      "Img": "/static/data/new/3.png",
      "Img1": "/static/data/new/3a.png",
      "EpisodesCount": 1021,
      "IsEnd": 0,
      "Comment": 29876
    },
    {
      "Id": 6,
      "Title": "排球少年 第四季",
      "SubTitle": "排球部迎接更艰巨的挑战",
      "AddTime": 1581423724,
      "Img": "/static/data/new/6.png",
      "Img1": "/static/data/new/6a.png",
      "EpisodesCount": 5,
      "IsEnd": 0,
      "Comment": 23564
    },
    {
      "Id": 4,
      "Title": "博人传 火影忍者新时代",
      "SubTitle": "鸣人之子续写忍者传奇",
      "AddTime": 1581423724,
      "Img": "/static/data/new/4.png",
      "Img1": "/static/data/new/4a.png",
      "EpisodesCount": 137,
      "IsEnd": 0,
      "Comment": 18976
    },
    {
      "Id": 2,
      "Title": "阿衰 第二季",
      "SubTitle": "怕踢学校搞笑故事再次归来",
      "AddTime": 1581423724,
      "Img": "/static/data/new/2.png",
      "Img1": "/static/data/new/2a.png",
      "EpisodesCount": 15,
      "IsEnd": 0,
      "Comment": 17860
    },
    {
      "Id": 22,
      "Title": "元气少女缘结神 第一季",
      "SubTitle": "软萌少女结缘狐神",
      "AddTime": 1581423724,
      "Img": "/static/data/new/22.png",
      "Img1": "/static/data/new/22a.png",
      "EpisodesCount": 13,
      "IsEnd": 1,
      "Comment": 9988
    },
    {
      "Id": 23,
      "Title": "爱神巧克力",
      "SubTitle": "高中生成后宫之王",
      "AddTime": 1581423724,
      "Img": "/static/data/new/23.png",
      "Img1": "/static/data/new/23a.png",
      "EpisodesCount": 15,
      "IsEnd": 1,
      "Comment": 9678
    },
    {
      "Id": 26,
      "Title": "飞天小女警Z",
      "SubTitle": "超能力少女卫东京",
      "AddTime": 1581423724,
      "Img": "/static/data/new/26.png",
      "Img1": "/static/data/new/26a.png",
      "EpisodesCount": 52,
      "IsEnd": 1,
      "Comment": 8735
    },
    {
      "Id": 24,
      "Title": "魔法使的新娘",
      "SubTitle": "骨叔新娘养成计划",
      "AddTime": 1581423724,
      "Img": "/static/data/new/24.png",
      "Img1": "/static/data/new/24a.png",
      "EpisodesCount": 24,
      "IsEnd": 1,
      "Comment": 8080
    },
    {
      "Id": 37,
      "Title": "小姐姐真漂亮",
      "SubTitle": "都属于的小姐姐都好漂亮啊",
      "AddTime": 1699949304,
      "Img": "",
      "Img1": "",
      "EpisodesCount": 1,
      "IsEnd": 1,
      "Comment": 1
    }
  ],
  "count": 11
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

## GET 类型排名

GET /type/top

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|typeId|query|string| 否 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "success",
  "items": [
    {
      "Id": 22,
      "Title": "元气少女缘结神 第一季",
      "SubTitle": "软萌少女结缘狐神",
      "AddTime": 1581423724,
      "Img": "/static/data/new/22.png",
      "Img1": "/static/data/new/22a.png",
      "EpisodesCount": 13,
      "IsEnd": 1,
      "Comment": 9988
    },
    {
      "Id": 23,
      "Title": "爱神巧克力",
      "SubTitle": "高中生成后宫之王",
      "AddTime": 1581423724,
      "Img": "/static/data/new/23.png",
      "Img1": "/static/data/new/23a.png",
      "EpisodesCount": 15,
      "IsEnd": 1,
      "Comment": 9678
    },
    {
      "Id": 26,
      "Title": "飞天小女警Z",
      "SubTitle": "超能力少女卫东京",
      "AddTime": 1581423724,
      "Img": "/static/data/new/26.png",
      "Img1": "/static/data/new/26a.png",
      "EpisodesCount": 52,
      "IsEnd": 1,
      "Comment": 8735
    },
    {
      "Id": 24,
      "Title": "魔法使的新娘",
      "SubTitle": "骨叔新娘养成计划",
      "AddTime": 1581423724,
      "Img": "/static/data/new/24.png",
      "Img1": "/static/data/new/24a.png",
      "EpisodesCount": 24,
      "IsEnd": 1,
      "Comment": 8080
    },
    {
      "Id": 25,
      "Title": "悬崖上的金鱼公主",
      "SubTitle": "宫崎骏版小美人鱼",
      "AddTime": 1581423724,
      "Img": "/static/data/new/25.png",
      "Img1": "/static/data/new/25a.png",
      "EpisodesCount": 1,
      "IsEnd": 1,
      "Comment": 5674
    },
    {
      "Id": 27,
      "Title": "一起来看流星雨",
      "SubTitle": "灰姑娘的校园爱情",
      "AddTime": 1581423724,
      "Img": "/static/data/new/27.png",
      "Img1": "/static/data/new/27a.png",
      "EpisodesCount": 52,
      "IsEnd": 1,
      "Comment": 3476
    },
    {
      "Id": 37,
      "Title": "小姐姐真漂亮",
      "SubTitle": "都属于的小姐姐都好漂亮啊",
      "AddTime": 1699949304,
      "Img": "",
      "Img1": "",
      "EpisodesCount": 1,
      "IsEnd": 1,
      "Comment": 1
    },
    {
      "Id": 41,
      "Title": "抖音博主",
      "SubTitle": "这是来自抖音的博主",
      "AddTime": 1699963622,
      "Img": "",
      "Img1": "",
      "EpisodesCount": 1,
      "IsEnd": 1,
      "Comment": 0
    },
    {
      "Id": 15,
      "Title": "守护甜心 第一季",
      "SubTitle": "能实现愿望的魔法蛋",
      "AddTime": 1581423724,
      "Img": "/static/data/new/15.png",
      "Img1": "/static/data/new/15a.png",
      "EpisodesCount": 51,
      "IsEnd": 1,
      "Comment": 0
    },
    {
      "Id": 13,
      "Title": "蜡笔小新 第三季",
      "SubTitle": "野原新之助的日常生活",
      "AddTime": 1581423724,
      "Img": "/static/data/new/13.png",
      "Img1": "/static/data/new/13a.png",
      "EpisodesCount": 156,
      "IsEnd": 1,
      "Comment": 0
    },
    {
      "Id": 10,
      "Title": "蜡笔小新 第二季",
      "SubTitle": "人小鬼大的野原新之助",
      "AddTime": 1581423724,
      "Img": "/static/data/new/10.png",
      "Img1": "/static/data/new/10a.png",
      "EpisodesCount": 873,
      "IsEnd": 1,
      "Comment": 0
    }
  ],
  "count": 11
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

# 视频模块/视频播放

## GET 视频详细信息

GET /video/info

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|videoId|query|string| 否 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "success",
  "items": {
    "Id": 43,
    "Title": "斗破苍穹",
    "SubTitle": "双帝之战",
    "AddTime": 1699966499,
    "Img": "/static/data/img/ddfhidfew8dfn43293dnf.png",
    "Img1": "",
    "EpisodesCount": 1,
    "IsEnd": 1,
    "ChannelId": 1,
    "Status": 1,
    "RegionId": 2,
    "TypeId": 2,
    "EpisodesUpdateTime": 1699966499,
    "Comment": 1,
    "UserId": 2,
    "IsRecommend": 0
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
|»» Id|integer|true|none||none|
|»» Title|string|true|none||none|
|»» SubTitle|string|true|none||none|
|»» AddTime|integer|true|none||none|
|»» Img|string|true|none||none|
|»» Img1|string|true|none||none|
|»» EpisodesCount|integer|true|none||none|
|»» IsEnd|integer|true|none||none|
|»» ChannelId|integer|true|none||none|
|»» Status|integer|true|none||none|
|»» RegionId|integer|true|none||none|
|»» TypeId|integer|true|none||none|
|»» EpisodesUpdateTime|integer|true|none||none|
|»» Comment|integer|true|none||none|
|»» UserId|integer|true|none||none|
|»» IsRecommend|integer|true|none||none|
|» count|integer|true|none||none|

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

# 视频模块/视频搜索

## POST 视频搜索

POST /video/search

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|key_word|query|string| 是 |none|
|limit|query|integer| 否 |none|
|offset|query|integer| 否 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "",
  "items": [
    {
      "Id": 40,
      "Title": "抖音小姐姐",
      "SubTitle": "",
      "AddTime": 0,
      "Img": "",
      "Img1": "",
      "EpisodesCount": 0,
      "IsEnd": 0,
      "ChannelId": 0,
      "Status": 1,
      "RegionId": 0,
      "TypeId": 0,
      "EpisodesUpdateTime": 0,
      "Comment": 0,
      "UserId": 0,
      "IsRecommend": 0
    },
    {
      "Id": 38,
      "Title": "小姐姐",
      "SubTitle": "",
      "AddTime": 0,
      "Img": "",
      "Img1": "",
      "EpisodesCount": 0,
      "IsEnd": 0,
      "ChannelId": 0,
      "Status": 1,
      "RegionId": 0,
      "TypeId": 0,
      "EpisodesUpdateTime": 0,
      "Comment": 0,
      "UserId": 0,
      "IsRecommend": 0
    },
    {
      "Id": 37,
      "Title": "小姐姐真漂亮",
      "SubTitle": "",
      "AddTime": 0,
      "Img": "",
      "Img1": "",
      "EpisodesCount": 0,
      "IsEnd": 0,
      "ChannelId": 0,
      "Status": 1,
      "RegionId": 0,
      "TypeId": 0,
      "EpisodesUpdateTime": 0,
      "Comment": 0,
      "UserId": 0,
      "IsRecommend": 0
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
|»» SubTitle|string|true|none||none|
|»» AddTime|integer|true|none||none|
|»» Img|string|true|none||none|
|»» Img1|string|true|none||none|
|»» EpisodesCount|integer|true|none||none|
|»» IsEnd|integer|true|none||none|
|»» ChannelId|integer|true|none||none|
|»» Status|integer|true|none||none|
|»» RegionId|integer|true|none||none|
|»» TypeId|integer|true|none||none|
|»» EpisodesUpdateTime|integer|true|none||none|
|»» Comment|integer|true|none||none|
|»» UserId|integer|true|none||none|
|»» IsRecommend|integer|true|none||none|
|» count|integer|true|none||none|

# ip模块

## GET 获取ip归属国

GET /geoip/country

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|ip|query|string| 是 |none|
|lan|query|string| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "success",
  "items": {
    "Continent": "亚洲",
    "code": "CN",
    "country": "中国"
  },
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
|» items|object|true|none||none|
|»» Continent|string|true|none||none|
|»» code|string|true|none||none|
|»» country|string|true|none||none|
|» count|integer|true|none||none|

## GET 获取ip归属市

GET /geoip/city

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|ip|query|string| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "success",
  "items": {
    "Continent": "北美洲",
    "Location": {
      "TimeZone": "America/Indiana/Indianapolis",
      "Latitude": 39.968,
      "Longitude": -86.0853,
      "MetroCode": 527,
      "AccuracyRadius": 20
    },
    "city": "卡梅尔",
    "code": "US",
    "country": "美国"
  },
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
|» items|object|true|none||none|
|»» Continent|string|true|none||none|
|»» Location|object|true|none||none|
|»»» TimeZone|string|true|none||none|
|»»» Latitude|number|true|none||none|
|»»» Longitude|number|true|none||none|
|»»» MetroCode|integer|true|none||none|
|»»» AccuracyRadius|integer|true|none||none|
|»» city|string|true|none||none|
|»» code|string|true|none||none|
|»» country|string|true|none||none|
|» count|integer|true|none||none|

# 广告模块/广告

## GET 频道-顶部广告

GET /channel/advert

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|channelId|query|string| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "success",
  "items": [
    {
      "Id": 1,
      "Title": "黑色四叶草",
      "SubTitle": "黑色暴牛骑士团",
      "AddTime": 1581600338,
      "Img": "/static/data/new/1b.png",
      "Url": "/show?id=9"
    }
  ],
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
|» items|[object]|true|none||none|
|»» Id|integer|false|none||none|
|»» Title|string|false|none||none|
|»» SubTitle|string|false|none||none|
|»» AddTime|integer|false|none||none|
|»» Img|string|false|none||none|
|»» Url|string|false|none||none|
|» count|integer|true|none||none|

# 弹幕模块/视频弹幕

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
{
  "code": 0,
  "msg": "string",
  "items": "string",
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
|» items|string|true|none||none|
|» count|integer|true|none||none|



## WX 获取弹幕

ws://127.0.0.1:8099/barrage/ws

### 请求参数

```json
{
    "currentTime":0,
    "episodesId":18
}
```

### 说明

发送消息：
> {
> "currentTime": 1,   //当前播放时间
> "episodesId": 1     //当前播放视频ID
> }

时间起始到60秒的弹幕返回：
> [{"id":47,"content":"蔡徐坤","currentTime":6},{"id":38,"content":"满满的回忆","currentTime":10},{"id":32,"content":"厉害","currentTime":15},{"id":37,"content":"满满的回忆呀","currentTime":15},{"id":39,"content":"好厉害呀","currentTime":25},{"id":33,"content":"佐助貌似没有准备时间","currentTime":28},{"id":40,"content":"佐助","currentTime":36},{"id":46,"content":"buc","currentTime":42},{"id":41,"content":"好可爱哦","currentTime":49},{"id":34,"content":"鸣人加油","currentTime":50}]



# 评论模块/视频评论

## GET 评论列表

GET /comment/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|episodesId|query|string| 是 |none|

> 返回示例

> 成功

```json
"{\n  \"code\": 0,\n  \"msg\": \"success\",\n  \"items\": [\n    {\n      \"id\": 52,\n      \"content\": \"吃了吗？\",\n      \"addTime\": 1700561028,\n      \"addTimeTitle\": \"2023-11-21\",\n      \"userId\": 2,\n      \"stamp\": 0,\n      \"praiseCount\": 0,\n      \"userinfo\": {\n        \"id\": 2,\n        \"name\": \"蔡徐坤的鸽鸽\",\n        \"addTime\": 1699374857,\n        \"avatar\": \"\"\n      },\n      \"episodesId\": 1\n    },\n    {\n      \"id\": 50,\n      \"content\": \"嘿嘿嘿\",\n      \"addTime\": 1699886746,\n      \"addTimeTitle\": \"2023-11-13\",\n      \"userId\": 2,\n      \"stamp\": 0,\n      \"praiseCount\": 0,\n      \"userinfo\": {\n        \"id\": 2,\n        \"name\": \"蔡徐坤的鸽鸽\",\n        \"addTime\": 1699374857,\n        \"avatar\": \"\"\n      },\n      \"episodesId\": 1\n    },\n    {\n      \"id\": 49,\n      \"content\": \"鸽鸽好\",\n      \"addTime\": 1699885608,\n      \"addTimeTitle\": \"2023-11-13\",\n      \"userId\": 2,\n      \"stamp\": 0,\n      \"praiseCount\": 0,\n      \"userinfo\": {\n        \"id\": 2,\n        \"name\": \"蔡徐坤的鸽鸽\",\n        \"addTime\": 1699374857,\n        \"avatar\": \"\"\n      },\n      \"episodesId\": 1\n    },\n    {\n      \"id\": 48,\n      \"content\": \"嘿嘿嘿\",\n      \"addTime\": 1699885507,\n      \"addTimeTitle\": \"2023-11-13\",\n      \"userId\": 2,\n      \"stamp\": 0,\n      \"praiseCount\": 0,\n      \"userinfo\": {\n        \"id\": 2,\n        \"name\": \"蔡徐坤的鸽鸽\",\n        \"addTime\": 1699374857,\n        \"avatar\": \"\"\n      },\n      \"episodesId\": 1\n    },\n    {\n      \"id\": 42,\n      \"content\": \"鸡你太美！\",\n      \"addTime\": 1699884311,\n      \"addTimeTitle\": \"2023-11-13\",\n      \"userId\": 2,\n      \"stamp\": 0,\n      \"praiseCount\": 0,\n      \"userinfo\": {\n        \"id\": 2,\n        \"name\": \"蔡徐坤的鸽鸽\",\n        \"addTime\": 1699374857,\n        \"avatar\": \"\"\n      },\n      \"episodesId\": 1\n    },\n    {\n      \"id\": 41,\n      \"content\": \"再来\",\n      \"addTime\": 1587224736,\n      \"addTimeTitle\": \"2020-04-18\",\n      \"userId\": 10,\n      \"stamp\": 0,\n      \"praiseCount\": 0,\n      \"userinfo\": {\n        \"id\": 0,\n        \"name\": \"\",\n        \"addTime\": 0,\n        \"avatar\": \"\"\n      },\n      \"episodesId\": 1\n    },\n    {\n      \"id\": 40,\n      \"content\": \"评论一下\",\n      \"addTime\": 1587224714,\n      \"addTimeTitle\": \"2020-04-18\",\n      \"userId\": 10,\n      \"stamp\": 0,\n      \"praiseCount\": 0,\n      \"userinfo\": {\n        \"id\": 0,\n        \"name\": \"\",\n        \"addTime\": 0,\n        \"avatar\": \"\"\n      },\n      \"episodesId\": 1\n    },\n    {\n      \"id\": 39,\n      \"content\": \"pinlun\",\n      \"addTime\": 1587217428,\n      \"addTimeTitle\": \"2020-04-18\",\n      \"userId\": 10,\n      \"stamp\": 0,\n      \"praiseCount\": 0,\n      \"userinfo\": {\n        \"id\": 0,\n        \"name\": \"\",\n        \"addTime\": 0,\n        \"avatar\": \"\"\n      },\n      \"episodesId\": 1\n    },\n    {\n      \"id\": 38,\n      \"content\": \"qq\",\n      \"addTime\": 1587217316,\n      \"addTimeTitle\": \"2020-04-18\",\n      \"userId\": 10,\n      \"stamp\": 0,\n      \"praiseCount\": 0,\n      \"userinfo\": {\n        \"id\": 0,\n        \"name\": \"\",\n        \"addTime\": 0,\n        \"avatar\": \"\"\n      },\n      \"episodesId\": 1\n    },\n    {\n      \"id\": 37,\n      \"content\": \"buco\",\n      \"addTime\": 1587217184,\n      \"addTimeTitle\": \"2020-04-18\",\n      \"userId\": 10,\n      \"stamp\": 0,\n      \"praiseCount\": 0,\n      \"userinfo\": {\n        \"id\": 0,\n        \"name\": \"\",\n        \"addTime\": 0,\n        \"avatar\": \"\"\n      },\n      \"episodesId\": 1\n    },\n    {\n      \"id\": 36,\n      \"content\": \"2\",\n      \"addTime\": 1586686812,\n      \"addTimeTitle\": \"2020-04-12\",\n      \"userId\": 10,\n      \"stamp\": 0,\n      \"praiseCount\": 0,\n      \"userinfo\": {\n        \"id\": 0,\n        \"name\": \"\",\n        \"addTime\": 0,\n        \"avatar\": \"\"\n      },\n      \"episodesId\": 1\n    },\n    {\n      \"id\": 35,\n      \"content\": \"1\",\n      \"addTime\": 1586686560,\n      \"addTimeTitle\": \"2020-04-12\",\n      \"userId\": 10,\n      \"stamp\": 0,\n      \"praiseCount\": 0,\n      \"userinfo\": {\n        \"id\": 0,\n        \"name\": \"\",\n        \"addTime\": 0,\n        \"avatar\": \"\"\n      },\n      \"episodesId\": 1\n    }\n  ],\n  \"count\": 20\n"
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

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» items|null|true|none||none|
|» count|integer|true|none||none|

# 消息模块/消息通知

## POST 通知用户

POST /send/message

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|uids|query|string| 是 |none|
|content|query|string| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "msg": "发送成功",
  "items": "",
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
|» items|string|true|none||none|
|» count|integer|true|none||none|

