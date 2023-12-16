[toc]



## 概述

仿优酷视频项目(fyouku)的后台服务，这是一个类优酷视频的的简版项目，基于Go开发。

## 功能特点

### 主要模块

- 视频模块
- 广告模块
- 弹幕模块
- 评论模块
- 用户模块
- 消息模块

### 主要特点

* 基于Redis缓存提示部分接口性能。
* 基于MQ实现消息和排行榜计算异步处理，提升接口性能。
* 基于ES完成视频关键字词检索，提高接口性能和用户体验。
* 基于WX完成弹幕实时推送和实时在线用户，提高了弹幕实时性和节省资源。
* 基于Go协程完成部分接口并发查询聚合，提高接口性能。
* 基于MMDB实现内部离线IP检查和定位，提示系统维护性和拓展性。
* 基于Aliyun云点播一站式视频管理。

### 目录结构

```
.
├── README.md
├── cmd.go  //子任务入口
├── command //子任务实现
│   ├── cmd.go
│   └── mq
│       └── consume.go
├── conf    //配置文件
│   └── app.conf
├── controllers  //控制器
│   ├── aliyun.go
│   ├── barrage.go
│   ├── base.go
│   ├── comment.go
│   ├── common.go
│   ├── geoip.go
│   ├── top.go
│   ├── user.go
│   └── video.go
├── fyouku.sql   //项目数据库DDL以及数据    
├── go.mod       //项目依赖
├── go.sum       //项目依赖
├── lastupdate.tmp  
├── main.go      //server启动入口
├── models       //数据库models
│   ├── advert.go
│   ├── barrage.go
│   ├── base.go
│   ├── comment.go
│   ├── message.go
│   ├── user.go
│   └── video.go
├── mq           //mq消费逻辑，已迁移至command
│   ├── send
│   │   └── main.go
│   └── top
│       └── main.go
├── resource    //ip数据库
│   └── libs
│       ├── GeoLite2-City.mmdb
│       └── GeoLite2-Country.mmdb
├── routers     //路由配置
│   └── router.go
├── services    //组件
│   ├── es
│   │   └── Es.go
│   ├── mq
│   │   └── Mq.go
│   └── redis
│       └── Redis.go
└── tests      //接口测试，技术demo
    ├── cobra
    │   ├── cmd
    │   │   └── root.go
    │   └── main.go
    ├── default_test.go
    └── es_test
        └── es.go
```



## 技术栈

> Go、Beego、Gorm、Cobra、Geoip2、MySQL、Redis、RabbitMQ、ElasticSearch、Docker



## 快速开始

### 下载项目

```shell
git clone https://github.com/iceymoss/fyoukuApi.git
```

### 构建数据库

直接使用```fyouku.sql```创建表和导入数据。

### 创建ES索引

对ES和kibana的安装可以参考：[Elasticsearch学习总结](https://learnku.com/articles/72845)，注意还要需要安装ik分词器，直接将ik分词器项目下载至elasticsearch的plugins目录中。

kibana创建index（这里只是方法之一）：

```json
# 创建index和mapping
PUT /fyouku_video
{
    "mappings": {
      "properties": {
        "id": {
          "type": "integer"
        },
        "title": {
          "type": "text",
          "analyzer": "ik_max_word",
          "search_analyzer": "ik_max_word"
        },
        "sub_title": {
            "type": "text",
            "analyzer": "ik_max_word",
            "search_analyzer": "ik_max_word"
        },
        "add_time": {
            "type": "integer"
        },
         "img": {
        "type": "keyword"
        },
        "img1": {
            "type": "keyword"
        },
        "episodes_count": {
            "type": "integer"
        },
        "is_end": {
            "type": "integer"
        },
        "channel_id": {
            "type": "integer"
        },
        "status": {
            "type": "integer"
        },
        "region_id": {
            "type": "integer"
        },
        "type_id": {
            "type": "integer"
        },
        "episodes_update_time": {
            "type": "date"
        },
        "comment": {
            "type": "integer"
        },
        "user_id": {
            "type": "integer"
        },
        "is_recommend": {
            "type": "integer"
        }
      }
    }
}

```

### 修改配置

连接到你自己的相关组件

```
appname = demo
httpport = 8099
runmode = dev

[dev]
defaultdb = root:123456789@(127.0.0.1:3306)/fyouku?charset=utf8
redisdb = 127.0.0.1:6379
rabbitmq = amqp://guest:guest@127.0.0.1:5672/

[prod]
defaultdb = root:123456789@(127.0.0.1:3306)/fyouku?charset=utf8
redisdb = 127.0.0.1:6379
rabbitmq = amqp://guest:guest@127.0.0.1:5672/

```

**注意**

在项目中：/controllers/aliyun.go你需要修改您的阿里云key，当然这里最好将其配置到配置文件或者环境变量中：

```go
var (
	accessKeyId     = "Gky82i3Q"
	accessKeySecret = "iYQIWSFejU"
)
```

### 拉取依赖

```shell
go get 
```

### 启动项目

排行榜MQ任务

```shell
go run cmd.go rank
```

消息MQ任务

```shell
go run cmd.go send
```

启动服务

```shell
go run main.go
```