package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/CodyGuo/godaemon"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("html/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", "")
	})
	r.Run(":80")
}
