package handler



import (
	"context"
	"encoding/json"
	"errors"

	"github.com/micro/go-log"

	ext "tenno.ucenter/proto/ext"
	c "tenno.ucenter/common"
	"tenno.ucenter/model"
)

type Ext struct{}

//查询扩展信息
func (e *Ext) GetInfo(ctx context.Context, req *ext.GetRequest, rsp *ext.Response) error {
	log.Log("Received Ext.GetInfo request")

	if req.Uid <= 0 {
		rsp.Msg = c.Rsp(4001, errors.New("uid 为空"), nil)
		return nil
	}

	var err error
	var ext model.UserExt
	data := map[string]interface{}{"bgImg": "", "notification": false}
	if err = ext.CacheGetOne(req.Uid); err != nil {
		rsp.Msg = c.Rsp(4002, err, nil)
		return nil
	}

	if ext.Uid > 0 {
		json.Unmarshal([]byte(ext.Ext), &data)
	}

	rsp.Msg = c.Rsp(1, errors.New("success"), data)
	return nil
}



func (e *Ext) SetInfo(ctx context.Context, req *ext.SetRequest, rsp *ext.Response) error {
	log.Log("Received Ext.SetInfo request")

	if req.Uid <= 0 {
		rsp.Msg = c.Rsp(4001, errors.New("uid 为空"), nil)
		return nil
	}
	
	var ext model.UserExt
	var err error
	data := map[string]interface{}{"bgImg": "", "notification": false}

	//1.查询扩展内容
	if err = ext.CacheGetOne(req.Uid); err != nil {
		rsp.Msg = c.Rsp(4002, err, nil)
		return nil
	}

	if ext.Uid > 0 {
		json.Unmarshal([]byte(ext.Ext), &data)
	}
	
	if req.BgImg != "" {
		data["bgImg"] = req.BgImg
	}
	if req.Notification != false {
		data["notification"] = req.Notification
	}
	json, _ := json.Marshal(data)
	ext.Ext = string(json)

	//2.有记录->更新
	if ext.Uid > 0 {
		if err = ext.UpdateOne(req.Uid); err != nil {
			rsp.Msg = c.Rsp(4001, err, nil)
			return nil
		}
		rsp.Msg = c.Rsp(1, errors.New("success"), data)
		return nil
	}

	//3.没记录->添加
	ext.Uid  = req.Uid
	if err = ext.AddOne(); err != nil {
		rsp.Msg = c.Rsp(4001, err, nil)
		return nil
	}
	
	rsp.Msg = c.Rsp(1, errors.New("success"), data)
	return nil
}
