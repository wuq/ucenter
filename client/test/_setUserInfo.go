package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"
	base "tenno.ucenter/proto/base"
)



func main() {
	service := micro.NewService(micro.Name("tenno.user.ucenter"))
	service.Init()

	// Create new tenno client
	client := base.NewBaseService("", service.Client())

	// Call the tenno
	rsp, err := client.SetUserInfo(context.TODO(), &base.SetRequest{Uid: 180001, NickName: "dio911"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Msg)

}