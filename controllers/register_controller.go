package controllers

import (
	"github.com/astaxie/beego"
	"newproject/models"
	"newproject/utils"
	)

type RegisterController struct{
	beego.Controller
}

func(this *RegisterController) URLMapping(){
	this.Mapping("Get",this.Get)
	this.Mapping("Post",this.Post)
}

func(this *RegisterController) Get(){

}
//Post ...
//@Title Post
//@Param   username  query  string true  "username"
//@Param   password  query  string true  "password"
//@Param   repassword query  string true "repassword"
// @Success 200 {object} models.User
// @Failure 403
// @router / [post]
func(this *RegisterController) Post(){
	username := this.Ctx.Input.Query("username")
	password := this.Ctx.Input.Query("password")
	repassword := this.Ctx.Input.Query("repassword")
	 res:= models.QueryUser(username)
	if password != repassword{
		this.Data["json"] = map[string]interface{}{"code":0,"message":"两次密码不一致"}
		this.ServeJSON()
		return
	}
	if res {
		this.Data["json"] = map[string]interface{}{"code":0,"message":"用户名已经存在"}
		this.ServeJSON()
		return
	}

	password = utils.MD5(password)
	user := models.User{0,username,password,1}
	err := models.InsertUser(user)

	if err != nil{
		this.Data["json"] = map[string]interface{}{"code":0,"message":"注册失败"}
		this.ServeJSON()
	}else{
		this.Data["json"] = map[string]interface{}{"code":1,"message":"注册成功"}
		this.ServeJSON()
	}
}