package controllers

import (
    "github.com/astaxie/beego"
    "encoding/json"
    "net/http"
    "github.com/KenmyZhang/mini-chat/models"
    "github.com/mssola/user_agent"    
    "fmt"
)

// USER API
type UserController struct {
	MainController	
}

// @Title create user
// @Description create user by email
// @Success 200 {object} models.User the details of User
// @Param   username  body   string true       "username"
// @Param   password  body   string true       "password"
// @Param   email     body   string true       "email"
// @Failure 400 no enough input
// @Failure 500 get products common error
// @router /create [post]
func (c *UserController) CreateUser() {
	ob := &models.User{}
    json.Unmarshal(c.Ctx.Input.RequestBody, ob)

	if ruser, err := models.CreateUser(ob); err != nil {
		models.AppError(c.Ctx, err.Error(), http.StatusBadRequest)
	} else {
    	c.Ctx.WriteString(models.UserToJson(ruser))
    	return
    }
}

// @Title login
// @Description login by username and password
// @Success 200 {object}  models.User the details of User
// @Param   username    body   string true       "用户名"
// @Param   password    body   string true       "密码"
// @Failure 400 no enough input
// @Failure 500 get products common error
// @router /login [post]
func (c *UserController) Login() {
    ob := &models.User{}
    json.Unmarshal(c.Ctx.Input.RequestBody, ob)

    if len(ob.Password) == 0 {
    	models.SetInvalidParam(c.Ctx, "password", http.StatusBadRequest)
    	return
    }

    if len(ob.Username) == 0 {
    	models.SetInvalidParam(c.Ctx, "username", http.StatusBadRequest)
    	return
    }    

	ua := user_agent.New(c.Ctx.Input.UserAgent())

	plat := ua.Platform()
	if plat == "" {
		plat = "unknown"
	}

	os := ua.OS()
	if os == "" {
		os = "unknown"
	}

	bname, bversion := ua.Browser()

	if bname == "" {
		bname = "unknown"
	}

	if bversion == "" {
		bversion = "0.0"
	}

	c.SetSession(models.SESSION_PROP_PLATFORM, plat)
	c.SetSession(models.SESSION_PROP_OS, os)
	c.SetSession(models.SESSION_PROP_BROWSER, fmt.Sprintf("%v/%v", bname, bversion))

	if user, err := models.AuthenticateUserForLogin(ob.Username, ob.Password); err != nil {
		models.AppError(c.Ctx, err.Error(), http.StatusUnauthorized)
		return
	} else {
		c.SetSession("user_id",user.Id)
		c.SetSession("roles",user.Roles)
        c.Ctx.WriteString(models.UserToJson(user))
    }
}
// @Title logout
// @Description logout
// @Success 200 success
// @Failure 400 no enough input
// @Failure 500 get products common error
// @router /logout [post]
func (c *UserController) Logout() {
	c.DestroySession()

}

func (c *UserController) test() {

  c.ServeXML()
}

func (c *UserController) Sendsms() {

}

func (c *UserController) Create() {

}

func (c *UserController) GetUsersByIds() {

}

func (c *UserController) GetUsersByUsernames() {

}

func (c *UserController) GetUserByUserId() {

}

func (c *UserController) GetImage() {

}

func (c *UserController) SetImage() {

}

func (c *UserController) ResetPassword() {

}

func (c *UserController) GetUserByUsername() {
	beego.Debug(c.Ctx.Input.Param(":username"))
}