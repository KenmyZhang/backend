package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"encoding/json"
)

type ChannelMember struct {
	Id           string    `orm:"pk"`
	ChannelId    string    `orm:"null"`
	UserId       string    `orm:"null"`
	Roles        string    `orm:"null"`
	LastViewedAt int64     `orm:"null"`
	MsgCount     int64     `orm:"null"`
	MentionCount int64     `orm:"null"`
	LastUpdateAt int64     `orm:"null"`
	IsDelete     bool      `orm:"null"`
}

func init() {
    orm.RegisterModel(new(ChannelMember))
}

// 多字段唯一键
func (o *ChannelMember) TableUnique() [][]string {
    return [][]string{
        []string{"ChannelId", "UserId"},
    }
}

func (o *ChannelMember) PreSave() {
	if o.Id == "" {
		o.Id = NewId()
	}	
	o.LastUpdateAt = GetMillis()
}

type ChannelMembers []ChannelMember

func (o *ChannelMembers) ToJson() string {
	if b, err := json.Marshal(o); err != nil {
		return "[]"
	} else {
		return string(b)
	}
}

func (o *ChannelMember) GetRoles() []string {
	return strings.Fields(o.Roles)
}

func GetChannelMember(channelId string, userId string) (*ChannelMember, error) {
	var member ChannelMember
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM ChannelMembers WHERE ChannelId = :ChannelId AND UserId = ?", 
		userId).QueryRows(&member)
	if err != nil {
    	return &member, err
	}	

	return  &member, nil
}			