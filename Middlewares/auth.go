package middlewares

import (
	"fmt"
	"mestorage/security"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

		if idVal, ok := result["_id"].(string); ok {
			ctx.Set("user", idVal)
			ctx.Next()
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "empty",
			})
			return
		}
	}
}
