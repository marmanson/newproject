package controllers
type ExitController struct {
	BaseController
}

//Get ...
//@Title Get
// @Success 200 {string} success
// @Failure 403
// @router / [get]
func (this *ExitController)Get(){
	//清除该用户登录状态的数据
	this.DelSession("loginuser")
	this.Redirect("/home",302)
}
