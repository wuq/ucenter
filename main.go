package main

import (
//	"context"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"tenno.ucenter/handler"
//	"tenno.ucenter/subscriber"

	auth "tenno.ucenter/proto/auth"
	base "tenno.ucenter/proto/base"
	ext  "tenno.ucenter/proto/ext"
)


func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.tenno.ucenter"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	auth.RegisterAuthHandler(service.Server(), new(handler.Auth))
	base.RegisterBaseHandler(service.Server(), new(handler.Base))
	ext.RegisterExtHandler(service.Server(), new(handler.Ext))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
