package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/KenmyZhang/mini-chat/controllers:ChannelController"] = append(beego.GlobalControllerRouter["github.com/KenmyZhang/mini-chat/controllers:ChannelController"],
		beego.ControllerComments{
			Method: "CreateDirectChannel",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/KenmyZhang/mini-chat/controllers:FileController"] = append(beego.GlobalControllerRouter["github.com/KenmyZhang/mini-chat/controllers:FileController"],
		beego.ControllerComments{
			Method: "UploadFile",
			Router: `/upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/KenmyZhang/mini-chat/controllers:OauthController"] = append(beego.GlobalControllerRouter["github.com/KenmyZhang/mini-chat/controllers:OauthController"],
		beego.ControllerComments{
			Method: "LoginWithOauth",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/KenmyZhang/mini-chat/controllers:PostController"] = append(beego.GlobalControllerRouter["github.com/KenmyZhang/mini-chat/controllers:PostController"],
		beego.ControllerComments{
			Method: "GetPost",
			Router: `/:channel_id([A-Za-z0-9]+)/posts`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/KenmyZhang/mini-chat/controllers:PostController"] = append(beego.GlobalControllerRouter["github.com/KenmyZhang/mini-chat/controllers:PostController"],
		beego.ControllerComments{
			Method: "CreatePost",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/KenmyZhang/mini-chat/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/KenmyZhang/mini-chat/controllers:UserController"],
		beego.ControllerComments{
			Method: "CreateUser",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/KenmyZhang/mini-chat/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/KenmyZhang/mini-chat/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/KenmyZhang/mini-chat/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/KenmyZhang/mini-chat/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
