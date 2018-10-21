package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"
	base "tenno.ucenter/proto/base"
)



func main() {
	service := micro.NewService(micro.Name("go.micro.srv.tenno.ucenter"))
	service.Init()

	// Create new tenno client
	client := base.NewBaseService("", service.Client())

	// Call the tenno
	rsp, err := client.GetUserInfo(context.TODO(), &base.GetRequest{Uid: 180001})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Msg)

}