package controllers

import (
	"context"
	"mestorage/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func contextGin(res string)

func (c *controller) Add(ctx *gin.Context) {
	accountid := ctx.Param("accountid")
	accountids, _ := primitive.ObjectIDFromHex(accountid)
	var addfile models.ImagePost
	a, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := ctx.ShouldBindJSON(&addfile)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}

	reslut, err := c.service.Add(a, addfile, accountids)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": reslut,
		})
	}
}

func (c *controller) Find(ctx *gin.Context) {
	accountid := ctx.Param("accountid")
	accountids, _ := primitive.ObjectIDFromHex(accountid)
	a, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	reslut, err := c.service.Get(a, accountids)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": reslut,
		})
	}
}
