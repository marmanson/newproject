package controllers

import (
	"newproject/models"
	"fmt"
	"newproject/operator"
)
type SubscribeController struct{
	BaseController
}

func(this *SubscribeController) URLMapping(){
	this.Mapping("Get",this.Get)
	this.Mapping("Post",this.Post)
	this.Mapping("Delete",this.Delete)
	this.Mapping("Request",this.Request)
}
//Get ...
//@Title Get
// @Success 200 {string} success
// @Failure 403
// @router / [get]
func(this *SubscribeController) Get(){
	username := this.Loginuser
	fmt.Println(username)
	res,err := models.QuerySubscribe(username.(string))
	if err != nil{
		this.Data["json"] = map[string]interface{}{"code":0,"message":"获取用户订阅服务失败"}
		this.ServeJSON()
	}else{
		this.Data["json"] = []models.UserSubscribe(res)
		fmt.Println(res[0].ServiceName,res[0].ServiceRoute)
		this.ServeJSON()
	}
}
func(this *SubscribeController)Post(){

}
//Post ...
//@Title Post
//@Param  serviceroute query string true  "serviceroute"
//@Param  token query string true "token"
//@Param  param query string true  "param"
// @Success 200 {string} success
// @Failure 403
// @router / [post]
func(this *SubscribeController)Request(){
		username := this.Loginuser
		serviceroute := this.Ctx.Input.Query("serviceroute")
		token := this.Ctx.Input.Query("token")
		param := this.Ctx.Input.Query("param")
		s := operator.RequestService(username.(string),serviceroute,token,param)
		this.Data["json"] = string(s)
		this.ServeJSON()
}
//Delete ...
//@Title Delete
//@Param  serviceroute query string  true "serviceroute"
//@Param  servicename  query string  true "servicename"
//@Param  token  query string  true "token"
//@Success 200 {string} success
//@Failure 403
//@router / [delete]
func(this *SubscribeController)Delete(){
		username := this.Loginuser
		serviceroute := this.Ctx.Input.Query("serviceroute")
		servicename := this.Ctx.Input.Query("servicename")
		token := this.Ctx.Input.Query("token")
		sub := models.Subscribe{servicename,serviceroute,username.(string),token,0}
		res := operator.DeleteServiceandRoute(sub)
		this.Data["json"] = string(res)
		this.ServeJSON()
}
