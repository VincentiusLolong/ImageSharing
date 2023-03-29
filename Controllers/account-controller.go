package controllers

import (
	"context"
	"fmt"
	"mestorage/models"
	"mestorage/security"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (c *controller) CreateAccount(ctx *gin.Context) {
	var accounts models.Account
	a, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	v := &accounts.Account_Created
	*v = time.Now()
	err := ctx.ShouldBindJSON(&accounts)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	security.HashPassword(&accounts, accounts.Password)

	res, errs := c.service.CreateAccount(a, accounts)
	if errs != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errs.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": res,
		})
		return
	}
}

func (c *controller) SignIn(ctx *gin.Context) { //
	a, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var signin models.SignIn
	err := ctx.ShouldBindJSON(&signin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	response, errs := c.service.LoginAccount(a, signin)
	if errs != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	if passerr := security.CheckPassword(response, signin.Password); passerr != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": passerr.Error(),
		})
		return
	}

	t, jwterr := security.GenerateJWT(&models.UserToken{
		Id:       response.Id,
		Email:    response.Email,
		Username: response.Username,
	}, 15)
	if jwterr != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": jwterr.Error(),
		})
		return
	}

	ts, refresherr := security.GenerateJWT(&models.RefreshDataToken{
		Id: response.Id,
	}, 43200)
	if refresherr != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": jwterr.Error(),
		})
		return
	}

	// tokens := models.Token{
	// 	Refreshtoken: ts,
	// }
	rediserr := c.redis.Set(response.Id.Hex(), ts, 2592000)
	if rediserr != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": rediserr.Error(),
		})
		return
	} else {
		session := sessions.Default(ctx)
		session.Set("usertoken", t)
		session.Save()
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Login Success",
		})
		return
	}
}

func (c *controller) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	key := session.Get("usertoken")
	if key == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "no session",
		})
		return
	}
	authToken := fmt.Sprintf("%v", key)

	result, err := security.ValidateToken(authToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	idVal, ok := result["_id"].(string)
	if !ok {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "interface {} is nil",
		})
		return
	} else {
		session.Clear() // this will mark the session as "written" only if there's
		// at least one key to delete
		session.Options(sessions.Options{MaxAge: -1})
		session.Save()
		redisdel, err := c.redis.Delete(idVal)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": redisdel,
			})
			return
		}
	}

}
