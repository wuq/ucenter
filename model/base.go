package model

import (
	"fmt"
	"errors"
	"encoding/json"
	"time"
	"strconv"

    "github.com/micro/go-log"
	"github.com/go-xorm/core"

	m "tenno.ucenter/common/mysql"
//	redis "tenno.ucenter/common/redis"
//	L0 "tenno.ucenter/common/localcache"
	L1 "tenno.ucenter/common/memcache"
)


var (
	errUidMiss = errors.New("uid 不能为空")
	errUidRange = errors.New("无效的uid范围")
	DefaultAvatar string = "https://apic.douyucdn.cn/upload/avatar/face/201606/21/9d542c9e6e2862da9b58f0b278ce9f91_big.jpg?rltime"
)


const (
	//用户信息
	KEY_MC_USER_BASE_INFO string = "key_mc_user_base_info_%d"
	//KEY_MC_TICKET_MAXID string = "key_mc_user_maxid"
)



const (
	GENDER_DEFAULT = 0
	GENDER_MALE = 1
	GENDER_FEMALE = 2
)


const (
	STATE_NORMAL = 0	//正常
	STATE_SEAL = 1		//封号
	STATE_CROSSOUT = 2  //注销
	STATE_FREEZE = 3	//冻结
)


type UserBase struct{
	Uid uint64
	NickName string
	UserName string
	RealName string
	Gender int32
	Age int32
	Phone uint64
	Email string
	Appid uint32
	Country string
	Province string
	City string
	Language string
	Avatar string
	State uint8
	CreateAt uint32 `xorm:"created"`
	UpdateAt uint32 `xorm:"updated"`
}


func init(){

}


//查询单条数据
func (b *UserBase)GetOne(uid uint64) error {
	log.Log("model UserBase.GetOne request, uid:", uid)
	if uid <= 0{
		return errUidMiss
	}

	maxid, _ := L1.Get(KEY_MC_TICKET_MAXID)
	mid, _ := strconv.ParseUint(maxid, 10, 64)
	if mid > 0 && uid > mid {
		return errUidRange
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
func (b *UserBase)CacheGetOne(uid uint64) error {
	log.Log("model UserBase.CacheGetOne request, uid:", uid)
	key := fmt.Sprintf(KEY_MC_USER_BASE_INFO, uid)

	maxid, _ := L1.Get(KEY_MC_TICKET_MAXID)
	mid, _ := strconv.ParseUint(maxid, 10, 64)
	if mid > 0 && uid > mid {
		return errUidRange
	}

/*
	// L0 cache get
	v, _ := L0.Get(key)
	if v == m.MC_NULL_VALUE {
		return nil
	}
	if v != "" {
		err := json.Unmarshal([]byte(v), b)
		return err
	}
*/

	// L1 cache get
	val, _ :=L1.Get(key)
	if val == m.MC_NULL_VALUE {
		//L0.Set(key, m.MC_NULL_VALUE)
		return nil
	}
	if val != "" {
		//L0.Set(key, string(val))
		err := json.Unmarshal([]byte(val), b)
		return err
	}

	// DB get
	if err := b.GetOne(uid); err != nil{
		return err
	}

	if b.Uid > 0 {
		json, _ := json.Marshal(b)
		//L0.Set(key, string(json))
		L1.Set(key, string(json), 0)
	}else{
		//L0.Set(key, m.MC_NULL_VALUE)
		L1.Set(key, m.MC_NULL_VALUE, 0)
	}
	return nil
}


//添加一条数据
func (b *UserBase)AddOne()  error {
	log.Log("model UserBase.AddOne request")

	tbMapper := core.NewSuffixMapper(core.SnakeMapper{}, b.getSuffix(b.Uid))
	m.DB.SetTableMapper(tbMapper)
	//m.DB.ShowSQL(true)
	_, err = m.DB.Insert(b)

	L1.Del(fmt.Sprintf(KEY_MC_USER_BASE_INFO, b.Uid))
	return err
}

//更新一条数据
func (b *UserBase)UpdateOne(uid uint64) error {
	log.Log("model UserBase.UpdateOne request, uid:", uid)

	b.UpdateAt = uint32(time.Now().Unix())
	tbMapper := core.NewSuffixMapper(core.SnakeMapper{}, b.getSuffix(b.Uid))
	m.DB.SetTableMapper(tbMapper)
	_, err = m.DB.Where("uid = ?", uid).Update(b)

	L1.Del(fmt.Sprintf(KEY_MC_USER_BASE_INFO, uid))
	return err 
}


//取分表id
func (b UserBase)getSuffix(uid uint64) string{
	i := int(uid % m.SUB_TABLE_NUM)
	return fmt.Sprintf("_%d", i)
}