package handler



import (
	"context"
	"errors"

	"github.com/micro/go-log"

	base "tenno.ucenter/proto/base"
	c "tenno.ucenter/common"
	"tenno.ucenter/model"
)

type Base struct{}


//查询用户信息
func (b *Base)GetUserInfo(ctx context.Context, req *base.GetRequest, rsp *base.Response) error {
	log.Log("Received Base.GetUserInfo request")

	var err error
	var user model.UserBase
	if err = user.CacheGetOne(req.Uid); err != nil{
		rsp.Msg = c.Rsp(4004, err, nil)
		return nil
	}

	rsp.Msg = c.Rsp(1, errors.New("success"), user)
	return nil
}


//更新用户信息
func (b *Base)SetUserInfo(ctx context.Context, req *base.SetRequest, rsp *base.Response) error {
	log.Log("Received Base.SetUserInfo request")

	var err error
	var user model.UserBase

	//1.判断用户是否存在
	if err = user.CacheGetOne(req.Uid); err != nil{
		rsp.Msg = c.Rsp(4004, err, nil)
		return nil
	}

	if user.Uid <= 0 {
		rsp.Msg = c.Rsp(4004, errors.New("用户不存在"), nil)
		return nil
	}

	//2.填充更新字段
	if req.NickName != "" {
		user.NickName = req.NickName
	}
	if req.UserName != "" {
		user.UserName = req.UserName
	}
	if req.RealName != "" {
		user.RealName = req.RealName
	}
	if req.Gender != 0 {
		user.Gender = req.Gender
	}
	if req.Age != 0 {
		user.Age = req.Age
	}
	if req.Phone != 0 {
		user.Phone = req.Phone
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Country != "" {
		user.Country = req.Country
	}
	if req.Province != "" {
		user.Province = req.Province
	}
	if req.City != "" {
		user.City = req.City
	}
	if req.Language != "" {
		user.Language = req.Language
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.State != 0 {
		user.State = uint8(req.State)
	}

	//更新
	if err = user.UpdateOne(req.Uid); err != nil {
		rsp.Msg = c.Rsp(1, err, nil)
		return nil
	}
	rsp.Msg = c.Rsp(1, errors.New("success"), nil)
	return nil
}