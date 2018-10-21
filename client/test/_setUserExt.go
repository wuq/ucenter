package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"
	ext "tenno.ucenter/proto/ext"
)



func main() {
	service := micro.NewService(micro.Name("tenno.user.ucenter"))
	service.Init()

	// Create new tenno client
	client := ext.NewExtService("", service.Client())

	// Call the tenno
	//rsp, err := client.SetInfo(context.TODO(), &ext.SetRequest{Uid: 180001, BgImg: "data/b.jpg"})
	rsp, err := client.GetInfo(context.TODO(), &ext.GetRequest{Uid: 180001})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Msg)

}