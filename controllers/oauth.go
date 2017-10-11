package controllers

import (
    "github.com/KenmyZhang/mini-chat/models"
    "github.com/KenmyZhang/mini-chat/config"
    "net/http"
)

// Oauth API
type OauthController struct {
	MainController	
}
// @Title login
// @Description login by username and password
// @Success 200  
// @Param   category_id     query   int false       "category id"
// @Param   brand_id    query   int false       "brand id"
// @Param   query   query   string  false       "query of search"
// @Param   segment query   string  false       "segment"
// @Param   sort    query   string  false       "sort option"
// @Param   dir     query   string  false       "direction asc or desc"
// @Param   offset  query   int     false       "offset"
// @Param   limit   query   int     false       "count limit"
// @Param   price           query   float       false       "price"
// @Param   special_price   query   bool        false       "whether this is special price"
// @Param   size            query   string      false       "size filter"
// @Param   color           query   string      false       "color filter"
// @Param   format          query   bool        false       "choose return format"
// @Failure 400 no enough input
// @Failure 500 get products common error
// @router /create [post]
func (c *OauthController) LoginWithOauth() {

	if c.Ctx.Input.Param(":service") != config.OauthCfg.String("service") {
		models.SetInvalidParam(c.Ctx, "service", http.StatusBadRequest)
		return
	}

	if isEnable, _ := config.OauthCfg.Bool("enable"); isEnable != true {
		models.AppError(c.Ctx, "oauth is't enable", http.StatusNotImplemented)
	}

	secure := false
	if c.Ctx.Input.Scheme() == "https" {
		secure = true
	}


  	clientId := config.OauthCfg.String("id")
  	endpoint := config.OauthCfg.String("authendpoint")
  	scope := config.OauthCfg.String("scope")

  	_=clientId
  	_=endpoint
  	_=scope
  	_=secure


    c.Ctx.WriteString("success")
}
