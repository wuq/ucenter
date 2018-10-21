package handler

import (
	"context"
	"regexp"
	"errors"
	"encoding/json"
	"strconv"
	"time"
	"fmt"


	"github.com/micro/go-log"

	auth "tenno.ucenter/proto/auth"
	c "tenno.ucenter/common"
	"tenno.ucenter/model"
)


type Auth struct{}


//手机登陆
func (a *Auth) Login(ctx context.Context, req *auth.LoginRequest, rsp *auth.Response) error {
	log.Log("Received Auth.Login request")

	var err error
	var user model.UserBase
	var pRelation model.UserRelationPhone
	var uid uint64

	//1.验证短信
	_, err = a.checkSMS(req.Phone, req.Code)
	if err != nil{
		rsp.Msg = c.Rsp(4001, err, nil)
		return nil
	}

	//2.是否有手机账号
	if err = pRelation.CacheGetOne(req.Phone); err != nil{
		rsp.Msg = c.Rsp(4002, err, nil)
		return nil
	}

	//3.没有账号->去注册
	uid = pRelation.Uid
	if pRelation.Uid <= 0 {

		randNick := fmt.Sprintf("user_%d", req.Phone)
		uid, err = a.reg(model.LOGIN_TYPE_PHONE, randNick, 0, req.Phone, "", "", req.Appid, "")
		if err != nil {
			rsp.Msg = c.Rsp(4003, err, nil)
			return nil
		}
	}

	//4.有账号->生成auth	
	if err = user.CacheGetOne(uid); err != nil{
		rsp.Msg = c.Rsp(4004, err, nil)
		return nil
	}
	auth, err := a.genAuth(uid, req.Appid, model.LOGIN_TYPE_PHONE)
	if err != nil {
		rsp.Msg = c.Rsp(4005, err, nil)
		return nil
	}
	res := map[string]interface{}{"userInfo" : user,"authInfo" : auth}

	rsp.Msg = c.Rsp(1, errors.New("success"), res)
	return nil
}


//三方登陆
func (a *Auth) SDKLogin(ctx context.Context, req *auth.SDKLoginRequest, rsp *auth.Response) error {
	log.Log("Received Auth.SDKLogin request")

	if req.Unionid == "" {
		rsp.Msg = c.Rsp(4001, errors.New("unionid为空"), nil)
		return nil
	}

	var uid uint64
	var err error
	var user model.UserBase
	var wxRelation model.UserRelationWx

	//1.是否有微信登陆关系
	if err = wxRelation.CacheGetOne(req.Unionid); err != nil{
		rsp.Msg = c.Rsp(4002, err, nil)
		return nil
	}
	uid = wxRelation.Uid
	
	//2.没有->去注册
	if uid <= 0 {
		uid, err = a.reg(uint8(req.Type), req.NickName, req.Gender, 0, req.Unionid, "", req.Appid, req.Avatar)
		if err != nil {
			rsp.Msg = c.Rsp(4003, err, nil)
			return nil
		}
	}

	//3.有账号->生成auth	
	if err = user.CacheGetOne(uid); err != nil{
		rsp.Msg = c.Rsp(4004, err, nil)
		return nil
	}
	auth, err := a.genAuth(uid, req.Appid, model.LOGIN_TYPE_PHONE)
	if err != nil {
		rsp.Msg = c.Rsp(4005, err, nil)
		return nil
	}
	res := map[string]interface{}{"userInfo" : user,"authInfo" : auth}

	rsp.Msg = c.Rsp(1, errors.New("success"), res)
	return nil
}


//检查token是否有效
func (a *Auth) Check(ctx context.Context, req *auth.CheckRequest, rsp *auth.Response) error {
	log.Log("Received Auth.Check request")

	//1.验参数
	if req.Uid <= 0 {
		rsp.Msg = c.Rsp(4001, errors.New("uid为空"), nil)
		return nil
	}
	if req.Appid <= 0 {
		rsp.Msg = c.Rsp(4002, errors.New("appid为空"), nil)
		return nil
	}

	//2.查auth记录
	var auth model.UserAuth
	var err error
	if err = auth.CacheGetOne(req.Uid, uint32(req.Appid)); err != nil {
		rsp.Msg = c.Rsp(4003, err, nil)
		return nil
	}

	//3.判断token
	if auth.Uid <= 0 {
		rsp.Msg = c.Rsp(4004, errors.New("没有授权记录"), nil)
		return nil
	}
	if auth.AccessToken != req.Token {
		rsp.Msg = c.Rsp(4004, errors.New("无效的token"), nil)
		return nil
	}
	if auth.Expire <= uint32(time.Now().Unix()) {
		rsp.Msg = c.Rsp(4005, errors.New("token已过期,请重新登陆"), nil)
		return nil
	}

	rsp.Msg = c.Rsp(1, errors.New("success"), nil)
	return nil
}


//登出
func (a *Auth) Logout(ctx context.Context, req *auth.LogoutRequest, rsp *auth.Response) error {
	log.Log("Received Auth.Logout request")

	//1.验参数
	if req.Uid <= 0 {
		rsp.Msg = c.Rsp(4001, errors.New("uid为空"), nil)
		return nil
	}
	if req.Appid <= 0 {
		rsp.Msg = c.Rsp(4002, errors.New("appid为空"), nil)
		return nil
	}

	//2.查auth记录
	var auth model.UserAuth
	var err error
	if err = auth.CacheGetOne(req.Uid, uint32(req.Appid)); err != nil {
		rsp.Msg = c.Rsp(4003, err, nil)
		return nil
	}

	//3.验token
	if auth.Uid <= 0 {
		rsp.Msg = c.Rsp(4004, errors.New("没有授权记录"), nil)
		return nil
	}
	if auth.AccessToken != req.Token {
		rsp.Msg = c.Rsp(4005, errors.New("无效的token"), nil)
		return nil
	}

	//4.改为过期状态
	auth.Expire = 0
	if err = auth.UpdateOne(req.Uid, uint32(req.Appid)); err != nil {
		rsp.Msg = c.Rsp(4006, err, nil)
		return nil
	}
	
	rsp.Msg = c.Rsp(1, errors.New("success"), nil)
	return nil
}


//发送登陆验证码
func (a *Auth) SendSMS(ctx context.Context, req *auth.SmsRequest, rsp *auth.Response) error {
	log.Log("Received Auth.SendSMS request")
	rsp.Msg = c.Rsp(1, errors.New("success"), nil)
	return nil
}



//检查短信
func (a Auth)checkSMS(phone uint64, code int32) (bool, error){
	regular := "^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}$"
	str := strconv.FormatUint(phone, 10)
	reg := regexp.MustCompile(regular)
 	if !reg.MatchString(str){
 		return false, errors.New("手机格式错误")
 	}

 	//短信功能没写,用默认值
 	var defCode int32 = 9911
 	if code != defCode{
 		return false, errors.New("验证码错误")
 	}

 	return true, nil
}


//注册(仅填充基本信息)
func (a Auth)reg(loginType uint8, nick string, gender int32, phone uint64, unionid string, email string, appid uint32, avatar string) (uid uint64, err error){
	log.Log("Received Auth.reg request")

	if appid <= 0 {
		return 0, errors.New("注册来源不能为空")
	}
	switch loginType{
	case model.LOGIN_TYPE_PHONE:
		if phone <= 0 {
			return 0, errors.New("手机号为空")
		}
	case model.LOGIN_TYPE_WX:
		if unionid == "" {
			return 0, errors.New("unionid 不能为空")
		}
	case model.LOGIN_TYPE_MAIL:
		if email == "" {
			return 0, errors.New("邮箱不能为空")
		}
	default:
		return 0, errors.New("注册类型不能为空")
	}
	
	//var err error
	var t model.Ticket
	//1.申请ID
	if err = t.GetId(); err != nil {
		log.Log("get id err=",err)
		return 0, errors.New("注册失败")
	}

	//2.写入基本信息
	var user model.UserBase
	if avatar == ""{
		avatar = model.DefaultAvatar
	}
	user.Uid = t.Id
	user.NickName = nick
	user.Gender = gender
	user.Phone = phone
	user.Email = email
	user.Appid = appid
	user.Avatar = avatar
	user.State = model.STATE_FREEZE
	if err = user.AddOne(); err != nil {
		return 0, err
	}

	//3.写入登陆关系
	var p model.UserRelationPhone
	if loginType == model.LOGIN_TYPE_PHONE {
		p.Uid = t.Id
		p.Phone = phone
		if err  = p.AddOne(); err != nil {
			return 0, err
		}
	}
	if loginType == model.LOGIN_TYPE_WX {
		//TODO: wx relation
	}

	//4.写入绑定关系
	var bind model.UserBind
	s := strconv.FormatUint(phone, 10)
	d := make(map[int]string)
	d[model.LOGIN_TYPE_PHONE] = s
	json, _ := json.Marshal(d)
	bind.Uid  = t.Id
	bind.Data = string(json)
	if err = bind.AddOne(); err != nil {
		//失败:删除登陆关系
		if loginType == model.LOGIN_TYPE_PHONE{
			p.DelOne(phone)
		}
		return 0, err
	}

	//5.把用户改为正常状态
	var u model.UserBase
	u.State = model.STATE_NORMAL
	u.UpdateOne(t.Id)

	return t.Id, nil
}


//生成auth信息
func (a Auth) genAuth(uid uint64, appid uint32, loginType uint8) (model.UserAuth, error) {
	var auth model.UserAuth
	var err error
	if uid <= 0 {
		return auth, errors.New("uid 不能为空")
	}
	if appid <= 0 {
		return auth, errors.New("appid 不能为空")
	}
	if loginType <= 0 {
		return auth, errors.New("loginType 不能为空")
	}

	token  := c.RandStr(16, "Aa0")
	secret := c.RandStr(32, "Aa0")
	
	if err = auth.CacheGetOne(uid, appid); err != nil {
		return auth, err
	}

	loc, _:= time.LoadLocation("Asia/Shanghai")
	tm := uint32(time.Now().In(loc).Unix()) + 86400
	auth.Type = loginType
	auth.AccessToken  = token
	auth.AccessSecret = secret
	auth.Expire = uint32(tm)

	//已有记录,更新
	if auth.Uid > 0 {
		auth.UpdateOne(uid, appid)
		return auth, nil
	}

	//写入新纪录
	auth.Uid = uid
	auth.Appid = appid
	err = auth.AddOne()
	if err != nil {
		return auth, err
	}

	return auth, nil
}
