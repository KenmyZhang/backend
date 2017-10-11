package controllers

import (
    "github.com/KenmyZhang/mini-chat/models"
    "encoding/json"
    "net/http"
)

type ChannelController struct {
	MainController	
}


// @Title CreateDirectChannel
// @Description CreateDirectChannel
// @Success 200 {object} models.Channel the details of Channel
// @Param   userIds  body   []string true       "userIds"
// @Failure 400 no enough input
// @Failure 500 get products common error
// @router /create [post]
func (c *ChannelController) CreateDirectChannel() {
	channelUserIds :=  &models.ChannelUserIds{}
	json.Unmarshal(c.Ctx.Input.RequestBody, channelUserIds)
	userIds := channelUserIds.UserIds
	if len(userIds) != 1 {
		models.SetInvalidParam(c.Ctx, "user_ids", http.StatusBadRequest)
		return
	}	
	allowed := false
	userId := c.GetSession("user_id")

	if !models.SessionHasPermissionTo(c.Ctx, models.PERMISSION_CREATE_DIRECT_CHANNEL) {
		models.SetPermissionError(c.Ctx, models.PERMISSION_CREATE_DIRECT_CHANNEL, http.StatusForbidden)
		return
	}

	if !allowed && !models.SessionHasPermissionTo(c.Ctx, models.PERMISSION_MANAGE_SYSTEM) {
		models.SetPermissionError(c.Ctx, models.PERMISSION_MANAGE_SYSTEM, http.StatusForbidden)
		return
	}

	if directChaTonnel, err := models.CreateDirectChannel(userId.(string), userIds[0]); err != nil {
		models.AppError(c.Ctx, err.Error(), http.StatusBadRequest)
	} else {
    	c.Ctx.WriteString(directChaTonnel.ToJson())
    	return
    }
}