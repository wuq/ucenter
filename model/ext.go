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




const (
	//用户扩展信息
	KEY_MC_USER_EXT_INFO string = "key_mc_user_ext_%d"
)



type UserExt struct{
	Uid uint64
	Ext string
	CreateAt uint32 `xorm:"created"`
	UpdateAt uint32 `xorm:"updated"`
}


func init(){

}


//查询单条数据
func (e *UserExt)GetOne(uid uint64) error{
	log.Log("model UserExt.GetOne request, uid:", uid)

	if uid <= 0{
		return errors.New("uid 不能为空")
	}

	tbMapper := core.NewSuffixMapper(core.SnakeMapper{}, e.getSuffix(uid))
	m.DB.SetTableMapper(tbMapper)

	_, err := m.DB.Where("uid = ?", uid).Get(e)
	if  err != nil{
		log.Log("mysql error:", err)
	}
	
	return err
}


//缓存查询单条数据
func (e *UserExt)CacheGetOne(uid uint64) error{
	log.Log("model UserExt.CacheGetOne request, uid:", uid)
	key := fmt.Sprintf(KEY_MC_USER_EXT_INFO, uid)


	// L1 cache get
	val, _ :=L1.Get(key)

	if val == m.MC_NULL_VALUE {
		return nil
	}
	if val != "" {
		err := json.Unmarshal([]byte(val), e)
		return err
	}

	// DB get
	if err := e.GetOne(uid); err != nil{
		return err
	}

	if e.Uid > 0 {
		json, _ := json.Marshal(e)
		L1.Set(key, string(json), 0)
	}else{
		L1.Set(key, m.MC_NULL_VALUE, 0)
	}
	return nil
}


//添加一条数据
func (e *UserExt)AddOne() error {
	log.Log("model UserExt.AddOne request")

	tbMapper := core.NewSuffixMapper(core.SnakeMapper{}, e.getSuffix(e.Uid))
	m.DB.SetTableMapper(tbMapper)
	_, err = m.DB.Insert(e)
	
	L1.Del(fmt.Sprintf(KEY_MC_USER_EXT_INFO, e.Uid))
	return err
}


//更新一条数据
func (e *UserExt)UpdateOne(uid uint64) error {
	log.Log("model UserExt.UpdateOne request, uid:", uid)

	tbMapper := core.NewSuffixMapper(core.SnakeMapper{}, e.getSuffix(uid))
	m.DB.SetTableMapper(tbMapper)
	_, err = m.DB.Where("uid = ?", uid).Update(e)

	L1.Del(fmt.Sprintf(KEY_MC_USER_EXT_INFO, uid))
	return err 
}


//取分表id
func (e UserExt)getSuffix(uid uint64) string{
	i := int(uid % m.SUB_TABLE_NUM)
	return fmt.Sprintf("_%d", i)
}