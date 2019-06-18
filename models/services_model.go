package models

import (
	"newproject/utils"
	)

type  Service struct{
	Servicename        string
	Path               string
	Statement          string
}

func InsertService(ser  Service)error{
	stmt,errs := utils.Db.Prepare("INSERT services SET servicename =?,path =?,statement =?")
	if errs != nil{
		return errs
	}
	defer stmt.Close()
	_, err1 := stmt.Exec(ser.Servicename,ser.Path,ser.Statement)
	if err1 != nil{
		return err1
	}
	return nil
}


func QueryServiceExist(servicename string)bool{
	row:= utils.Db.QueryRow("select * from services where servicename =?",servicename)
	s := Service{}
	row.Scan(&s.Servicename,&s.Path,&s.Statement)
	if s.Servicename != ""{
		return true
	}
	return false
}


func QueryService()([]Service,error){
	S := []Service{}
	row ,err1 := utils.Db.Query("select * from services")
	if err1 != nil{
		return nil,err1
	}

	for row.Next(){
		s := Service{}
		row.Scan(&s.Servicename,&s.Path,&s.Statement)
		S = append(S,s)
	}
	return S,nil
}

func DeleteService(servicename string)error{
	_,err := utils.ModifyDB("delete from services where servicename =?",servicename)
	if err != nil{
		return err
	}
	return nil
}

func UpdateService(servicename string,ser Service)error{
	_,err := utils.ModifyDB("update services set servicename =?,path =?,statement =? where servicename =?",ser.Servicename,ser.Path,ser.Statement,servicename)
	if err != nil{
		return err
	}
	return nil
}