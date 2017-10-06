package main

import (
    _ "github.com/astaxie/beego/session/mysql"
	_ "github.com/KenmyZhang/mini-chat/routers"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/KenmyZhang/mini-chat/config"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"time"
)

func init() {
    maxIdle := 30
	maxConn := 30
    orm.RegisterDriver(beego.BConfig.WebConfig.Session.SessionProvider, orm.DRMySQL)	
 	orm.RegisterDataBase("default", beego.BConfig.WebConfig.Session.SessionProvider, beego.BConfig.WebConfig.Session.SessionProviderConfig, maxIdle, maxConn)
    orm.DefaultTimeLoc = time.UTC
    orm.RunSyncdb("default", false, true)  //自动建表
    orm.Debug = true //[ORM] - 时间 - [Queries/数据库名] - [执行操作/执行时间] - [SQL语句] - 使用标点 `,` 分隔的参数列表 - 打印遇到的错误
}

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.SetLogger("file", `{"filename":"logs/test.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	beego.SetLogFuncCall(true)
	if beego.BConfig.RunMode == "dev" {
		beego.SetLevel(beego.LevelDebug)
	} else {
		beego.SetLevel(beego.LevelInformational)
	}	

	beego.SetStaticPath("/down1", "download1")
	beego.Run()
}





