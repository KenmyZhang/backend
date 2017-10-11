package models

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"strings"
)

type ChannelMember struct {
	Id           string `orm:"pk" json:"id"`
	ChannelId    string `orm:"null" json:"channel_id"`
	UserId       string `orm:"null" json:"user_id"`
	Roles        string `orm:"null" json:"roles"`
	LastViewedAt int64  `orm:"null" json:"last_viewed_at"`
	MsgCount     int64  `orm:"null" json:"msg_count"`
	MentionCount int64  `orm:"null" json:"mention_count"`
	LastUpdateAt int64  `orm:"null" json:"last_update_at"`
	IsDelete     bool   `orm:"null" json:"is_delete"`
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

	return &member, nil
}
