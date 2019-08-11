package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func main() {
	staticServe()
}

func staticServe() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	router.NoRoute(func(c *gin.Context) {
		c.File("./client/build/index.html")
	})
	fmt.Println(router)

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	api2 := router.Group("/api2")
	{
		api2.GET("/", func(c *gin.Context) {
			nums := 10000
			var x = make([]float64, nums)
			for i := 0; i < nums; i++ {
				s := rand.Float64()
				x[i] += s
			}
			c.JSON(http.StatusOK, gin.H{
				"message": x,
			})
		})
	}

	_ = router.Run(":5000")
}
