package dao

import (
	"Purejuice/model"
	"Purejuice/utils"
)

//在数据库中查询一条数据
func CheckUserNameAndPassword(username, password string) bool {
	var u model.User
	err := utils.Db.Where("username = ? AND password = ?",username,password).Find(&u)
	if err.Error != nil {
		return false
	}else {
		return true
	}
}

func CheckUserName(username string) bool {
	var u model.User
	err := utils.Db.Where("username = ? ",username).Find(&u)
	if err.Error != nil {
		return false
	}else {
		return true
	}
}
func CheckInvitationInMysql(Invitation string) bool {
	var i model.InvitationCode
	err := utils.Db.Where("code = ?",Invitation).Find(&i)
	if err.Error != nil {
		return false
	}else {
		return true
	}
}

func SaveUser(username, password string) bool {
	var u model.User
	u.Username = username
	u.Password = password
	err := utils.Db.Create(&u)
	if err.Error != nil {
		return false
	}else {
		return true
	}
}

func DeleteInvitationCode(In string) {
	var i model.InvitationCode
	utils.Db.Where("code = ?",In).Find(&i)
	//utils.Db.Where(model.InvitationCode{Code: In}).FirstOrInit(&i)
	utils.Db.Where("ID = ?",i.ID).Delete(&i)
}