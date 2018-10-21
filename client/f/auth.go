package f

import (
	"context"
	"strconv"

	"github.com/micro/go-log"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"

	auth "tenno.ucenter/proto/auth"
)


var asc auth.AuthService

type AuthClient struct {}



func init() {
	asc = auth.NewAuthService("", client.DefaultClient)
}


func (b AuthClient)Login(c *gin.Context) {
	log.Log("Received Auth.Login API request")

	phone, _ := strconv.ParseUint(c.PostForm("phone"), 10, 64)
	appid, _ := strconv.ParseUint(c.PostForm("appid"), 10, 32)
	code, _ := strconv.ParseInt(c.PostForm("code"), 10, 32)

	rsp, _ := asc.Login(context.TODO(), &auth.LoginRequest{
		Phone: phone,
		Appid: uint32(appid),
		Code: int32(code),
	})

	c.String(200, "%v", rsp.Msg)
}


func (b AuthClient)Check(uid uint64, appid uint32, token string) *auth.Response {
	log.Log("Received Auth.Check API request")

	rsp, _ := asc.Check(context.TODO(), &auth.CheckRequest{
		Uid: uid,
		Appid: appid,
		Token: token,
	})

	return rsp
}