package model

import (
	"fmt"
	"errors"
	"encoding/json"
	"hash/crc32"
//	"strconv"

    "github.com/micro/go-log"
	"github.com/go-xorm/core"

	m "tenno.ucenter/common/mysql"
//	redis "tenno.ucenter/common/redis"
//	L0 "tenno.ucenter/common/localcache"
	L1 "tenno.ucenter/common/memcache"
)


var (
	errWxUinonidNull = errors.New("微信unionid不能为空")
)


const (
	//用户微信账号
	KEY_MC_USER_RELATION_WX string = "key_mc_user_relation_wx_%d"
)



type UserRelationWx struct{
	Uid uint64
	Unionid string
	CreateAt uint32 `xorm:"created"`
	UpdateAt uint32 `xorm:"updated"`
}


func init(){

}


//查询单条数据
func (wx *UserRelationWx)GetOne(unionid string) error{
	log.Log("model UserRelationWx.GetOne request, unionid:", unionid)

	if unionid == "" {
		return errPhoneNull
	}

	tbMapper := core.NewSuffixMapper(core.SnakeMapper{}, wx.getSuffix(unionid))
	m.DB.SetTableMapper(tbMapper)

	_, err := m.DB.Where("unionid = ?", unionid).Get(wx)
	if  err != nil{
		log.Log("mysql error:", err)
	}
	
	return err
}


//缓存查询单条数据
func (wx *UserRelationWx)CacheGetOne(unionid string) error{
	log.Log("model UserRelationWx.CacheGetOne request, unionid:", unionid)
	key := fmt.Sprintf(KEY_MC_USER_RELATION_WX, unionid)


	// L1 cache get
	val, _ :=L1.Get(key)
	if val == m.MC_NULL_VALUE {
		return nil
	}
	if val != "" {
		err := json.Unmarshal([]byte(val), wx)
		return err
	}

	// DB get
	if err := wx.GetOne(unionid); err != nil{
		return err
	}

	if wx.Unionid != "" {
		json, _ := json.Marshal(wx)
		L1.Set(key, string(json), 0)
	}else{
		L1.Set(key, m.MC_NULL_VALUE, 0)
	}
	return nil
}


//添加一条数据
func (wx *UserRelationWx)AddOne() (err error){
	log.Log("model UserRelationWx.AddOne request")

	tbMapper := core.NewSuffixMapper(core.SnakeMapper{}, wx.getSuffix(wx.Unionid))
	m.DB.SetTableMapper(tbMapper)
	_, err = m.DB.Insert(wx)

	L1.Del(fmt.Sprintf(KEY_MC_USER_RELATION_WX, wx.Unionid))
	return err
}

//删除微信登陆关系
func (wx *UserRelationWx)DelOne(unionid string) (err error){
	log.Log("model UserRelationWx.DelOne request")

	tbMapper := core.NewSuffixMapper(core.SnakeMapper{}, wx.getSuffix(unionid))
	m.DB.SetTableMapper(tbMapper)
	_, err = m.DB.Where("unionid = ?", unionid).Delete(wx)

	L1.Del(fmt.Sprintf(KEY_MC_USER_RELATION_WX, unionid))
	return err
}


//取分表id
func (w UserRelationWx)getSuffix(unionid string) string{
	hash := uint64(crc32.ChecksumIEEE([]byte(unionid)))
	i := int(hash % m.SUB_TABLE_NUM)
	return fmt.Sprintf("_%d", i)
}