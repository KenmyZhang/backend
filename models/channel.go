package models

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"sort"
	"strings"
	"strconv"
	"time"
	"errors"
	"regexp"
	"fmt"
	"unicode/utf8"
	"github.com/astaxie/beego/orm"
)

const (
	CHANNEL_DIRECT                 = "D"
	CHANNEL_GROUP                  = "G"
	CHANNEL_GROUP_MAX_USERS        = 500
	CHANNEL_GROUP_MIN_USERS        = 3
	CHANNEL_DISPLAY_NAME_MAX_RUNES = 64
	CHANNEL_NAME_MIN_LENGTH        = 2
	CHANNEL_NAME_MAX_LENGTH        = 26
	CHANNEL_HEADER_MAX_RUNES       = 1024
	CHANNEL_PURPOSE_MAX_RUNES      = 250
	CHANNEL_CACHE_SIZE             = 25000
)

type Channel struct {
	Id            string   `orm:"pk" json:"id"`
	CreateAt      int64    `orm:"null"`
	UpdateAt      int64    `orm:"null"`
	DeleteAt      int64    `orm:"null"`
	Type          string   `orm:"null"`
	DisplayName   string   `orm:"null"`
	Name          string   `orm:"size(50); unique" valid:"Required"`
	Header        string   `orm:"null"`
	Purpose       string   `orm:"null"`
	LastPostAt    int64    `orm:"null"`
	TotalMsgCount int64    `orm:"null"`
	CreatorId     string   `orm:"null"`
}

type ChannelPatch struct {
	DisplayName *string `json:"display_name"`
	Name        *string `json:"name"`
	Header      *string `json:"header"`
	Purpose     *string `json:"purpose"`
}

type allChannelMemberRoles struct {
	ChannelId string "channelId"
	Roles     string "roles"
}

type ChannelUserIds struct {
	UserIds []string `json:"userIds"`
}

func init() {
    orm.RegisterModel(new(Channel))
}

func (o *Channel) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func (o *ChannelPatch) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func ChannelFromJson(data io.Reader) *Channel {
	decoder := json.NewDecoder(data)
	var o Channel
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

func ChannelPatchFromJson(data io.Reader) *ChannelPatch {
	decoder := json.NewDecoder(data)
	var o ChannelPatch
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

func (o *Channel) Etag() string {
	return Etag(o.Id, o.UpdateAt)
}

var CurrentVersion string = "0.0.1"

func Etag(parts ...interface{}) string {

	etag := CurrentVersion

	for _, part := range parts {
		etag += fmt.Sprintf(".%v", part)
	}

	return etag
}

func (o *Channel) IsValid() error {

	if len(o.Id) != 26 {
		return errors.New("channel Id is invalid")
	}

	if o.CreateAt == 0 {
		return errors.New("channel CreateAt is invalid")
	}

	if o.UpdateAt == 0 {
		return errors.New("channel UpdateAt is invalid")	
	}

	if utf8.RuneCountInString(o.DisplayName) > CHANNEL_DISPLAY_NAME_MAX_RUNES {
		return errors.New("channel DisplayName is invalid")	
	}

	if !IsValidChlIdentifier(o.Name) {
		return errors.New("channel Name is invalid")	
	}

	if o.Type == CHANNEL_DIRECT || o.Type == CHANNEL_GROUP {
		errors.New("channel Type is invalid")	
	}

	if utf8.RuneCountInString(o.Header) > CHANNEL_HEADER_MAX_RUNES {
		return errors.New("channel Header is invalid")		
	}

	if utf8.RuneCountInString(o.Purpose) > CHANNEL_PURPOSE_MAX_RUNES {
		return errors.New("channel Purpose is invalid")
	}

	if len(o.CreatorId) > 26 {
		return errors.New("channel CreatorId is invalid")	
	}

	return nil
}

func (o *Channel) PreSave() {
	if o.Id == "" {
		o.Id = NewId()
	}

	o.CreateAt = GetMillis()
	o.UpdateAt = o.CreateAt
}

func (o *Channel) PreUpdate() {
	o.UpdateAt = GetMillis()
}

func (o *Channel) IsGroupOrDirect() bool {
	return o.Type == CHANNEL_DIRECT || o.Type == CHANNEL_GROUP
}

func (o *Channel) Patch(patch *ChannelPatch) {
	if patch.DisplayName != nil {
		o.DisplayName = *patch.DisplayName
	}

	if patch.Name != nil {
		o.Name = *patch.Name
	}

	if patch.Header != nil {
		o.Header = *patch.Header
	}

	if patch.Purpose != nil {
		o.Purpose = *patch.Purpose
	}
}

func GetDMNameFromIds(userId1, userId2 string) string {
	if userId1 > userId2 {
		return userId2 + "__" + userId1
	} else {
		return userId1 + "__" + userId2
	}
}

func GetGDisplayNameFromUsers(users []*User, truncate bool) string {
	usernames := make([]string, len(users))
	for index, user := range users {
		usernames[index] = user.Nickname
	}

	sort.Strings(usernames)

	name := strings.Join(usernames, ", ")

	if truncate && utf8.RuneCountInString(name) > CHANNEL_NAME_MAX_LENGTH {
		name = name[:CHANNEL_NAME_MAX_LENGTH]
	}

	return name
}

func IsValidChlIdentifier(s string) bool {

	validAlphaNumHyphenUnderscore := regexp.MustCompile(`^[a-z0-9]+([a-z\-\_0-9]+|(__)?)[a-z0-9]+$`)
	return validAlphaNumHyphenUnderscore.MatchString(s)

	if len(s) < CHANNEL_NAME_MIN_LENGTH {
		return false
	}

	return true
}

func GetGroupNameFromUserIds(userIds []string) string {
	sort.Strings(userIds)

	h := sha1.New()
	for _, id := range userIds {
		io.WriteString(h, id)
	}
	io.WriteString(h, strconv.FormatInt(time.Now().UnixNano() / int64(time.Millisecond), 10))

	return hex.EncodeToString(h.Sum(nil))
}

func GetAllChannelMembersForUser(userId string) (map[string]string, error) {
	var data []allChannelMemberRoles
	ids := make(map[string]string)
	o := orm.NewOrm()
	_, err := o.Raw("SELECT ChannelId, Roles FROM Channels, ChannelMembers WHERE Channels.Id = ChannelMembers.ChannelId AND ChannelMembers.UserId = ? AND Channels.DeleteAt = 0", 
		userId).QueryRows(&data)
	if err != nil {
    	return ids, err
	}

	for i := range data {
		ids[data[i].ChannelId] = data[i].Roles
	}
	return  ids, nil
}	

func CreateDirectChannel(creatorId string, otherUserId string) (*Channel, error) {
	if _, err := GetUser(creatorId); err != nil {
		return nil, err
	}

	if _, err := GetUser(otherUserId); err != nil {
		return nil, err
	}

	channel := new(Channel)

	channel.DisplayName = ""
	channel.Name = GetDMNameFromIds(otherUserId, creatorId)

	channel.Header = ""
	channel.Type = CHANNEL_DIRECT
	channel.CreatorId = creatorId

	o := orm.NewOrm()
	if err := o.Begin(); err != nil {
		return nil, err
	}
	// 事务处理过程
	channel.PreSave()
	if _, err := o.Insert(channel); err != nil {
		o.Rollback()
		return nil, err
	}

	cm1 := &ChannelMember{
		UserId:      creatorId,
		Roles:       ROLE_CHANNEL_USER.Id,
		ChannelId:   channel.Id,
	}
	cm2 := &ChannelMember{
		UserId:      otherUserId,
		Roles:       ROLE_CHANNEL_USER.Id,
		ChannelId:   channel.Id,		
	}
	cm1.PreSave()
	if _, err := o.Insert(cm1); err != nil {
		o.Rollback()
		return nil, err
	}	

	cm2.PreSave()
	if _, err := o.Insert(cm2); err != nil {
		o.Rollback()
		return nil, err
	}	

	// 此过程中的所有使用 o Ormer 对象的查询都在事务处理范围内
    if err := o.Commit(); err != nil {
    	return nil, err
    }
	return channel, nil
}
