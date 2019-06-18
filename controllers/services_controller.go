package controllers

import (
	"newproject/models"
	"fmt"
	"newproject/operator"
)

type ServiceController struct{
	BaseController
}


func(this *ServiceController) URLMapping(){
	//this.Mapping("Get",this.Get)
	this.Mapping("Post",this.Post)
	this.Mapping("Delete",this.Delete)
	this.Mapping("Patch",this.Patch)
	this.Mapping("Request",this.Request)
}


func(this *ServiceController)Get(){
	S, err := models.QueryService()
	if err != nil {
		fmt.Println(err)
	}
	this.Data["json"] = []models.Service(S)
	this.ServeJSON()
}
//Post ...
//@Title Post
//@Param   servicename  query  string true  "servicename"
//@Param   path  query  string true  "path"
//@Param   statement  query string true "statement"
// @Success 200 {object} models.Service
// @Failure 403
// @router / [post]
func(this *ServiceController)Post(){
	servicename := this.Ctx.Input.Query("servicename")
	path := this.Ctx.Input.Query("path")
	statement := this.Ctx.Input.Query("statement")
	res := models.QueryServiceExist(servicename)
	if res{
		this.Data["json"]=map[string]interface{}{"code":0,"message":"服务已经存在"}
	}else {
		s := models.Service{servicename,path,statement}
		err := models.InsertService(s)
		if err != nil{
			this.Data["json"]=map[string]interface{}{"code":0,"message":"添加服务失败"}
		}else{
			this.Data["json"]=map[string]interface{}{"code":1,"message":"添加服务成功"}
		}
	}
	this.ServeJSON()
}
//Delete ...
//@Title Delete
//@Param  servicename    query  string true  "servicename"
//@Success 200 {string} success
//@Failure  403
//@router / [delete]
func(this *ServiceController)Delete(){
	servicename := this.Ctx.Input.Query("servicename")
	res := models.QueryServiceExist(servicename)
	if res {
		err := models.DeleteService(servicename)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": "服务删除失败"}
		} else {
			this.Data["json"] = map[string]interface{}{"code": 1, "message": "服务删除成功"}
		}
		this.ServeJSON()
	}else{
		this.Data["json"] = map[string]interface{}{"code": 0,"message": "服务不存在，无法删除"}
		this.ServeJSON()
	}
}
//Patch ...
//@Title Patch
//@Param servicename    query   string true "servicename"
//@Param updateservicename  query  string true "updateservicename"
//@Param updatepath   query  string  true "updatepath"
//@Param updatestatement  query  string true "updatestatement"
//@Success 200  {object}  models.Service
//@Failure  403
//@router  / [patch]
func(this *ServiceController)Patch(){
	servicename := this.Ctx.Input.Query("servicename")
	updateservicename := this.Ctx.Input.Query("updateservicename")
	updatepath := this.Ctx.Input.Query("updatepath")
	updatestatement := this.Ctx.Input.Query("updatestatement")
	res := models.QueryServiceExist(servicename)
	if res {
		S := models.Service{updateservicename, updatepath, updatestatement}
		err := models.UpdateService(servicename, S)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"code": 0, "message": "更新服务失败"}
		} else {
			this.Data["json"] = map[string]interface{}{"code": 1, "message": "更新服务成功"}
		}
		this.ServeJSON()
	}else{
		this.Data["json"] = map[string]interface{}{"cpde":0,"message":"所更新服务不存在"}
		this.ServeJSON()
	}
}

//Request ...
//@Title Request
//@Param servicename    query   string true "servicename"
//@Param servicepath    query  string true "servicepath"
//@Success 200  {string}  success
//@Failure  403
//@router  / [get]
func(this *ServiceController)Request(){
	servicename := this.Ctx.Input.Query("servicename")
	servicepath := this.Ctx.Input.Query("servicepath")
	username := this.Loginuser
	fmt.Println(username)
	U,err := operator.K8sCreateService(username.(string),servicename,servicepath)
	if err != nil{
		this.Data["json"] = map[string]interface{}{"code":0,"message":"请求服务失败1"}
	}else{
		fmt.Println(U.Serviceurl)
		err1 := operator.Register(U)
		if err1 != nil{
			this.Data["json"] = map[string]interface{}{"code":0,"message":"请求服务失败2"}
		}else{
			this.Data["json"] = map[string]interface{}{"code":1,"message":"服务请求成功，请查询自己的订阅表"}
		}
	}
	this.ServeJSON()
}

