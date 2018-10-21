package f

import (
	"context"
	"strconv"
	"encoding/json"
//	"fmt"

	"github.com/micro/go-log"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"

	base "tenno.ucenter/proto/base"
)


var bsc base.BaseService

type BaseClient struct {}



func init() {
	bsc = base.NewBaseService("", client.DefaultClient)
}


func (b BaseClient)GetUserInfo(c *gin.Context) {
	log.Log("Received BaseClient.GetUserInfo API request")

	//1.接收参数
	uid, _ := strconv.ParseUint(c.Param("uid"), 10, 64)
	appid, _ := strconv.ParseUint(c.PostForm("appid"), 10, 32)
	token := c.PostForm("token")

	//2.验证token
	msg := make(map[string]interface{})
	var auth AuthClient
	var code float64 = 1
	res := auth.Check(id, uint32(appid), token)
	json.Unmarshal([]byte(res.Msg), &msg)

	if msg["code"].(float64) != code {
		c.String(200, "%v", res.Msg)
		return
	}

	//3.返回信息
	rsp, _ := bsc.GetUserInfo(context.TODO(), &base.GetRequest{
		Uid: id,
	})

	c.String(200, "%v", rsp.Msg)
}


func (b BaseClient)SetUserInfo(c *gin.Context) {
	log.Log("Received BaseClient.SetUserInfo API request")

	//1.接收参数
	uid, _ := strconv.ParseUint(c.Param("uid"), 10, 64)
	appid, _ := strconv.ParseUint(c.PostForm("appid"), 10, 32)
	token := c.PostForm("token")

	//2.验证token
	msg := make(map[string]interface{})
	var auth AuthClient
	var code float64 = 1
	res := auth.Check(id, uint32(appid), token)
	json.Unmarshal([]byte(res.Msg), &msg)

	if msg["code"].(float64) != code {
		c.String(200, "%v", res.Msg)
		return
	}

	//2. 返回信息
	rsp, _ := bsc.SetUserInfo(context.TODO(), &base.GetRequest{
		Uid: id,
	})

	c.String(200, "%v", rsp.Msg)
}