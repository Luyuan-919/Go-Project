package main

import (
	"Purejuice/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	r := gin.Default()
	r.Static("css","static/css")
	r.Static("script","static/script")
	r.Static("img","static/css/img")
	//r.LoadHTMLFiles("static/login.html","static/pages/loginSuccessful.html")
	//r.LoadHTMLFiles("static/login.html")
	r.LoadHTMLGlob("static/pages/*")
	//去首页
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",gin.H{})
	})
	r.POST("/login",controllers.Login)
	r.POST("/registered",controllers.Registered)
	r.POST("/registeredHtml", func(c *gin.Context) {
		c.HTML(http.StatusOK,"registered.html",gin.H{})
	})
	r.POST("/checkInvitation",controllers.CheckInvitation)
	r.POST("/checkUserNameExist",controllers.CheckUsername)
	r.POST("/upload",controllers.UpLoad)
	r.Run(":8080")
}
