package httpserver

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func StartServer() {
	r := gin.Default()

	r.POST("/cost/long", func(c *gin.Context) {
		time.Sleep(100000 * time.Millisecond)
		c.JSON(200, gin.H{
			"message": "long post succeed!",
		})
	})

	r.POST("/cost/random", func(c *gin.Context) {
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		c.JSON(200, gin.H{
			"message": "random post succeed!",
		})
	})

	r.POST("/cost/short", func(c *gin.Context) {
		time.Sleep(10 * time.Millisecond);
		c.JSON(200, gin.H{
			"message": "short post succeed!",
		})
	})

    r.Run()
}
