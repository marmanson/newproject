package models

import(
	"newproject/utils"
	)

type Subscribe struct{
	ServiceName    string
	ServiceRoute   string
	UserName       string
	Token          string
	DeadLine       int64
}


type UserSubscribe struct{
	ServiceName    string
	ServiceRoute   string
	Token          string
}


func InsertSubscribe(sub Subscribe)error{
	_,err := utils.ModifyDB("insert subscribe set servicename=?,serviceroute=?,username=?,token=?,deadline=?",sub.ServiceName,sub.ServiceRoute,sub.UserName,sub.Token,sub.DeadLine)
	if err != nil{
		return err
	}
	return nil
}

func QuerySubscribe(username string)([]UserSubscribe,error){
	rows ,errs := utils.Db.Query("select servicename,serviceroute,token from subscribe where username =?",username)
	if errs != nil{
		panic(errs)
	}
	U := []UserSubscribe{}
	for rows.Next(){
		var sn  string
		var sr  string
		var tk  string
		rows.Scan(&sn,&sr,&tk)
		u := UserSubscribe{sn,sr,tk}
		U = append(U,u)
	}
	return U,nil
}
func QueryUserExist(username string,serviceroute string)bool{
	rows := utils.Db.QueryRow("select username from subscribe where serviceroute =?",serviceroute)
	var u string
	rows.Scan(&u)
	if u == username{
		return true
	}else{
		return false
	}
}
func QueryToken(serviceroute string)string{
	rows := utils.Db.QueryRow("select token from subscribe where serviceroute =?",serviceroute)
	var tk string
	rows.Scan(&tk)
	return tk
}
func UpdateSubscribe(updatesuscribe Subscribe)error{
	_,err := utils.ModifyDB("update subscribe set token=?,deadline=? where Username=? and ServiceRoute=?",updatesuscribe.Token,updatesuscribe.DeadLine,updatesuscribe.UserName,updatesuscribe.ServiceRoute)
	if err != nil{
		return err
	}
	return nil
}

func DeleteSubscribe(user Subscribe)error{
	_,err := utils.ModifyDB("delete from subscribe where serviceroute=?",user.ServiceRoute)
	if err != nil{
		return err
	}
	return nil
}