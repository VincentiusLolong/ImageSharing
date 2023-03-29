package controllers

import (
	"mestorage/conf/database"
	"mestorage/services"

	"github.com/gin-gonic/gin"
)

type FileController interface {
	Add(ctx *gin.Context)
	Find(ctx *gin.Context)
	CreateAccount(ctx *gin.Context)
	SignIn(ctx *gin.Context)
	Test(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Refresh(ctx *gin.Context)
}

type controller struct {
	service services.PrivateServices
	redis   database.Redis
}

func New(srv services.PrivateServices, redis database.Redis) FileController {
	return &controller{
		service: srv,
		redis:   redis,
	}
}

func (c *controller) MainSession() {

}
