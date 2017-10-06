package controllers

import (
	"github.com/astaxie/beego/context"
)

var FilterUser = func(ctx *context.Context) {
    _, ok := ctx.Input.Session("uid").(int)
    if !ok {
        ctx.Redirect(302, "/v1/user/login")
    }
}