package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"
	auth "tenno.ucenter/proto/auth"
)



func main() {
	service := micro.NewService(micro.Name("tenno.user.ucenter"))
	service.Init()

	// Create new tenno client
	client := auth.NewAuthService("", service.Client())

	// Call the tenno
	rsp, err := client.Login(context.TODO(), &auth.LoginRequest{Phone: 15652191455, Code: 9911, Appid: 10000})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Msg)

}