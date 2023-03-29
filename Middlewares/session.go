package middlewares

import (
	"mestorage/conf/env"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware() gin.HandlerFunc {
	store := cookie.NewStore([]byte(env.Env("SESSION_SECRET_KEY")))
	store.Options(sessions.Options{
		HttpOnly: true,
	})

	return sessions.Sessions("my-session", store)
}
