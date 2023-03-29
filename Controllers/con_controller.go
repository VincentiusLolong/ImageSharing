package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *controller) Test(ctx *gin.Context) {
	value, bol := ctx.Get("user")
	if !bol {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "interface {} is nil",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": value,
		})
		return
	}

}

func (c *controller) Refresh(ctx *gin.Context) {

}
