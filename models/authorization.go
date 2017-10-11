package models

import (
	"strings"

	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
)

type Permission struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Role struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}

var PERMISSION_MANAGE_SYSTEM *Permission
var PERMISSION_EDIT_OTHER_USERS *Permission 
var PERMISSION_CREATE_DIRECT_CHANNEL *Permission
var PERMISSION_MANAGE_CHANNEL_MEMBERS *Permission
var PERMISSION_READ_CHANNEL *Permission


var ROLE_SYSTEM_ADMIN *Role
var ROLE_NORMAL_USER *Role 
var ROLE_CHANNEL_USER *Role
var ROLE_CHANNEL_ADMIN *Role


var BuiltInRoles map[string]*Role

func InitalPermissions() {
	PERMISSION_MANAGE_SYSTEM = &Permission{
		"manage_system",
		"permissions.manage_system.name",
		"permissions.manage_system.description",
	}
	PERMISSION_EDIT_OTHER_USERS = &Permission{
		"edit_other_users",
		"permissions.edit_other_users.name",
		"permissions.edit_other_users.description",
	}	
	PERMISSION_CREATE_DIRECT_CHANNEL = &Permission{
		"create_direct_channel",
		"permissions.create_direct_channel.name",
		"permissions.edit_otcreate_direct_channelher_users.description",
	}
	PERMISSION_MANAGE_CHANNEL_MEMBERS = &Permission{
		"manage_channel_members",
		"permissions.manage_channel_members.name",
		"permissions.manage_channel_members.description",
	}	
	PERMISSION_READ_CHANNEL = &Permission{
		"read_channel",
		"read_channel.name",
		"read_channel.description",
	}	
}

func InitRoles() {
	InitalPermissions()
	BuiltInRoles = make(map[string]*Role)

	ROLE_SYSTEM_ADMIN = &Role{
		"system_admin",
		"roles.channel_user.name",
		"roles.channel_user.description",
		[]string{
			PERMISSION_MANAGE_SYSTEM.Id,
			PERMISSION_EDIT_OTHER_USERS.Id,
			PERMISSION_CREATE_DIRECT_CHANNEL.Id,
		},
	}
	BuiltInRoles[ROLE_SYSTEM_ADMIN.Id] = ROLE_SYSTEM_ADMIN

	ROLE_NORMAL_USER = &Role{
		"normal_user",
		"roles.channel_user.name",
		"roles.channel_user.description",
		[]string{
			PERMISSION_CREATE_DIRECT_CHANNEL.Id,
		},
	}
	BuiltInRoles[ROLE_NORMAL_USER.Id] = ROLE_NORMAL_USER
	ROLE_CHANNEL_USER = &Role{
		"channel_user",
		"roles.channel_user.name",
		"roles.channel_user.description",
		[]string{
			PERMISSION_READ_CHANNEL.Id,
		},
	}
	BuiltInRoles[ROLE_CHANNEL_USER.Id] = ROLE_CHANNEL_USER
	ROLE_CHANNEL_ADMIN = &Role{
		"channel_admin",
		"roles.channel_admin.name",
		"roles.channel_admin.description",
		[]string{
			PERMISSION_MANAGE_CHANNEL_MEMBERS.Id,		
		},
	}
	BuiltInRoles[ROLE_CHANNEL_ADMIN.Id] = ROLE_CHANNEL_ADMIN
}

func init() {
	InitRoles()
}

func GetUserRoles(roles string) []string {
	return strings.Fields(roles)
}

func SessionHasPermissionTo(ctx *context.Context, permission *Permission) bool {
	roles := ctx.Input.Session("roles").(string)
	return IfRolesGrantPermission(GetUserRoles(roles), permission.Id)
}

func SessionHasPermissionToChannel(ctx *context.Context, channelId string, permission *Permission) bool {
	if channelId == "" {
		return false
	}
	userId := ctx.Input.Session("user_id").(string)
	cmc, err := GetAllChannelMembersForUser(userId); 
	if err != nil {
		return false
	}

	var channelRoles []string
	if roles, ok := cmc[channelId]; ok {
		channelRoles = strings.Fields(roles)
		if IfRolesGrantPermission(channelRoles, permission.Id) {
			return true
		}
	}

	return SessionHasPermissionTo(ctx, permission)
}

func SessionHasPermissionToUser(ctx *context.Context, userId string) bool {
	if userId == "" {
		return false
	}

	if ctx.Input.Session("user_id") == userId {
		return true
	}

	if SessionHasPermissionTo(ctx, PERMISSION_EDIT_OTHER_USERS) {
		return true
	}

	return false
}

func HasPermissionTo(askingUserId string, permission *Permission) bool {
	user, err := GetUser(askingUserId)
	if err != nil {
		return false
	}

	roles := user.GetRoles()

	return IfRolesGrantPermission(roles, permission.Id)
}

func HasPermissionToChannel(askingUserId string, channelId string, permission *Permission) bool {
	if channelId == "" || askingUserId == "" {
		return false
	}

	channelMember, err := GetChannelMember(channelId, askingUserId)
	if err == nil {
		roles := channelMember.GetRoles()
		if IfRolesGrantPermission(roles, permission.Id) {
			return true
		}
	}

	return HasPermissionTo(askingUserId, permission)
}

func HasPermissionToUser(askingUserId string, userId string) bool {
	if askingUserId == userId {
		return true
	}

	if HasPermissionTo(askingUserId, PERMISSION_EDIT_OTHER_USERS) {
		return true
	}

	return false
}

func IfRolesGrantPermission(roles []string, permissionId string) bool {
	for _, roleId := range roles {
		if role, ok := BuiltInRoles[roleId]; !ok {
			beego.Error("Bad role in system " + roleId)
			return false
		} else {
			permissions := role.Permissions
			for _, permission := range permissions {
				if permission == permissionId {
					return true
				}
			}
		}
	}

	return false
}

