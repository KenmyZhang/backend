package config 

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

var OauthCfg config.Configer
func init() {
	var err error
	OauthCfg, err = config.NewConfig("ini", "./conf/oauth.conf")
	if err != nil {
    	beego.Error(err)
	}
}