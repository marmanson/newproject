package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["newproject/controllers:LoginController"] = append(beego.GlobalControllerRouter["newproject/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["newproject/controllers:RegisterController"] = append(beego.GlobalControllerRouter["newproject/controllers:RegisterController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
