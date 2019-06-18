package controllers

import "fmt"
type HomeController struct{
	BaseController
}
func(this *HomeController) URLMapping(){
	this.Mapping("Get",this.Get)
}
//Get ...
//@Title Get
// @Success 200 {string} success
// @Failure 403
// @router / [get]
func (this *HomeController) Get(){
	fmt.Println("Islogin:",this.IsLogin,this.Loginuser)
}