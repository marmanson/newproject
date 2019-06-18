package models

import (
	"newproject/utils"
)

type User struct{
	Id       int
	Username string
	Password string
	Role   int //用户身份
}

//插入
func InsertUser(user User)(error){
	stmt,errs := utils.Db.Prepare("INSERT users SET username =?,password =?,role =?")
	if errs != nil{
		panic(errs)
	}
	defer stmt.Close()
	_, err1 := stmt.Exec(user.Username,user.Password,user.Role)
	if err1 != nil{
		panic(err1)
	}
	return nil
}

func QueryUser(username string)(bool){
	rows := utils.Db.QueryRow("select username from users where username =?",username)
	var id string
	_ = rows.Scan(&id)
	if id == username{
		return true
	}else{
		return false
	}
}

func QueryUserParam(username string,password string)bool{
	rows:= utils.Db.QueryRow("select username from users where username =? and password =?",username,password)
	var id string
	_ = rows.Scan(&id)
	if id == username{
		return true
	}else{
		return false
	}
}

func QueryUserRole(username string)int{
	rows := utils.Db.QueryRow("select role from users where username =?",username)
	var role int
	_ = rows.Scan(&role)
	return role
}

