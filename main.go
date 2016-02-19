package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// init
	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Static("/js", "./js")
	router.Static("/css", "./css")
	router.LoadHTMLGlob("html/*")

	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Comming Soon",
		})
	})

	router.Run(":80")

}
