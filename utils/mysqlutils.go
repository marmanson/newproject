package utils

import (
	"database/sql"
	"crypto/md5"
	"log"
	"fmt"
	_ "github.com/mysql"
	)

var Db *sql.DB

//type DB struct{
//	Drivername  string
//	User        string
//	Password    string
//	Host        string
//	DbName      string
//}
const Drivername      string = "mysql"
const DataSourceName  string = "root:12300@/cloudplatform"
func init(){
		if Db == nil {
			Db, _ = sql.Open(Drivername,DataSourceName)
			CreateTableWithUser()
			CreateTableWithService()
			CreateTableWithSubscribe()
		}

}
//创建用户表
func CreateTableWithUser(){
	sql := `CREATE TABLE IF NOT EXISTS users(
		username VARCHAR(20) PRIMARY KEY NOT NULL,
		password VARCHAR(64) NOT NULL,
		role INT(2) NOT NULL
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	ModifyDB(sql)
}
//创建服务表
func CreateTableWithService(){
	sql := `CREATE TABLE IF NOT EXISTS services(
		servicename  VARCHAR(50) PRIMARY KEY NOT NULL,
		path         VARCHAR(50) NOT NULL,
		statement    VARCHAR(200)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	ModifyDB(sql)
}
//创建kubeflow表



//创建用户订阅表
func CreateTableWithSubscribe(){
	sql := `CREATE TABLE IF NOT EXISTS subscribe(
		servicename VARCHAR(50) NOT NULL,
		serviceroute VARCHAR(50) PRIMARY KEY NOT NULL,
		username  VARCHAR(20) NOT NULL,
		token     VARCHAR(20) NOT NULL,
		deadline  INT(10),
		CONSTRAINT username FOREIGN KEY(username) REFERENCES users(username)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	ModifyDB(sql)
}
//操作数据库
func ModifyDB(sql1 string, args ...interface{})(int64,error){
	result,err := Db.Exec(sql1,args...)
	if err != nil{
		log.Println(err)
		return 0,err
	}
	count ,err := result.RowsAffected()
	if err != nil{
		log.Println(err)
		return 0,err
	}
	return count,nil
}

func MD5(str string) string{
	md5str := fmt.Sprintf("%x",md5.Sum([]byte(str)))
	return md5str
}
