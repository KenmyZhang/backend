package models

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

var ROLE_SYSTEM_ADMIN *Role
var ROLE_NORMAL_USER *Role 


var BuiltInRoles map[string]*Role

func InitalizePermissions() {
	PERMISSION_MANAGE_SYSTEM = &Permission{
		"manage_system",
		"authentication.permissions.manage_system.name",
		"authentication.permissions.manage_system.description",
	}
}

func InitalizeRoles() {
	InitalizePermissions()
	BuiltInRoles = make(map[string]*Role)

	ROLE_SYSTEM_ADMIN = &Role{
		"system_admin",
		"authentication.roles.channel_user.name",
		"authentication.roles.channel_user.description",
		[]string{
			PERMISSION_MANAGE_SYSTEM.Id,
		},
	}
	BuiltInRoles[ROLE_SYSTEM_ADMIN.Id] = ROLE_SYSTEM_ADMIN

	ROLE_NORMAL_USER = &Role{
		"normal_user",
		"authentication.roles.channel_user.name",
		"authentication.roles.channel_user.description",
		[]string{},
	}
	BuiltInRoles[ROLE_NORMAL_USER.Id] = ROLE_NORMAL_USER

}

func init() {
	InitalizeRoles()
}
