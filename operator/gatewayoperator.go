package operator

import (
	"math/rand"
	"time"
	"github.com/gorequest"
	"github.com/kong-sdk"
	"newproject/models"
	"fmt"
)

const endpoint = "http://localhost:8001"
const repoint = "http://localhost:8000/"
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func Register(ser UserSubscribeService)error{
	random := GetRandomString(20)
	//fmt.Println(ser.Servicename)
	name := random + ser.Servicename
	service := kong_sdk.ServiceRequest{name,ser.Serviceurl,ser.Servicepath}
	//fmt.Println(name,ser.Serviceurl,ser.Servicepath)
	s,err :=kong_sdk.CreateService(service,endpoint)
	if err != nil{
		kong_sdk.DeleteService(endpoint,s.Name)
		fmt.Println(err)
		return err

	}
	random = GetRandomString(20)
	h := []string{random+s.Name}
	route := kong_sdk.RouteRequest{random+s.Name,h}
	r ,err1 := kong_sdk.CreateRoute(route,endpoint,s.Name)
	if err1 != nil{
		kong_sdk.DeleteRoute(endpoint,r.Id)
		kong_sdk.DeleteService(endpoint,s.Name)
		return err1
	}
	random = GetRandomString(20)
	subscirbe := models.Subscribe{s.Name,r.Name,ser.Username,random,time.Now().Unix()}
	err2 := models.InsertSubscribe(subscirbe)
	if err2 != nil{
		kong_sdk.DeleteRoute(endpoint,r.Id)
		kong_sdk.DeleteService(endpoint,s.Name)
		return err2
	}
	return nil
}

func DeleteServiceandRoute(sub models.Subscribe)string{
	res := models.QueryUserExist(sub.UserName,sub.ServiceRoute)
	tk := models.QueryToken(sub.ServiceRoute)
	if res && (tk == sub.Token) {
		if !K8sDeleteService(){
			return "无法删除该服务"
		}
		err := kong_sdk.DeleteRoute(endpoint, sub.ServiceRoute)
		if err != nil {
			return "无法删除该服务"
		}
		err1 := kong_sdk.DeleteService(endpoint, sub.ServiceName)
		if err1 != nil {
			return "无法删除该服务"
		}
		err2 := models.DeleteSubscribe(sub)
		if err2 != nil {
			return "无法删除该服务"
		}
	}else{
		return "无法删除该服务"
	}
	return "删除成功"
}

func RequestService(username string, serviceroute string,token string,param string)(string){
	res := models.QueryUserExist(username,serviceroute)
	tk := models.QueryToken(serviceroute)
	if res && (tk == token){
		r, body, err := gorequest.New().Post(repoint).AppendHeader("Host", serviceroute).End()
		if err != nil {
			return "无法获取该服务"
		}

		if r.StatusCode == 401 || r.StatusCode == 403 {
			return "无权请求该服务"
		}
		return body
	}else{
		return "无权请求该服务"
	}
}
