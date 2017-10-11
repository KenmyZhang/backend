package controllers

import (
    "github.com/KenmyZhang/mini-chat/models"
    "net/http"
    "encoding/json"
)

type PostController struct {
	MainController	
}

func (this *PostController) Prepare() {
	//this.StopRun()
}

// @Title create post
// @Description create post
// @Success 200 {Post}  models.Post
// @Param   channel_id   body   string true       "message"
// @Param   message      body   string true       "密码"
// @Failure 400 no enough input
// @Failure 500 get products common error
// @router /create [post]
func (c *PostController) CreatePost() {
    ob := &models.Post{}
    json.Unmarshal(c.Ctx.Input.RequestBody, ob)

    if len(ob.Message) == 0 {
    	models.SetInvalidParam(c.Ctx, "Post.message", http.StatusBadRequest)
    	return
    }

    if len(ob.ChannelId) == 0 {
    	models.SetInvalidParam(c.Ctx, "Post.ChannelId", http.StatusBadRequest)
    	return
    }

    ob.UserId = c.GetSession("user_id").(string)

	if post, err := models.CreatePost(ob); err != nil {
		models.AppError(c.Ctx, err.Error(), http.StatusUnauthorized)
		return
	} else {
        c.Ctx.WriteString(post.ToJson())
    }
}

// @Title create post
// @Description create post
// @Success 200 {Post}  models.Post
// @Param   message    body   string true       "message"
// @Param   password    body   string true       "密码"
// @Failure 400 no enough input
// @Failure 500 get products common error
// @router /:channel_id([A-Za-z0-9]+)/posts [get]
func (c *PostController) GetPost() {
    channelId := c.Ctx.Input.Param(":channel_id")
    page, _ := c.GetInt("page")
    perpage, _ := c.GetInt("perpage")

	if posts, err := models.GetPost(channelId, page, perpage); err != nil {
		models.AppError(c.Ctx, err.Error(), http.StatusInternalServerError)
		return
	} else {
        c.Ctx.WriteString(models.PostsToJson(posts))
    }
}