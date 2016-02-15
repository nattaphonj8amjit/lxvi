package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// init
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.LoadHTMLGlob("templates/**/*")


	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "LXVI",
		})
	})

	router.Run(":8080")

}

