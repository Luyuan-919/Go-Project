package controllers

import (
	"Purejuice/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckInvitation(c *gin.Context) {
	Invitation := c.PostForm("InvitationCode")
	if ok := dao.CheckInvitationInMysql(Invitation); !ok {
		c.JSON(http.StatusOK,"此邀请码不存在！")
		//c.HTML(http.StatusOK,"registered.html","此邀请码不存在！")
		}else {
		c.JSON(http.StatusOK,"")
		//c.HTML(http.StatusOK,"registered.html","邀请码可用！")
	}
}

func CheckUsername(c *gin.Context) {
	username := c.PostForm("username")
	if ok := dao.CheckUserName(username); ok {
		c.JSON(http.StatusOK,"此账号已经存在！")
		//c.HTML(http.StatusOK,"registered.html","此邀请码不存在！")
	}else {
		c.JSON(http.StatusOK,"")
		//c.HTML(http.StatusOK,"registered.html","邀请码可用！")
	}
}