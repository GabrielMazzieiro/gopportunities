package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	initializeRoutes(router)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "SERVER OK",
		})
	})

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "pong",
		})
	})
	router.Run(":8080")
}
