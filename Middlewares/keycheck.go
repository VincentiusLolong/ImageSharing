package middlewares

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var refreshMutex sync.Mutex

func isLoggingin(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	key := session.Get("usertoken")
	return key != nil
}

// func RefreshToken() func() {
// 	return func() {
// 		refreshTicker := time.Tick(1 * time.Second)
// 		for range refreshTicker {
// 			// Call your function here
// 			fmt.Println("test")
// 		}
// 	}
// }

// k > i > not > next /1
// k > i > yes > built-in go func   /2 != deadlock!
// k > i > yes > built-in go > func  /3 > func > func > func
// handlerfunc  2 func > top func

func KeyMiddelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshCh := make(chan bool)
		// refreshMutex.Lock()
		// defer refreshMutex.Unlock()

		// if isRefreshing, ok := c.Get("isRefreshing"); ok && !isRefreshing.(bool) {
		// 	c.Next()
		// 	return}
		if !isLoggingin(c) {
			refreshCh <- true
			c.Next()
			return
		} else {
			go func() {
				refreshMutex.Lock()
				defer refreshMutex.Unlock()
				// go func(stop chan bool, c *gin.Context) {
				// c.Set("isRefreshing", true)
				// Perform the refresh operation here...
				refreshTicker := time.Tick(1 * time.Minute)
				// for range refreshTicker {
				// 	if isLoggingin(c) {
				// 		fmt.Println("test")
				// 	} else {
				// 		stop <- true
				// 		return
				// 	}
				// }
			refreshLoop:
				for {
					select {
					case <-refreshTicker:
						fmt.Println("test")
					case <-refreshCh:
						break refreshLoop
					}
				}
				// for {
				// 	select {
				// 	case <-refreshTicker:
				// 		// Perform the refresh operation here...
				// 		fmt.Println("test")

				// 	case <-c.Done():
				// 		close(refreshCh)
				// 		return
				// 	}
				// }
				// }(refreshCh, c)
				<-refreshCh
				// c.Set("isRefreshing", false)
			}()
			// go RefreshToken()()
			// <-refreshCh
			c.Next()
			return
		}

	}
}

// func TestRefresh() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if !isLoggingin(c) {
// 			c.Next()
// 			return
// 		} else {
// 			header := isLoggingin(c)
// 			done := make(chan bool)
// 			go func(headers bool, dones chan bool) {
// 				refreshMutex.Lock()
// 				defer refreshMutex.Unlock()
// 				refreshTicker := time.Tick(1 * time.Minute)
// 				for {
// 					select {
// 					case <-refreshTicker:
// 						if !headers {
// 							done <- true
// 							return
// 						} else {
// 							fmt.Println("test")
// 						}
// 						// Perform the refresh operation here...

// 					case <-done:
// 						return
// 					}
// 				}
// 			}(header, done)
// 			done <- true
// 		}
// 	}
// }

// return func() {
// 	time.Sleep(1 * time.Minute)
// 	refreshTicker := time.Tick(1 * time.Minute)
// 	for {
// 		select {
// 		case <-refreshTicker:
// 			fmt.Println("test")
// 			// var idtoken models.SessionData
// 			// session := sessions.Default(ctx)
// 			// key := session.Get("usertoken")
// 			// byteData, ok := key.([]byte)
// 			// if !ok {
// 			// 	ctx.JSON(http.StatusConflict, gin.H{
// 			// 		"message": "Confilict in Json Token Convert",
// 			// 	})
// 			// 	return
// 			// }
// 			// err := json.Unmarshal([]byte(byteData), &idtoken)
// 			// if err != nil {
// 			// 	ctx.JSON(http.StatusConflict, gin.H{
// 			// 		"message": err.Error(),
// 			// 	})
// 			// 	return
// 			// }
// 			// // result, err := security.ValidateToken(idtoken.AccessToken)
// 			// // if err != nil {
// 			// // 	ctx.JSON(http.StatusBadRequest, gin.H{
// 			// // 		"message": err.Error(),
// 			// // 	})
// 			// // 	return
// 			// // }
// 			// reinsight, rediserr := redisinsight.Get(idtoken.Id)
// 			// if rediserr != nil {
// 			// 	ctx.JSON(http.StatusBadRequest, gin.H{
// 			// 		"message": err.Error(),
// 			// 	})
// 			// 	return
// 			// }
// 			// var redisrefreshtoken string
// 			// converterr := json.Unmarshal([]byte(reinsight), &redisrefreshtoken)
// 			// if err != nil {
// 			// 	ctx.JSON(http.StatusConflict, gin.H{
// 			// 		"message": converterr.Error(),
// 			// 	})
// 			// 	return
// 			// }
// 			// result, err := security.ValidateToken(redisrefreshtoken)
// 			// if err != nil {
// 			// 	ctx.JSON(http.StatusBadRequest, gin.H{
// 			// 		"message": err.Error(),
// 			// 	})
// 			// 	return
// 			// }
// 		}
// 	}
// }
