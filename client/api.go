package main

import (

	"github.com/micro/go-log"
	"github.com/gin-gonic/gin"
	
	"github.com/micro/go-web"

	"tenno.ucenter/client/f"
)


var err error



func main() {

	// Create service
	service := web.NewService(
		web.Name("go.micro.api.ucenter"),
	)
	service.Init()

	// Create RESTful handler (using Gin)
	router := gin.Default()

	bc := new(f.BaseClient)
	router.POST("/ucenter/getUserInfo/:uid", bc.GetUserInfo)

	ac := new(f.AuthClient)
	router.POST("/ucenter/login", ac.Login)



	// Register Handler
	//service.HandleFunc("/ucenter/get", cc)
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
