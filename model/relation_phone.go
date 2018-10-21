package model

import (
	"fmt"
	"errors"
	"encoding/json"
	"hash/crc32"
	"strconv"

    "github.com/micro/go-log"
	"github.com/go-xorm/core"

	m "tenno.ucenter/common/mysql"
//	redis "tenno.ucenter/common/redis"
//	L0 "tenno.ucenter/common/localcache"
	L1 "tenno.ucenter/common/memcache"
)


var (
	errPhoneNull = errors.New("手机号不能为空")
)


const (
	//用户手机账号
	KEY_MC_USER_RELATION_PHONE string = "key_mc_user_relation_phone_%d"
)



type UserRelationPhone struct{
	Uid uint64
	Phone uint64
	CreateAt uint32 `xorm:"created"`
	UpdateAt uint32 `xorm:"updated"`
}


func init(){

}


//查询单条数据
func (b *UserRelationPhone)GetOne(phone uint64) error{
	log.Log("model UserRelationPhone.GetOne request, phone:", phone)

	if phone <= 0{
		return errPhoneNull
	}

	tbMapper := core.NewSuffixMapper(core.SnakeMapper{}, b.getSuffix(phone))
	m.DB.SetTableMapper(tbMapper)

	_, err := m.DB.Where("phone = ?", phone).Get(b)
	if  err != nil{
		log.Log("mysql error:", err)
	}
	
	return err
}


//缓存查询单条数据
func (b *UserRelationPhone)CacheGetOne(phone uint64) error{
	log.Log("model UserRelationPhone.CacheGetOne request, phone:", phone)
	key := fmt.Sprintf(KEY_MC_USER_RELATION_PHONE, phone)


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
	if err := b.GetOne(phone); err != nil{
		return err
	}

	if b.Phone > 0 {
		json, _ := json.Marshal(b)
		L1.Set(key, string(json), 0)
	}else{
		L1.Set(key, m.MC_NULL_VALUE, 0)
	}
	return nil
}


//添加一条数据
func (b *UserRelationPhone)AddOne() (err error){
	log.Log("model UserRelationPhone.AddOne request")

	tbMapper := core.NewSuffixMapper(core.SnakeMapper{}, b.getSuffix(b.Phone))
	m.DB.SetTableMapper(tbMapper)
	_, err = m.DB.Insert(b)

	L1.Del(fmt.Sprintf(KEY_MC_USER_RELATION_PHONE, b.Phone))
	return err
}

//删除登陆关系
func (b *UserRelationPhone)DelOne(phone uint64) (err error){
	log.Log("model UserRelationPhone.DelOne request")

	tbMapper := core.NewSuffixMapper(core.SnakeMapper{}, b.getSuffix(phone))
	m.DB.SetTableMapper(tbMapper)
	_, err = m.DB.Where("phone = ?", phone).Delete(b)

	L1.Del(fmt.Sprintf(KEY_MC_USER_RELATION_PHONE, phone))
	return err
}


//取分表id
func (b UserRelationPhone)getSuffix(phone uint64) string{
	str  := strconv.FormatUint(phone, 10)
	hash := uint64(crc32.ChecksumIEEE([]byte(str)))
	i := int(hash % m.SUB_TABLE_NUM)
	return fmt.Sprintf("_%d", i)
}