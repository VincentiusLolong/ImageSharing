package main

import (
	"fmt"
	middlewares "mestorage/Middlewares"
	"mestorage/conf"
	"mestorage/conf/database"
	"mestorage/router"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	conf.Init()
	database.MongoConnect()
	redis := database.Setup()
	if redis != nil {
		panic(fmt.Errorf("fatal error : %s", redis.Error()))
	}
	server.Use(middlewares.SessionMiddleware())
	router.Routers(server)
	// go func() {
	// 	refreshTicker := time.Tick(14 * time.Second)
	// 	for range refreshTicker {
	// 		// Call your function here
	// 		fmt.Println("test")
	// 	}
	// }()
	gin.SetMode(gin.ReleaseMode)
	server.Run(":8080")
}
