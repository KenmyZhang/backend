package models

import (
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

type Post struct {
	Id            string          `bson:"_id" json:"id"`
	CreateAt      int64           `bson:"createAt" json:"create_at"`
	UpdateAt      int64           `bson:"updateAt" json:"update_at"`
	EditAt        int64           `bson:"editAt" json:"edit_at"`
	DeleteAt      int64           `bson:"deleteAt" json:"delete_at"`
	UserId        string          `bson:"userId" json:"user_id"`
	ChannelId     string          `bson:"channelId" json:"channel_id"`
	Message       string          `bson:"message" json:"message"`
	Type          string          `bson:"type" json:"type"`
	Props         map[string] interface{} `bson:"props" json:"props"`
	Filenames     []string     `bson:"filenames" json:"filenames,omitempty"` // Deprecated, do not use this field any more
	FileIds       []string     `bson:"fileIds" json:"file_ids,omitempty"`
	FileTypes     []string     `bson:"fileTypes" json:"file_types,omitempty"`
	PendingPostId string          `bson:"-" db:"-" json:"pending_post_id"`
	LastPictureUpdate  int64      `bson:"lastPictureUpdate" json:"last_picture_update,omitempty"`
}

func PostToJson(u *Post) string {
	b, err := json.Marshal(u)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func PostsToJson(u *[]Post) string {
	b, err := json.Marshal(u)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}



func (o *Post) PreSave() {
	if o.Id == "" {
		o.Id = NewId()
	}

	if o.CreateAt == 0 {
		o.CreateAt = GetMillis()
	}

	o.UpdateAt = o.CreateAt

	if o.Props == nil {
		o.Props = make(map[string]interface{})
	}

	if o.Filenames == nil {
		o.Filenames = []string{}
	}

	if o.FileIds == nil {
		o.FileIds = []string{}
	}
}

func CreatePost(post *Post) (*Post, error) {
	post.PreSave()
    o := orm.NewOrm()
    o.Using("default")
	if _, err := o.Insert(post); err != nil {
		return nil, err
	}
	return post, nil
}

func GetPost(channelId string, page, perpage int) ( *[]Post, error) {
	posts := &[]Post{}
    o := orm.NewOrm()
    o.Using("default")

	return posts,nil
}