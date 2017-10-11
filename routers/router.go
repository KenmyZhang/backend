// @APIVersion 1.0.0
// @Title mini-chat API
// @Description  IM.
// @Contact 1027837952@qq.com
// @TermsOfServiceUrl
// @License
// @LicenseUrl
package routers

import (
	"github.com/KenmyZhang/mini-chat/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	ns :=
		beego.NewNamespace("/v1",
        	
        	//此处正式版时改为验证加密请求
        	beego.NSCond(func(ctx *context.Context) bool {
            	if ua := ctx.Input.UserAgent(); ua != "" {
                	return true
            	}
            	return false
        	}),
        	
        	beego.NSNamespace("/user",
            	beego.NSInclude(
                    &controllers.UserController{},
                ),
            ),
                	
            beego.NSNamespace("/post",
                beego.NSInclude(
                    &controllers.PostController{},
                ),
            ),

            beego.NSNamespace("/channel",
                beego.NSInclude(
                    &controllers.ChannelController{},
                ),
            ),

            beego.NSNamespace("/file",
                beego.NSInclude(
                    &controllers.FileController{},
                ),
            ), 
    	)

	beego.AddNamespace(ns)

    beego.InsertFilter("/user/:id([0-9]+)",beego.BeforeRouter,controllers.FilterUser)

    beego.ErrorController(&controllers.ErrorController{})
}
 