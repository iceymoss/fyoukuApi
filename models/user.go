package models

import (
	redisClient "fyoukuApi/services/redis"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

// User 表结构
type User struct {
	Id       int
	Name     string
	Password string
	Status   int
	Mobile   string
	AddTime  int64
	Avatar   string
}

type UserInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	AddTime int64  `json:"addTime"`
	Avatar  string `json:"avatar"`
}

func init() {
	orm.RegisterModel(new(User))
}

// GetUserByMobile 查询用户
func GetUserByMobile(mobile string) bool {
	o := orm.NewOrm()
	user := User{Mobile: mobile}
	err := o.Read(&user, "mobile")
	if err == orm.ErrNoRows {
		return false
	} else if err == orm.ErrMissPK {
		return false
	}
	return true
}

func UserSave(mobile string, password string) error {
	o := orm.NewOrm()
	user := User{
		Name:     "",
		Password: password,
		Status:   1,
		Mobile:   mobile,
		AddTime:  time.Now().Unix(),
	}
	_, err := o.Insert(&user)
	return err
}

// IsMobileLogin 登录功能
func IsMobileLogin(mobile string, password string) (int, string) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("mobile", mobile).Filter("password", password).One(&user)
	if err == orm.ErrNoRows {
		return 0, ""
	} else if err == orm.ErrMissPK {
		return 0, ""
	}
	return user.Id, user.Name
}

// GetUserInfo 根据用户ID获取用户信息
func GetUserInfo(uid int) (UserInfo, error) {
	o := orm.NewOrm()
	var user UserInfo
	err := o.Raw("SELECT id,name,add_time,avatar FROM user WHERE id=? LIMIT 1", uid).QueryRow(&user)
	return user, err
}

func RedisGetUserInfo(uid int) (UserInfo, error) {
	var userInfo UserInfo
	var err error
	//到Redis查询，没有则去MySQL查询然后写入Redis中
	conn := redisClient.PoolConnect()
	defer conn.Close()

	key := "user:id:" + strconv.Itoa(uid)
	exists, err := redis.Bool(conn.Do("exists", key))
	if exists { //key存在
		res, _ := redis.Values(conn.Do("hgetall", key))
		err = redis.ScanStruct(res, &userInfo)
	} else { //key不存在，去MySQL查，写入Redis中
		o := orm.NewOrm()
		err = o.Raw("SELECT id,name,add_time,avatar FROM user WHERE id=? LIMIT 1", uid).QueryRow(&userInfo)
		if err == nil {
			_, err = conn.Do("hmset", redis.Args{key}.AddFlat(&userInfo)...)
			if err == nil {
				conn.Do("expire", key, 86400*time.Second)
			}
		}
	}
	return userInfo, err
}
