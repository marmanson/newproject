package controllers

import (
	"github.com/astaxie/beego"
	"newproject/models"
	"newproject/utils"
)

type LoginController struct{
	beego.Controller
}
const C_SESSION = "loginuser"
func(this *LoginController) URLMapping(){
	this.Mapping("Get",this.Get)
	this.Mapping("Post",this.Post)
}

func (this *LoginController) Get(){

}
//Post ...
//@Title Post
//@Param   username  query  string true  "username"
//@Param   password  query  string true  "password"
// @Success 200 {object} models.User
// @Failure 403
// @router / [post]
func (this *LoginController) Post(){
	username := this.Ctx.Input.Query("username")
	password := this.Ctx.Input.Query("password")
	res := models.QueryUserParam(username,utils.MD5(password))
	//fmt.Println(id)
	if res {
		this.SetSession(C_SESSION,username)
		this.Data["json"] = map[string]interface{}{"code":1,"message":"登录成功"}
	}else{
		this.Data["json"] = map[string]interface{}{"code":0,"message":"登录失败"}
	}
	this.ServeJSON()
}