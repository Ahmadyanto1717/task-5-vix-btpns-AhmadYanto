package main

import (
	"task-5-vix-btpns-Ahmad_Yanto/controllers"
	"task-5-vix-btpns-Ahmad_Yanto/router"
)

func main() {
	server := controllers.Server{}
	server.Initialize()
	router.InitRoutes(&server)
	server.Run(8000)
}
