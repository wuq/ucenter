package model

import (
	"fmt"
	"errors"
	"encoding/json"
	"time"
//	"hash/crc32"
//	"strconv"

    "github.com/micro/go-log"
	"github.com/go-xorm/core"

	m "tenno.ucenter/common/mysql"
//	redis "tenno.ucenter/common/redis"
//	L0 "tenno.ucenter/common/localcache"
	L1 "tenno.ucenter/common/memcache"
)


const (
	//用户绑定信息
	KEY_MC_USER_AUTH_INFO string = "key_mc_user_auth_%d_%d"
)


const (
	LOGIN_TYPE_PHONE = 1
	LOGIN_TYPE_WX = 2
	LOGIN_TYPE_QQ = 3
	LOGIN_TYPE_WB = 4
	LOGIN_TYPE_MAIL = 5
)


type UserAuth struct{
	Uid uint64
	Appid	uint32
	Type    uint8
	AccessToken string
	AccessSecret string
	Expire uint32
	CreateAt uint32 `xorm:"created"`
	UpdateAt uint32 `xorm:"updated"`
}


func init(){

}


//查询单条数据
func (b *UserAuth)GetOne(uid uint64, appid uint32) error{
	log.Log("model UserAuth.GetOne request, uid:", uid, "appid:", appid)

	if uid <= 0{
		return errors.New("uid不能为空")
	}

	tbMapper := core.NewSuffixMapper(core.SnakeMapper{}, b.getSuffix(uid))
	m.DB.SetTableMapper(tbMapper)

	_, err := m.DB.Where("uid = ? AND appid = ?", uid, appid).Get(b)
	if  err != nil{
		log.Log("mysql error:", err)
	}
	
	return err
}


//缓存查询单条数据
func (b *UserAuth)CacheGetOne(uid uint64, appid uint32) error{
	log.Log("model UserAuth.CacheGetOne request, uid:", uid, "appid:", appid)
	key := fmt.Sprintf(KEY_MC_USER_AUTH_INFO, uid, appid)

	// L1 cache get
	val, _ :=L1.Get(key)

	if val == m.MC_NULL_VALUE {
		return nil
	}
	if val != "" {
		err := json.Unmarshal([]byte(val), b)
		return err
	}

	// DB get
	if err := b.GetOne(uid, appid); err != nil{
		return err
	}

	if b.Uid > 0 {
		json, _ := json.Marshal(b)
		L1.Set(key, string(json), 0)
	}else{
		L1.Set(key, m.MC_NULL_VALUE, 0)
	}
	return nil
}


//添加一条数据
func (b *UserAuth)AddOne() (err error){
	log.Log("model UserAuth.AddOne request")

	tbMapper := core.NewSuffixMapper(core.SnakeMapper{}, b.getSuffix(b.Uid))
	m.DB.SetTableMapper(tbMapper)
	_, err = m.DB.Insert(b)
	
	L1.Del(fmt.Sprintf(KEY_MC_USER_AUTH_INFO, b.Uid, b.Appid))
	return err
}


//更新一条数据
func (b *UserAuth)UpdateOne(uid uint64, appid uint32) error {
	log.Log("model UserAuth.UpdateOne request, uid:", uid, "appid:", appid)

	b.UpdateAt = uint32(time.Now().Unix())
	tbMapper := core.NewSuffixMapper(core.SnakeMapper{}, b.getSuffix(b.Uid))
	m.DB.SetTableMapper(tbMapper)
	_, err = m.DB.Where("uid = ? AND appid = ?", uid, appid).Update(b)

	L1.Del(fmt.Sprintf(KEY_MC_USER_AUTH_INFO, uid, appid))
	return err 
}


//取分表id
func (b UserAuth)getSuffix(uid uint64) string{
	i := int(uid % m.SUB_TABLE_NUM)
	return fmt.Sprintf("_%d", i)
}