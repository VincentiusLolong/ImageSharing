package router

import (
	controllers "mestorage/Controllers"
	middlewares "mestorage/Middlewares"
	"mestorage/conf/database"
	"mestorage/services"

	"github.com/gin-gonic/gin"
)

var (
	db         database.Dbs               = database.New()
	redis      database.Redis             = database.NewRedis()
	serv       services.PrivateServices   = services.PrivateSerivce(db)
	controller controllers.FileController = controllers.New(serv, redis)
)

func Routers(app *gin.Engine) {
	privateRoute := app.Group("/myapp")
	privateRoute.POST("/register", controller.CreateAccount)
	privateRoute.POST("/log-in", controller.SignIn)
	privateRoute.POST("/logout", controller.Logout)
	userGroup := privateRoute.Group("/user").Use(middlewares.Auth())
	userGroup.GET("/test", controller.Test)
	// app.POST("/images", controller.Add)
	// app.GET("/images/:title", controller.Find)

}
