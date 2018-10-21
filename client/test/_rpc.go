package main

import (
	"context"
//	"strconv"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	base "tenno.ucenter/proto/base"
)



type BaseClient struct {
	Client base.BaseService
}

//client
func (bc *BaseClient) GetUserInfo(ctx context.Context, req *base.GetRequest, rsp *base.Response) error {
	log.Log("Received BaseClient.GetUserInfo API request")

	// make the request
	response, err := bc.Client.GetUserInfo(ctx, &base.GetRequest{Uid: 180001})
	if err != nil {
		return err
	}

	// set api response
	rsp.Msg = response.Msg
	return nil
}


func (bc *BaseClient) SetUserInfo(ctx context.Context, req *base.SetRequest, rsp *base.Response) error {
	log.Log("Received BaseClient.GetUserInfo API request")

	// make the request
	response, err := bc.Client.SetUserInfo(ctx, &base.SetRequest{Uid: 180001})
	if err != nil {
		return err
	}

	// set api response
	rsp.Msg = response.Msg
	return nil
}





func main() {

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.tenno.ucenter"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	base.RegisterBaseHandler(service.Server(), &BaseClient{
		// Create Service Client
		Client: base.NewBaseService("", service.Client()),
	})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}