package model

import (
	"fmt"
	"errors"
	"encoding/json"
//	"time"	
//	"hash/crc32"
//	"strconv"

    "github.com/micro/go-log"
	"github.com/go-xorm/core"

	m "tenno.ucenter/common/mysql"
//	redis "tenno.ucenter/common/redis"
//	L0 "tenno.ucenter/common/localcache"
	L1 "tenno.ucenter/common/memcache"
)


var (
	errUidNull = errors.New("uid不能为空")
)


const (
	//用户绑定信息
	KEY_MC_USER_BIND_INFO string = "key_mc_user_bind_%d"
)



type UserBind struct{
	Uid uint64
	Data string
	CreateAt uint32 `xorm:"created"`
	UpdateAt uint32 `xorm:"updated"`
}


func init(){

}


//查询单条数据
func (b *UserBind)GetOne(uid uint64) error{
	log.Log("model UserBind.GetOne request, uid:", uid)

	if uid <= 0{
		return errUidNull
	}

	tbMapper := core.NewSuffixMapper(core.SnakeMapper{}, b.getSuffix(uid))
	m.DB.SetTableMapper(tbMapper)

	_, err := m.DB.Where("uid = ?", uid).Get(b)
	if  err != nil{
		log.Log("mysql error:", err)
	}
	
	return err
}


//缓存查询单条数据
func (b *UserBind)CacheGetOne(uid uint64) error{
	log.Log("model UserBind.CacheGetOne request, uid:", uid)
	key := fmt.Sprintf(KEY_MC_USER_BIND_INFO, uid)


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
	if err := b.GetOne(uid); err != nil{
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
func (b *UserBind)AddOne() error {
	log.Log("model UserBind.AddOne request")

	tbMapper := core.NewSuffixMapper(core.SnakeMapper{}, b.getSuffix(b.Uid))
	m.DB.SetTableMapper(tbMapper)
	_, err = m.DB.Insert(b)
	
	L1.Del(fmt.Sprintf(KEY_MC_USER_BIND_INFO, b.Uid))
	return err
}


//取分表id
func (b UserBind)getSuffix(uid uint64) string{
	i := int(uid % m.SUB_TABLE_NUM)
	return fmt.Sprintf("_%d", i)
}