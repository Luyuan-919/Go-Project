package controllers

import (
	"Purejuice/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	Ok := dao.CheckUserNameAndPassword(username,password)
	if Ok {
		c.HTML(http.StatusOK,"loginSuccessful.html",gin.H{})

	}else{
		c.HTML(http.StatusOK,"login.html","用户名或密码不正确！")
	}
}
func Registered(c *gin.Context) {
	Invitation := c.PostForm("InvitationCode")
	if ok := dao.CheckInvitationInMysql(Invitation); !ok {
		c.HTML(http.StatusOK,"registered.html","error1")
		return
	}
	username := c.PostForm("Username")
	if ok := dao.CheckUserName(username); ok {
		c.HTML(http.StatusOK,"registered.html","error2")
		return
	}
	password := c.PostForm("Password")
	if ok := dao.SaveUser(username, password); ok {
		c.HTML(http.StatusOK,"registered.html","success")
	}

	dao.DeleteInvitationCode(Invitation)
}
